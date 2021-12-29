/*
 * Copyright © 2020 江苏远景科创智能科技有限公司
 * All rights reserved.
 *
 * Filename : time.go
 * Description :
 *
 * Created by Laiyuan YANG at 2020/7/10 19:43:0
 *
 * http://www.envision-group.com/
 */

package utils

import (
	"strconv"
	"time"
)

func CurrentMilliSecond() int64 {
	return time.Now().Unix() * 1000
}

func CurrentMilliSecondStr() string {
	t := CurrentMilliSecond()
	return strconv.FormatInt(t, 10)
}

func TimeStamp() string {
	t := time.Now()
	str := t.Format("20060102_15_04_05")
	return str
}
