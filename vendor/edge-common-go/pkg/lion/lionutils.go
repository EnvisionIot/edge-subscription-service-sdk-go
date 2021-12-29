package lion

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"reflect"
	"sort"
	"strconv"
	"strings"
	"sync"
)

/**
 * @Author: qianjialin
 * @Date: 2020/6/9 11:30
 */
var (
	//本地运行请确认DefaultLionConfigPath下存在lion.properties文件
	DefaultLionConfigPath = "/home/envuser/data/apps/config/lionconfig/config/lion.properties"
	//used for debug
	PrivateLionConfigPath = "../resources/lion_windows.properties"
)

var lionCache = make(map[string]interface{})
var appLion = make(map[string]interface{})
var appLionKeySlice []string
var appLionMutex sync.RWMutex

func (l *Default_LionConfig) LoadLionConfig(lionConfig interface{}) error {
	f, err := os.Open(DefaultLionConfigPath)
	if err != nil {
		fmt.Println("open DefaultLionConfigPath failed path=", DefaultLionConfigPath)
		fmt.Println("using windows test file")
		f, err = os.Open(PrivateLionConfigPath)
		if err != nil {
			fmt.Println("open windows debug file error", PrivateLionConfigPath)
		}

	}
	defer f.Close()

	rd := bufio.NewReader(f)
	var firstEOF = true
	for {
		line, err := rd.ReadString('\n') //以'\n'为结束符读入一行

		if err != nil && err != io.EOF {
			fmt.Println("read lion.properties error=", err.Error())
			break
		}

		if err == io.EOF && !firstEOF {
			break
		}

		if err == io.EOF && firstEOF {
			firstEOF = false
		}

		line = strings.Trim(line, " \r\n\t")

		if strings.HasPrefix(line, "#") {
			continue
		} else {
			keyValues := strings.Split(line, "=")
			if len(keyValues) == 2 {
				lionCache[keyValues[0]] = strings.Trim(keyValues[1], "\n")
			}
			if len(keyValues) > 2 {
				var index = 0
				var value = ""
				for _, keyValue := range keyValues {
					index++
					if index == 1 {
						continue
					}
					value = value + strings.Trim(keyValue, "\n") + "="
				}
				value = string([]rune(value)[:len(value)-1])
				lionCache[keyValues[0]] = value
			}
		}
	}

	//初始化一下一些aes相关的配置
	initLionAes()
	//解密需要解密的lion
	decryptLion()

	l.init(lionConfig)
	return nil
}

type Default_LionConfig struct {
}

func (l *Default_LionConfig) init(child interface{}) {
	ref := reflect.ValueOf(child)
	method := ref.MethodByName("InitLionConfig")
	if method.IsValid() {
		method.Call(make([]reflect.Value, 0))
	} else {
		fmt.Println("cloud not find InitLionConfig func")
	}
}

func (l *Default_LionConfig) InitLionConfig() {
	fmt.Println("default InitLionConfig func")
}

func GetAppLion() map[string]interface{} {
	appLionMutex.RLock()
	defer appLionMutex.RUnlock()
	return appLion
}

func GetAppLionWithOrderedSlice() (map[string]interface{}, []string) {
	appLionMutex.RLock()
	defer appLionMutex.RUnlock()
	sort.Strings(appLionKeySlice)
	return appLion, appLionKeySlice
}

func GetBoolValue(key string, defaultValue bool) bool {
	if _, exist := lionCache[key]; exist {
		res, ok := lionCache[key].(string)
		if ok {
			b, err := strconv.ParseBool(res)
			if err == nil {
				appLionMutex.Lock()
				defer appLionMutex.Unlock()
				appLion[key] = b
				appLionKeySlice = append(appLionKeySlice, key)
				return b
			}
		}
	}
	appLionMutex.Lock()
	defer appLionMutex.Unlock()
	appLion[key] = defaultValue
	appLionKeySlice = append(appLionKeySlice, key)
	return defaultValue
}

func GetStringValue(key string, defaultValue string) string {
	if _, exist := lionCache[key]; exist {
		res, ok := lionCache[key].(string)
		if ok {
			appLionMutex.Lock()
			defer appLionMutex.Unlock()
			appLion[key] = res
			appLionKeySlice = append(appLionKeySlice, key)
			return res
		}
	}
	appLionMutex.Lock()
	defer appLionMutex.Unlock()
	appLion[key] = defaultValue
	appLionKeySlice = append(appLionKeySlice, key)
	return defaultValue
}

func GetIntValue(key string, defaultValue int) int {
	if _, exist := lionCache[key]; exist {
		res, ok := lionCache[key].(string)
		if ok {
			b, err := strconv.ParseInt(res, 10, 64)
			if err == nil {
				appLionMutex.Lock()
				defer appLionMutex.Unlock()
				appLion[key] = int(b)
				appLionKeySlice = append(appLionKeySlice, key)
				return int(b)
			}
		}
	}
	appLionMutex.Lock()
	defer appLionMutex.Unlock()
	appLion[key] = defaultValue
	appLionKeySlice = append(appLionKeySlice, key)
	return defaultValue
}

func GetInt32Value(key string, defaultValue int32) int32 {
	if _, exist := lionCache[key]; exist {
		res, ok := lionCache[key].(string)
		if ok {
			b, err := strconv.ParseInt(res, 10, 64)
			if err == nil {
				appLionMutex.Lock()
				defer appLionMutex.Unlock()
				appLion[key] = int32(b)
				appLionKeySlice = append(appLionKeySlice, key)
				return int32(b)
			}
		}
	}
	appLionMutex.Lock()
	defer appLionMutex.Unlock()
	appLion[key] = defaultValue
	appLionKeySlice = append(appLionKeySlice, key)
	return defaultValue
}

func GetInt64Value(key string, defaultValue int64) int64 {
	if _, exist := lionCache[key]; exist {
		res, ok := lionCache[key].(string)
		if ok {
			b, err := strconv.ParseInt(res, 10, 64)
			if err == nil {
				appLionMutex.Lock()
				defer appLionMutex.Unlock()
				appLion[key] = b
				appLionKeySlice = append(appLionKeySlice, key)
				return b
			}
		}
	}
	appLionMutex.Lock()
	defer appLionMutex.Unlock()
	appLion[key] = defaultValue
	appLionKeySlice = append(appLionKeySlice, key)
	return defaultValue
}

func GetFloat32Value(key string, defaultValue float32) float32 {
	if _, exist := lionCache[key]; exist {
		res, ok := lionCache[key].(string)
		if ok {
			b, err := strconv.ParseFloat(res, 32)
			if err == nil {
				appLionMutex.Lock()
				defer appLionMutex.Unlock()
				appLion[key] = float32(b)
				appLionKeySlice = append(appLionKeySlice, key)
				return float32(b)
			}
		}
	}
	appLionMutex.Lock()
	defer appLionMutex.Unlock()
	appLion[key] = defaultValue
	appLionKeySlice = append(appLionKeySlice, key)
	return defaultValue
}

func GetFloat64Value(key string, defaultValue float64) float64 {
	if _, exist := lionCache[key]; exist {
		res, ok := lionCache[key].(string)
		if ok {
			b, err := strconv.ParseFloat(res, 64)
			if err == nil {
				appLionMutex.Lock()
				defer appLionMutex.Unlock()
				appLion[key] = b
				appLionKeySlice = append(appLionKeySlice, key)
				return b
			}
		}
	}
	appLionMutex.Lock()
	defer appLionMutex.Unlock()
	appLion[key] = defaultValue
	appLionKeySlice = append(appLionKeySlice, key)
	return defaultValue
}
