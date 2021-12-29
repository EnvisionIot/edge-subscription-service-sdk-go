/*
 * Copyright © 2020 江苏远景科创智能科技有限公司
 * All rights reserved.
 *
 * Filename : header.go
 * Description :
 *
 * Created by Laiyuan YANG at 2020/7/10 16:56:8
 *
 * http://www.envision-group.com/
 */

package apim

import (
	"net/http"
)

const (
	AccessTokenHeader = "apim-accesstoken"
	SignHeader        = "apim-signature"
	TimestampHeader   = "apim-timestamp"
)

func (c Client) setAccessToken(req *http.Request, token string) {
	req.Header.Add(AccessTokenHeader, token)
}

func (c Client) setSign(req *http.Request, sign string) {
	req.Header.Add(SignHeader, sign)
}

func (c Client) setTimestamp(req *http.Request, timestamp string) {
	req.Header.Add(TimestampHeader, timestamp)
}

func (c Client) setJsonHeader(req *http.Request) {
	req.Header.Add("Content-Type", "application/json;charset=UTF-8")
}
