package lion

import (
	"edge-common-go/pkg/aes"
	"fmt"
	"io/ioutil"
	"strings"
)

/**
 * Description:这边函数需要在LoadLionConfig后调用
 *
 * @author yang.zhang4
 * @date 2021/3/10
 */

const (
	kSubDir         = "k"
	kFirstFileName  = "k_1"
	kSecondFileName = "k_2"
	kThirdFileName  = "k_3"
)

var (
	enOnOff    bool
	firstFile  string
	secondFile string
	thirdFile  string
)

var (
	LionKeysNeedEn = map[string]string{
		"edge.redis.password":        "1",
		"edge.mongodb.password":      "1",
		"edge.mysql.password":        "1",
		"edge.pushgateway.pass_word": "1",
		"edge.influxdb.password":     "1",
	}

	//只要lion的key中包含以下字符串，自动解密
	LionKeysMatchNeedEn = map[string]string{
		".need_en": "1",
		"_need_en": "1",
	}
)

func EnValue(valueIn []byte) ([]byte, error) {
	var p []byte
	p, _ = GetP()
	if !enOnOff {
		return valueIn, nil
	}
	if len(p) <= 0 {
		return valueIn, nil
	}

	return aes.AES256Encrypt(valueIn, p)
}

func DeValue(valueIn []byte) ([]byte, error) {
	var p []byte
	p, _ = GetP()
	if !enOnOff {
		return valueIn, nil
	}
	if len(p) <= 0 {
		return valueIn, nil
	}

	return aes.AES256Decrypt(valueIn, p)
}

//获取内置密码（先从文件读，读不到取默认）
func GetP() ([]byte, error) {
	var err error
	var files = getPFiles()

	var pBytes []byte
	pBytes, err = GetPFromPFiles(files)
	if err != nil {
		return aes.GetDefaultP(), fmt.Errorf("GetPFromPFiles, %s", err)
	}

	return pBytes, nil
}

//根据密码文件数组，获取实际密码
func GetPFromPFiles(files []string) ([]byte, error) {
	var err error
	var bufAll []byte
	for _, v := range files {
		var buf []byte
		buf, err = getFileFirstLine(v)
		if err != nil {
			return nil, fmt.Errorf("getFileFirstLine, %s", err)
		}
		bufAll = append(bufAll, buf...)
	}

	return bufAll, nil
}

//解密需要解密的lion
func decryptLion() {
	var p []byte
	p, _ = GetP()
	if !enOnOff {
		return
	}
	if len(p) <= 0 {
		return
	}

	for k, v := range lionCache {
		var ok bool
		//判断是否是需要解密的key
		ok = keyNeedDe(k)
		if !ok {
			continue
		}
		var vString string
		vString, ok = v.(string)
		if !ok {
			continue
		}
		if len(vString) <= 0 {
			continue
		}
		var byteArray []byte
		var err error
		byteArray, err = aes.AES256Decrypt([]byte(vString), p)
		if err != nil {
			fmt.Printf("aes.AES256Decrypt error, %s\n", err)
		} else {
			lionCache[k] = string(byteArray)
		}
	}
}

func initLionAes() {
	enOnOff = GetBoolValue("config-manger.encrpyt.on_off", true)
	firstFile = GetStringValue("config-manger.encrpyt.file.first", "")
	secondFile = GetStringValue("config-manger.encrpyt.file.second", "")
	thirdFile = GetStringValue("config-manger.encrpyt.file.third", "")
}

//判断一个lion的key需不需要解密
func keyNeedDe(keyIn string) bool {
	var ok bool
	//判断是否是需要解密的key
	_, ok = LionKeysNeedEn[keyIn]
	if ok {
		return true
	}

	for k := range LionKeysMatchNeedEn {
		if strings.Contains(keyIn, k) {
			return true
		}
	}

	return false
}

//从lion中获取密码文件的数组，如果lion中没配则取默认路径
func getPFiles() []string {
	var files []string
	if len(firstFile) > 0 {
		files = append(files, firstFile)
	}
	if len(secondFile) > 0 {
		files = append(files, secondFile)
	}
	if len(thirdFile) > 0 {
		files = append(files, thirdFile)
	}

	if len(files) <= 0 {
		return getDefaultPFiles()
	}

	return files
}

//获取默认lion配置的目录
func getDefaultLionDir() string {
	var defaultLionPath = DefaultLionConfigPath
	var sep = "/"
	if strings.Index(defaultLionPath, "/") >= 0 {
		sep = "/"
	} else {
		sep = "\\"
	}

	var stringArray []string
	stringArray = strings.Split(defaultLionPath, sep)

	if len(stringArray) <= 1 {
		return sep
	}

	var stringArrayNew = stringArray[:len(stringArray)-1]

	var lionDir string
	lionDir = strings.Join(stringArrayNew, sep)

	return lionDir
}

//获取默认密码文件的数组，默认文件位置为在lion文件所在文件夹下的k/k_1、k/k_2、k/k_3
func getDefaultPFiles() []string {
	var defaultLionDir = getDefaultLionDir()
	var sep = "/"
	if strings.Index(defaultLionDir, "/") >= 0 {
		sep = "/"
	} else {
		sep = "\\"
	}
	var files []string
	files = append(files, defaultLionDir+sep+kSubDir+sep+kFirstFileName)
	files = append(files, defaultLionDir+sep+kSubDir+sep+kSecondFileName)
	files = append(files, defaultLionDir+sep+kSubDir+sep+kThirdFileName)

	return files
}

//获取密码文件第一行
func getFileFirstLine(filePath string) ([]byte, error) {
	filePath = strings.Trim(filePath, "\r\n\t ")
	if len(filePath) <= 0 {
		return nil, fmt.Errorf("invalid file path, %s", filePath)
	}

	var buf []byte
	var err error
	buf, err = ioutil.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	var stringArray = strings.Split(string(buf), "\n")
	if len(stringArray) <= 0 {
		return nil, nil
	}

	var firstLine string
	firstLine = strings.Trim(stringArray[0], "\r\n\t ")
	return []byte(firstLine), nil
}
