package lion

import (
	"edge-common-go/pkg/lion"
	"edge-common-go/pkg/utils"
	"fmt"
	"strings"
	"testing"
)

/**
 * Description:
 *
 * @author yang.zhang4
 * @date 2021/3/10
 */

//测试lion配置中加密
func TestEncryptValueOn(t *testing.T) {
	var thisFilePath = utils.GetFileName()
	var thisDir = getDirFromPath(thisFilePath)
	lion.DefaultLionConfigPath = thisDir + "/lion.properties"
	var defaultLionConfig MyLionConfig
	_ = defaultLionConfig.LoadLionConfig(&defaultLionConfig)

	var srcString = "aaaaaabbbbbbbcccccddde12643978rrryyyfff"
	var result1 []byte
	var err error
	result1, err = lion.EnValue([]byte(srcString))
	if err != nil {
		t.Fatalf("aes.EncryptValue error, %s", err)
	}

	var result2 []byte
	result2, err = lion.DeValue(result1)
	if err != nil {
		t.Fatalf("aes.DecryptValue error, %s", err)
	}

	fmt.Printf("src string=%s\n", srcString)
	fmt.Printf("encrypt string=%s\n", string(result1))
	fmt.Printf("decrypt string=%s\n", string(result2))

	if srcString != string(result2) {
		t.Fatalf("srcString(%s) != decrypt string(%s)\n", srcString, string(result2))
	}

}

//测试lion配置中不加密
func TestEncryptValueOff(t *testing.T) {
	var thisFilePath = utils.GetFileName()
	var thisDir = getDirFromPath(thisFilePath)
	lion.DefaultLionConfigPath = thisDir + "/lion1.properties"
	var defaultLionConfig MyLionConfig
	_ = defaultLionConfig.LoadLionConfig(&defaultLionConfig)

	var srcString = "aaaaaabbbbbbbcccccddde12643978rrryyyfff"
	var result1 []byte
	var err error

	result1, err = lion.EnValue([]byte(srcString))
	if err != nil {
		t.Fatalf("aes.EncryptValue error, %s", err)
	}

	var result2 []byte
	result2, err = lion.DeValue(result1)
	if err != nil {
		t.Fatalf("aes.DecryptValue error, %s", err)
	}

	fmt.Printf("src string=%s\n", srcString)
	fmt.Printf("encrypt string=%s\n", string(result1))
	fmt.Printf("decrypt string=%s\n", string(result2))

	if srcString != string(result1) {
		t.Fatalf("srcString(%s) != encrypt string(%s)\n", srcString, string(result1))
	}
	if srcString != string(result2) {
		t.Fatalf("srcString(%s) != decrypt string(%s)\n", srcString, string(result2))
	}
}

//测试lion配置中自动将需要解密的key解密
func TestLionAutoDecrypt(t *testing.T) {
	var thisFilePath = utils.GetFileName()
	var thisDir = getDirFromPath(thisFilePath)
	lion.DefaultLionConfigPath = thisDir + "/lion.properties"
	var defaultLionConfig MyLionConfig
	_ = defaultLionConfig.LoadLionConfig(&defaultLionConfig)

	var redisPExpect = "Envisi0nEdge@1"
	var mongoPExpect = "Envisi0nEdge@1"
	var mysqlPExpect = "Envisi0nEdge@1"
	var pushPExpect = "A,12345678"
	var xxxPExpect = "Envisi0nEdge@1"
	var aaaPExpect = "abcd"

	var redisP = lion.GetStringValue("edge.redis.password", "")
	var mongoP = lion.GetStringValue("edge.mongodb.password", "")
	var mysqlP = lion.GetStringValue("edge.mysql.password", "")
	var pushP = lion.GetStringValue("edge.pushgateway.pass_word", "")
	var xxxP = lion.GetStringValue("edge.need_en.xxx.p", "")
	var aaaP = lion.GetStringValue("edge.aaa.bbb", "")

	if redisP != redisPExpect {
		t.Fatalf("redisP(%s) != redisPExpect(%s)", redisP, redisPExpect)
	}

	if mongoP != mongoPExpect {
		t.Fatalf("mongoP(%s) != mongoPExpect(%s)", mongoP, mongoPExpect)
	}

	if mysqlP != mysqlPExpect {
		t.Fatalf("mysqlP(%s) != mysqlPExpect(%s)", mysqlP, mysqlPExpect)
	}

	if pushP != pushPExpect {
		t.Fatalf("pushP(%s) != pushPExpect(%s)", pushP, pushPExpect)
	}

	if xxxP != xxxPExpect {
		t.Fatalf("xxxP(%s) != xxxPExpect(%s)", xxxP, xxxPExpect)
	}

	if aaaP != aaaPExpect {
		t.Fatalf("aaaP(%s) != aaaPExpect(%s)", aaaP, aaaPExpect)
	}
}

//从文件路径获取目录
func getDirFromPath(filePath string) string {
	var sep = "/"
	if strings.Index(filePath, "/") >= 0 {
		sep = "/"
	} else {
		sep = "\\"
	}

	var stringArray []string
	stringArray = strings.Split(filePath, sep)

	if len(stringArray) <= 1 {
		return sep
	}

	var stringArrayNew = stringArray[:len(stringArray)-1]

	var dir string
	dir = strings.Join(stringArrayNew, sep)

	return dir
}
