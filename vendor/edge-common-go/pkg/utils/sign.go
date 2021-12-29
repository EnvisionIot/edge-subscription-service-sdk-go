/*
 * Copyright © 2020 江苏远景科创智能科技有限公司
 * All rights reserved.
 *
 * Filename : sign.go
 * Description :
 *
 * Created by Laiyuan YANG at 2020/7/10 19:43:33
 *
 * http://www.envision-group.com/
 */

package utils

import (
	"crypto/sha256"
	"encoding/hex"
	"strings"
)

func SignSha256(content string) string {
	h := sha256.New()
	h.Write([]byte(content))
	signedStr := hex.EncodeToString(h.Sum(nil))
	return strings.ToLower(signedStr)
}
