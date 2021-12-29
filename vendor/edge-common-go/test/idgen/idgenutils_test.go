package idgen

import (
	"edge-common-go/pkg/idgen"
	"fmt"
	"testing"
)

/**
 * @Author: qianjialin
 * @Date: 2021/2/25 14:25
 */
func Test_IDGenUtils(t *testing.T) {
	var myGenerator = idgen.IDGenerator{}
	myID, _ := myGenerator.New()
	myID2, _ := myGenerator.New()
	myID3, _ := myGenerator.New()
	fmt.Println(myID)
	fmt.Println(myID2)
	fmt.Println(myID3)
}
