package idgen

import (
	"fmt"
	"github.com/go-basic/uuid"
	"strings"
)

/**
 * @Author: qianjialin
 * @Date: 2021/2/25 13:55
 */
type IDGenerator struct {
}

func (g *IDGenerator) New() (string, error) {
	myUUID, err := uuid.GenerateUUID()
	if err != nil {
		fmt.Println(fmt.Sprintf("idgen New error=%s", err.Error()))
		return "", err
	}
	resID := strings.ReplaceAll(myUUID, "-", "")
	return resID, nil
}
