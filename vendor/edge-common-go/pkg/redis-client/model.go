package redis_client

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
	"github.com/spf13/cast"
)

/**
 * @Author: qianjialin
 * @Date: 2020/8/11 15:04
 */
// CurrentRedisDao 实例，全局唯一
var CurrentRedisDao *RedisDao

type RedisDao struct {
	pool *redis.Pool
}

// 初始化一个 CurrentRedisDao 结构体
func InitRedisDao(pool *redis.Pool) (currentRedisDao *RedisDao) {
	currentRedisDao = &RedisDao{pool: pool}
	return
}

func (this *RedisDao) GET(redisKey string) (value string, err error) {
	conn := this.pool.Get()
	defer func() {
		_ = conn.Close()
	}()

	res, err := redis.String(conn.Do("GET", redisKey))
	if err != nil {
		return
	}
	value = res
	return
}

//func (this *RedisDao) SETEX(localCache *cache2go.CacheTable) (result bool, err error) {
//	conn := this.pool.Get()
//	defer func() {
//		_ = conn.Close()
//	}()
//	err = conn.Send("MULTI")
//	if err != nil {
//		result = false
//		return
//	}
//	localCache.Foreach(
//		func(key interface{}, entry *cache2go.CacheItem) {
//			var value = cast.ToString(entry.Data())
//			var expireSeconds = cast.ToInt64(entry.LifeSpan().Seconds())
//			if expireSeconds > 0 {
//				_, err = conn.Do("SETEX", key, expireSeconds, value)
//			} else if expireSeconds == 0 {
//				_, err = conn.Do("SET", key, value)
//			}
//		})
//	_, err = redis.Values(conn.Do("EXEC"))
//	if err != nil {
//		fmt.Printf("set local cache to reids error: %v", err)
//		result = false
//		return
//	}
//	result = true
//	return
//}

func (this *RedisDao) HGet(redisKey string, hashKey string) (value string, err error) {
	conn := this.pool.Get()
	defer func() {
		_ = conn.Close()
	}()

	res, err := redis.String(conn.Do("hget", redisKey, hashKey))
	if err != nil {
		err = ERROR_DOES_NOT_EXIST
		return
	}
	value = res
	return
}

//删除一个key下的一批field
func (this *RedisDao) HDel(key string, delMap map[string]string) error {
	var conn redis.Conn
	var err error
	conn = this.pool.Get()
	defer func() {
		_ = conn.Close()
	}()

	var fieldArray []interface{}
	fieldArray = append(fieldArray, key)
	for field := range delMap {
		fieldArray = append(fieldArray, field)
	}
	if len(fieldArray) <= 1 {
		return nil
	}
	_, err = conn.Do("hdel", fieldArray...)
	if err != nil {
		return err
	}

	return nil
}

//hmget
func (this *RedisDao) HmGet(key string, fieldMap map[string]string) (map[string][]byte, error) {
	if len(fieldMap) <= 0 {
		return nil, nil
	}

	var err error
	conn := this.pool.Get()
	defer func() {
		_ = conn.Close()
	}()

	var fieldArray []interface{}
	fieldArray = append(fieldArray, key)
	for field := range fieldMap {
		fieldArray = append(fieldArray, field)
	}

	getValue, err := redis.Values(conn.Do("hmget", fieldArray...))
	if err != nil {
		return nil, err
	}

	var mapRet map[string][]byte
	mapRet = make(map[string][]byte)
	var fieldArrayIndex int
	var fieldString string
	for index, value := range getValue {
		fieldArrayIndex = index + 1
		if fieldArrayIndex >= len(fieldArray) {
			break
		}

		if value == nil {
			continue
		}

		fieldString = fieldArray[fieldArrayIndex].(string)
		mapRet[fieldString] = value.([]byte)
	}

	return mapRet, nil
}

func (this *RedisDao) HGETALL(redisKey string) (value map[string]string, err error) {
	value = make(map[string]string)
	conn := this.pool.Get()
	defer func() {
		_ = conn.Close()
	}()
	err = conn.Send("MULTI")
	if err != nil {
		return
	}
	err = conn.Send("hgetall", redisKey)
	if err != nil {
		return
	}
	res, err := redis.Values(conn.Do("EXEC"))
	if err != nil {
		return
	}
	var temp []interface{}
	var tempStr [2]string
	for _, resList := range res {
		temp = resList.([]interface{})
		index := 0
		for _, tempList := range temp {
			tempStr[index] = string(tempList.([]uint8))
			index++
			if index >= 2 {
				index = 0
				value[tempStr[0]] = tempStr[1]
			}
		}
	}
	return
}

func (this *RedisDao) GETByPattern(keyPattern string) (map[string]map[string]int64, error) {
	conn := this.pool.Get()
	defer func() {
		_ = conn.Close()
	}()
	var result = make(map[string]map[string]int64) //key.value.ttl
	res, err := redis.Values(conn.Do("KEYS", "*"+keyPattern+"*"))
	if err != nil {
		return nil, err
	}
	for _, key := range res {
		var keyStr = cast.ToString(key)
		value, err := redis.String(conn.Do("GET", keyStr))
		if err != nil || value == "" {
			continue
		}
		ttl, err := redis.Int64(conn.Do("TTL", keyStr))
		if err != nil || ttl < 0 {
			continue
		}
		if _, exist := result[keyStr]; !exist {
			result[keyStr] = make(map[string]int64)
		}
		result[keyStr][value] = ttl
	}
	return result, nil
}

func (this *RedisDao) GETKeysByPattern(keyPattern string) (map[string]string, error) {
	var err error
	conn := this.pool.Get()
	defer func() {
		_ = conn.Close()
	}()

	getValue, err := redis.Values(redis.Values(conn.Do("KEYS", keyPattern)))
	if err != nil {
		return nil, err
	}

	var keyMapRet = make(map[string]string)
	for _, key := range getValue {
		var keyStr string
		keyStr, err = cast.ToStringE(key)
		if err != nil {
			continue
		}
		keyMapRet[keyStr] = "1"
	}

	return keyMapRet, nil
}

//和HGETALL一样，不过返回的是map[string][]byte，因为[]byte转为string是一个比较耗时的操作
func (this *RedisDao) HGETALLBytes(redisKey string) (map[string][]byte, error) {
	var err error
	conn := this.pool.Get()
	defer func() {
		_ = conn.Close()
	}()

	getValue, err := redis.Values(conn.Do("hgetall", redisKey))
	if err != nil {
		return nil, err
	}

	var key string
	var flag bool
	flag = true
	var retMap = make(map[string][]byte)
	for _, value := range getValue {
		if flag {
			key = string(value.([]byte))
			flag = false
		} else {
			retMap[key] = value.([]byte)
			flag = true
		}
	}
	return retMap, err
}

func (this *RedisDao) HSET(redisKey string, hashKey string, value string, expireSeconds int) (result bool, err error) {
	conn := this.pool.Get()
	defer func() {
		_ = conn.Close()
	}()
	_, err = conn.Do("hset", redisKey, hashKey, value)
	if expireSeconds != 0 {
		_, err = conn.Do("expire", redisKey, hashKey, expireSeconds)
	}
	if err != nil {
		fmt.Printf("set to reids error: %v", err)
		return
	}
	result = true
	return
}

//写入map数据，key map[field]data，过期时间为整个key的过期时间
func (this *RedisDao) SingleHmSet(key string, fieldMap map[string][]byte, expireSeconds int64) error {
	if len(fieldMap) <= 0 {
		return nil
	}

	var err error
	conn := this.pool.Get()
	defer func() {
		_ = conn.Close()
	}()

	var fieldArray []interface{}
	fieldArray = append(fieldArray, key)

	for field, data := range fieldMap {
		fieldArray = append(fieldArray, field, data)
	}
	if len(fieldArray) <= 1 {
		return fmt.Errorf("fieldMap is empty")
	}

	_, err = conn.Do("hmset", fieldArray...)
	if err != nil {
		return err
	}

	if expireSeconds > 0 {
		_, err = conn.Do("expire", key, expireSeconds)
		if err != nil {
			return err
		}
	}

	return nil
}

//写入map数据，key map[field]data，过期时间为整个key的过期时间，先删除原来的key，再写
func (this *RedisDao) SingleHmSetOverride(key string, fieldMap map[string][]byte, expireSeconds int64) error {
	if len(fieldMap) <= 0 {
		return nil
	}

	var err error
	conn := this.pool.Get()
	defer func() {
		_ = conn.Close()
	}()

	var fieldArray []interface{}
	fieldArray = append(fieldArray, key)

	for field, data := range fieldMap {
		fieldArray = append(fieldArray, field, data)
	}
	if len(fieldArray) <= 1 {
		return fmt.Errorf("fieldMap is empty")
	}

	err = conn.Send("MULTI")
	if err != nil {
		return err
	}
	err = conn.Send("del", key)
	if err != nil {
		return err
	}
	err = conn.Send("hmset", fieldArray...)
	if err != nil {
		return err
	}

	if expireSeconds > 0 {
		err = conn.Send("expire", key, expireSeconds)
		if err != nil {
			return err
		}
	}

	_, err = conn.Do("EXEC")

	return err
}

//批量写入map数据，map[key][field]data，过期时间为所有key的过期时间
func (this *RedisDao) MultiHmSet(multiData map[string]map[string][]byte, expireSeconds int64) error {
	if len(multiData) <= 0 {
		return nil
	}

	var err error
	conn := this.pool.Get()
	defer func() {
		_ = conn.Close()
	}()

	for key, fieldMap := range multiData {
		var fieldArray []interface{}
		fieldArray = append(fieldArray, key)
		for field, data := range fieldMap {
			fieldArray = append(fieldArray, field, data)
		}
		if len(fieldArray) <= 1 {
			continue
		}
		_ = conn.Send("hmset", fieldArray...)
		if expireSeconds > 0 {
			_ = conn.Send("expire", key, expireSeconds)
		}
	}
	err = conn.Flush()
	if err != nil {
		return err
	}

	return nil
}

//func (this *RedisDao) FlushSolarSiteGroupRecordCache(redisKey string, localCache *cache2go.CacheTable) (result bool, err error) {
//	conn := this.pool.Get()
//	defer func() {
//		_ = conn.Close()
//	}()
//	err = conn.Send("MULTI")
//	if err != nil {
//		return
//	}
//	localCache.Foreach(
//		func(key interface{}, entry *cache2go.CacheItem) {
//			//value=api.Record 转换成 StreamingStructure 再转换为string(StreamingStructureForRedis)
//			var record = entry.Data().(api.Record)
//			fieldMap, _ := record.Get()
//			mapValue := fieldMap.Value.(map[string]*api.Field)
//			valueField, _ := mapValue["value"]
//			value, _ := valueField.GetValueAsInterface()
//			attrField, _ := mapValue["attr"]
//			attr, _ := attrField.GetValueAsMap()
//			StreamingStructure, err := stagescommon.TransformStreamingRecordToSS(record)
//			if err == nil {
//				var StreamingStructureForRedis = stagescommon.StreamingStructureForRedis{
//					OrgId:      StreamingStructure.OrgId,
//					ModelId:    StreamingStructure.ModelId,
//					AssetId:    StreamingStructure.AssetId,
//					PointId:    StreamingStructure.PointId,
//					Time:       StreamingStructure.Time,
//					Value:      value,
//					Quality:    StreamingStructure.Quality,
//					Dq:         StreamingStructure.Dq,
//					Attr:       attr,
//					ModelPath:  StreamingStructure.ModelPath,
//					PolicyType: StreamingStructure.PolicyType,
//				}
//				value, err := json.Marshal(StreamingStructureForRedis)
//				if err == nil {
//					var v = string(value)
//					conn.Send("hset", redisKey, key, v)
//				}
//			}
//		})
//	_, err = redis.Values(conn.Do("EXEC"))
//	if err != nil {
//		fmt.Printf("set local cache to reids error: %v", err)
//		return
//	}
//	result = true
//	return
//}

//func (this *RedisDao) FlushJSAccumulateCache(redisKey string, lastRecord map[string]commonutils.LastValue, accumulateValue map[string]float64) (result bool, err error) {
//	conn := this.pool.Get()
//	defer func() {
//		_ = conn.Close()
//	}()
//	conn.Send("MULTI")
//	for key, lastValue := range lastRecord {
//		var valueForRedis commonutils.AccumulateValueForRedis
//		valueForRedis.LastValue = lastValue.Value
//		valueForRedis.LastTimestamp = lastValue.Timestamp
//		accumulateValueI, exist := accumulateValue[key]
//		if exist {
//			valueForRedis.Value = accumulateValueI
//			value, err := json.Marshal(valueForRedis)
//			if err == nil {
//				var v = string(value)
//				conn.Send("hset", redisKey, key, v)
//			}
//		}
//	}
//	_, err = redis.Values(conn.Do("EXEC"))
//	if err != nil {
//		fmt.Printf("set local cache to reids error: %v", err)
//		return
//	}
//	result = true
//	return
//}
