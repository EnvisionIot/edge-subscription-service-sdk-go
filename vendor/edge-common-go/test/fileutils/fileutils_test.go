package fileutils

import (
	"edge-common-go/pkg/utils"
	"testing"
)

/**
 * Description:
 *
 * @author yang.zhang4
 * @date 2020/12/3
 */

func TestCopy(t *testing.T) {
	_ = utils.Copy("/aaa/bbb", "/aaa/ccc")
}
