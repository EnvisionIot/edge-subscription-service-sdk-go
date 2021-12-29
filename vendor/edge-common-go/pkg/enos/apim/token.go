/*
 * Copyright © 2020 江苏远景科创智能科技有限公司
 * All rights reserved.
 *
 * Filename : token.go
 * Description :
 *
 * Created by Laiyuan YANG at 2020/7/10 16:56:25
 *
 * http://www.envision-group.com/
 */

package apim

import (
	"edge-common-go/pkg/utils"
	"encoding/json"
	"errors"
)

const (
	GetTokenURI = "/apim-token-service/v2.0/token/get"
)

type TokenRequestBody struct {
	AppKey     string `json:"appKey"`
	Encryption string `json:"encryption"`
	Timestamp  string `json:"timestamp"`
}

type TokenInfo struct {
	AccessToken string `json:"accessToken"`
	Expire      int    `json:"expire"`
}
type TokenResponseBody struct {
	Status   int       `json:"status"`
	Message  string    `json:"msg"`
	Business string    `json:"business"`
	Token    TokenInfo `json:"data"`
}

func (c *Client) GetAccessToken() (string, error) {
	body := TokenRequestBody{AppKey: c.accessKey}
	ts := utils.CurrentMilliSecondStr()
	body.Timestamp = ts
	body.Encryption = signTokenReq(c.accessKey, c.secretKey, ts)

	rawBody, err := json.Marshal(body)
	if err != nil {
		return "", err
	}
	resp, err := c.Post(GetTokenURI, rawBody, nil, "")
	if err != nil {
		return "", err
	}

	var tokenResp TokenResponseBody
	err = json.Unmarshal(resp, &tokenResp)
	if err != nil {
		return "", err
	}

	if tokenResp.Status != 0 {
		return "", errors.New("service side error " + tokenResp.Message)
	}

	return tokenResp.Token.AccessToken, nil
}

func signTokenReq(appKey, appSecret, timestamp string) string {
	content := appKey + timestamp + appSecret
	return utils.SignSha256(content)
}
