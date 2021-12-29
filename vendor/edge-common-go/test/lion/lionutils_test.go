package lion

import (
	"edge-common-go/pkg/lion"
	"fmt"
	"testing"
)

/**
 * @Author: qianjialin
 * @Date: 2020/6/10 9:53
 */
var Example bool

type MyLionConfig struct {
	lion.Default_LionConfig
}

func (l *MyLionConfig) InitLionConfig() {
	Example = lion.GetBoolValue("edge.edge-archive.example.key", false)
}

func TestLionUtils_LoadLionConfig(t *testing.T) {
	var defaultLionConfig MyLionConfig
	defaultLionConfig.LoadLionConfig(&defaultLionConfig)
	fmt.Println("Example=", Example)
	fmt.Println(fmt.Sprintf("appCache= %+v", lion.GetAppLion()))
	for key, value := range lion.GetAppLion() {
		fmt.Println(fmt.Sprintf("%s=%v", key, value))
	}
	lionMap, lionKeySlice := lion.GetAppLionWithOrderedSlice()
	for _, key := range lionKeySlice {
		fmt.Println(fmt.Sprintf("%s=%v", key, lionMap[key]))
	}
	return
}
