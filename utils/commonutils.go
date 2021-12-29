package utils

import (
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"time"
)

var (
	PROGRAM_NAME = "edge-dataservice"
)

//判断文件夹是否存在
func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
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

//获取短文件名，就是从common.PROGRAM_NAME往后的文件路径
func GetShortFileName() string {
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

	var index = -1
	for k, v := range stringArray {
		if v == PROGRAM_NAME {
			index = k
			break
		}
	}

	if index < 0 {
		return file
	}

	var stringArrayNew = stringArray[index+1:]

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

func GetLocalTimeStrByUnix(timeSec int64, timeNsec int64) string {
	var tm time.Time
	var timeTemplate = "2006-01-02 15:04:05.999"
	var timeStr string
	tm = time.Unix(timeSec, timeNsec)
	timeStr = tm.Format(timeTemplate)

	return timeStr
}

//获取baseDir
func GetBaseDir() string {
	ex, err := os.Executable()
	if err != nil {
		os.Exit(1)
	}
	baseDir := strings.TrimSuffix(filepath.Dir(ex), "/bin")
	baseDir = strings.TrimSuffix(baseDir, "\\bin") // for windows
	return baseDir
}
