package utils

import (
	"github.com/mohae/deepcopy"
	"math"
	"os"
	"path/filepath"
	"runtime"
	"strings"
)

/**
* Description:
*
* @author yang.zhang4
* @date 2019/6/15
 */

const (
	UNABLE_TO_CAST_TYPE = "unable to cast %#v of type %T to "
	NO_SUCH_KEY         = "no such key=%s"
	FLOAT64             = "float64"
	STRING              = "string"
	INT32_SLICE         = "[]int32"
	INT64_SLICE         = "[]int64"
	FLOAT32_SLICE       = "[]float32"
	FLOAT64_SLICE       = "[]float64"
	STRING_SLICE        = "[]string"
)

//判断文件夹是否存在
func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, err
	}
	return false, err
}

//获取baseDir
func GetBaseDir() string {
	ex, err := os.Executable()
	if err != nil {
		panic(err)
	}
	baseDir := strings.TrimSuffix(filepath.Dir(ex), "/bin")
	baseDir = strings.TrimSuffix(baseDir, "\\bin") // for windows
	return baseDir
}

//获取文件名，大约执行一次几百纳秒
func GetFileName() string {
	_, file, _, ok := runtime.Caller(1)
	if !ok {
		return ""
	}
	return file
}

//获取文件名，只取最后的n级
func GetFileNameLastN(n int) string {
	_, file, _, ok := runtime.Caller(1)
	if !ok {
		return ""
	}
	var sep string = "/"
	if strings.Index(file, "/") >= 0 {
		sep = "/"
	} else {
		sep = "\\"
	}

	//sep := string(os.PathSeparator)

	var stringArray []string
	stringArray = strings.Split(file, sep)

	if len(stringArray) <= n {
		return file
	}

	var stringArrayNew = stringArray[len(stringArray)-n:]

	var filePathNew string
	filePathNew = strings.Join(stringArrayNew, sep)

	return filePathNew
}

//获取行号，大约执行一次几百纳秒
func GetLineNum() int64 {
	_, _, line, ok := runtime.Caller(1)
	if !ok {
		return 0
	}
	return int64(line)
}

//获取函数名，大约执行一次几百纳秒
func GetFuncName() string {
	pc, _, _, ok := runtime.Caller(1)
	if !ok {
		return ""
	}

	funcTemp := runtime.FuncForPC(pc)
	if funcTemp == nil {
		return ""
	}

	return funcTemp.Name()
}

//比较浮点数，x取值举例：0.000001
func FloatIsEqual(f1 float64, f2 float64) bool {
	var gap = math.Abs(f1 - f2)
	if gap <= 0.000001 {
		return true
	}

	return false
}

//比较浮点数，x取值举例：0.000001
func FloatIsEqualX(f1 float64, f2 float64, x float64) bool {
	var gap = math.Abs(f1 - f2)
	if gap <= x {
		return true
	}

	return false
}

//获取字符串指针
func NewStringPtr(str string) *string {
	return &str
}

//从字符串指针获取值，为nil就使用默认值
func GetValueFromStringPtr(str *string, defaultValue string) string {
	if str == nil {
		return defaultValue
	}
	return *str
}

//拷贝字符串指针，如果原指针为nil，则使用默认值指针
func CopyStringPtrNotNil(valuePtr *string, defaultValue string) *string {
	if valuePtr != nil {
		return valuePtr
	}
	return &defaultValue
}

//获取bool指针
func NewBoolPtr(value bool) *bool {
	return &value
}

//从bool指针获取值，为nil就使用默认值
func GetValueFromBoolPtr(valuePtr *bool, defaultValue bool) bool {
	if valuePtr == nil {
		return defaultValue
	}
	return *valuePtr
}

//拷贝bool指针，如果原指针为nil，则使用默认值指针
func CopyBoolPtrNotNil(valuePtr *bool, defaultValue bool) *bool {
	if valuePtr != nil {
		return valuePtr
	}
	return &defaultValue
}

//获取int32指针
func NewInt32Ptr(value int32) *int32 {
	return &value
}

//从int32指针获取值，为nil就使用默认值
func GetValueFromInt32Ptr(valuePtr *int32, defaultValue int32) int32 {
	if valuePtr == nil {
		return defaultValue
	}
	return *valuePtr
}

//拷贝int32指针，如果原指针为nil，则使用默认值指针
func CopyInt32PtrNotNil(valuePtr *int32, defaultValue int32) *int32 {
	if valuePtr != nil {
		return valuePtr
	}
	return &defaultValue
}

//获取int64指针
func NewInt64Ptr(value int64) *int64 {
	return &value
}

//从int64指针获取值，为nil就使用默认值
func GetValueFromInt64Ptr(valuePtr *int64, defaultValue int64) int64 {
	if valuePtr == nil {
		return defaultValue
	}
	return *valuePtr
}

//拷贝int64指针，如果原指针为nil，则使用默认值指针
func CopyInt64PtrNotNil(valuePtr *int64, defaultValue int64) *int64 {
	if valuePtr != nil {
		return valuePtr
	}
	return &defaultValue
}

//判断一个string是否在一个string数组中
func StringIsInArray(strArray []string, str string) bool {
	for _, v := range strArray {
		if v == str {
			return true
		}
	}

	return false
}

//合并多个string数组，只添加不重复的元素
func CombineStringArray(arrays ...[]string) []string {
	var resultArray []string
	var resultMap = make(map[string]string)
	var ok bool

	for _, arrayTemp := range arrays {
		for _, v := range arrayTemp {
			_, ok = resultMap[v]
			if ok {
				continue
			}
			resultMap[v] = "1"
			resultArray = append(resultArray, v)
		}
	}

	return resultArray
}

//深拷贝map[string]interface{}
func DeepCopyMap(mapTemp map[string]interface{}) map[string]interface{} {
	var ok bool
	var mapRet map[string]interface{}
	mapRet, ok = deepcopy.Copy(mapTemp).(map[string]interface{})
	if ok {
		return mapRet
	}

	return nil
}
