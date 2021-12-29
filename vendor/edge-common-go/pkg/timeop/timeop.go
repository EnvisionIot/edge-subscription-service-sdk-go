package timeop

import (
	"edge-common-go/pkg/metadata"
	"errors"
	"strconv"
	"strings"
	"sync"
	"time"
)

/**
 * Description:一些时区相关操作
 *
 * @author yang.zhang4
 * @date 2020/12/25
 */

var (
	defaultTimePattern   = "2006-01-02 15:04:05.000"
	timeLocationMap      = make(map[string]*time.Location)
	timeLocationMapMutex sync.RWMutex //操作timeLocationMap的锁
)

//根据城市名或"+08:00"的时区格式获取时区的秒
func GetTimezoneSecondByStr(timezoneStr string) (int64, error) {
	if len(timezoneStr) <= 0 {
		return 0, errors.New("timezoneStr empty error")
	}

	if strings.HasPrefix(timezoneStr, "+") || strings.HasPrefix(timezoneStr, "-") {
		return TransformDigitalTimezoneStrToSecond(timezoneStr), nil
	}

	return TransformLocationTimezoneStrToSecond(timezoneStr), nil
}

//根据指定时区获取时间字符串，timezoneStr可以是地区名可以是+08:00这种格式，+08:00的+或-不可以省略
func GetLocalTimeStrByUnixWithTimezone(timeNSec int64, timezoneStr string, timePattern string) string {
	var tm time.Time
	if len(timePattern) <= 0 {
		timePattern = defaultTimePattern
	}
	var timeStr string

	var timezoneOffset int64
	timezoneOffset, _ = GetTimezoneSecondByStr(timezoneStr)
	timeNSec += timezoneOffset * 1e9

	//因为已经在上面做了时区偏移的处理，这里需要按照UTC时区操作
	tm = time.Unix(0, timeNSec).UTC()
	timeStr = tm.Format(timePattern)

	return timeStr
}

//根据指定设备，查询设备所在时区，获取时间字符串，timezoneStr可以是地区名可以是+08:00这种格式，+08:00的+或-不可以省略
func GetLocalTimeStrByUnixWithAssetId(metaData *metadata.MetaData, assetId string, timeNSec int64, timePattern string) string {
	//获取设备时区
	var tslInstanceData metadata.TSLInstanceData
	var timezoneStr = "+00:00"
	var err error

	tslInstanceData, err = metadata.GetTSLInstanceDataByAssetId(metaData.Base.ModelAsset, assetId)
	if err == nil {
		timezoneStr = tslInstanceData.Timezone
	}

	var timeStrRet string
	timeStrRet = GetLocalTimeStrByUnixWithTimezone(timeNSec, timezoneStr, timePattern)

	return timeStrRet
}

//根据指定设备，查询设备所在时区，获取时区的毫秒，timezoneStr可以是地区名可以是+08:00这种格式，+08:00的+或-不可以省略
func GetTimezoneMsWithAssetId(metaData *metadata.MetaData, assetId string) (int64, error) {
	//获取设备时区
	var tslInstanceData metadata.TSLInstanceData
	var timezoneStr = "+00:00"
	var err error

	tslInstanceData, err = metadata.GetTSLInstanceDataByAssetId(metaData.Base.ModelAsset, assetId)
	if err == nil {
		timezoneStr = tslInstanceData.Timezone
	}

	var timezoneOffsetSec int64
	timezoneOffsetSec, _ = GetTimezoneSecondByStr(timezoneStr)

	var timezoneOffsetMs int64
	timezoneOffsetMs = timezoneOffsetSec * 1e3

	return timezoneOffsetMs, nil
}

//根据本机时区获取时间字符串
func GetLocalTimeStrByUnix(timeSec int64, timeNSec int64, timePattern string) string {
	var tm time.Time
	if len(timePattern) <= 0 {
		timePattern = defaultTimePattern
	}
	var timeStr string
	tm = time.Unix(timeSec, timeNSec)
	timeStr = tm.Format(timePattern)

	return timeStr
}

//根据城市地区名的时区格式统一转换为秒，地区的解析LoadLocation函数每次都要去解时区信息的压缩文件，比较慢，所以这里搞了缓存
func TransformLocationTimezoneStrToSecond(zoneStr string) int64 {
	var loc *time.Location
	var ok bool
	timeLocationMapMutex.RLock()
	loc, ok = timeLocationMap[zoneStr]
	timeLocationMapMutex.RUnlock()

	if ok {
		if loc == nil {
			return 0
		}
		var offset int
		_, offset = time.Now().In(loc).Zone()
		return int64(offset)
	}

	var err error
	var secondRet int64
	loc, err = time.LoadLocation(zoneStr)
	if err != nil {
		loc = nil
		secondRet = 0
	} else {
		var offset int
		_, offset = time.Now().In(loc).Zone()
		secondRet = int64(offset)
	}

	timeLocationMapMutex.Lock()
	timeLocationMap[zoneStr] = loc
	timeLocationMapMutex.Unlock()

	return secondRet
}

//根据"+08:00"的时区格式统一转换为秒，这里解析比较快，不需要搞缓存
func TransformDigitalTimezoneStrToSecond(zoneStr string) int64 {
	var zoneStrNumber = zoneStr[1:]
	var strArray []string
	strArray = strings.Split(zoneStrNumber, ":")

	var secondRet int64
	var numberTemp int64
	var err error
	if len(strArray) >= 1 {
		numberTemp, err = strconv.ParseInt(strArray[0], 10, 64)
		if err == nil {
			secondRet += numberTemp * 3600
		}
	}

	if len(strArray) >= 2 {
		numberTemp, err = strconv.ParseInt(strArray[1], 10, 64)
		if err == nil {
			secondRet += numberTemp * 60
		}
	}

	if strings.HasPrefix(zoneStr, "-") {
		return secondRet * (-1)
	}

	return secondRet
}
