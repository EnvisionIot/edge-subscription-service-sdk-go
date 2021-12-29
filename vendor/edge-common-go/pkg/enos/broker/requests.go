/*
 * Copyright © 2020 江苏远景科创智能科技有限公司
 * All rights reserved.
 *
 * Filename : requests.go
 * Description :
 *
 * Created by Laiyuan YANG at 2020/7/10 20:54:36
 *
 * http://www.envision-group.com/
 */

package broker

import (
	"edge-common-go/pkg/utils"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"sort"
)

const (
	signMethod      = "sha256"
	defaultLifeTime = 60000
	loginUri        = "/auth/%v/%v"
	success         = 200
)

type LoginReq struct {
	SignMethod string `json:"signMethod"`
	LifeTime   int64  `json:"lifetime"`
	Sign       string `json:"sign"`
}

type ResponseBody struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type LoginRespData struct {
	SessionId string `json:"sessionId"`
}

type LoginResp struct {
	ResponseBody
	Data LoginRespData `json:"data"`
}

func (c Client) get(uri string, query map[string]string) ([]byte, error) {
	req, err := c.client.GenReq(http.MethodGet, uri, nil, query)
	if err != nil {
		return nil, err
	}
	c.setJsonHeader(req)

	return c.client.Request(req)
}

func (c Client) setJsonHeader(req *http.Request) {
	req.Header.Add("Content-Type", "application/json;charset=UTF-8")
}

func (c Client) post(uri string, query map[string]string, rawBody []byte) ([]byte, error) {
	req, err := c.client.GenReq(http.MethodPost, uri, rawBody, query)
	if err != nil {
		return nil, err
	}
	c.setJsonHeader(req)

	return c.client.Request(req)
}

//now we only support http. and product not open the bidirectinal authentication
func (c Client) login(productKey, deviceKey, deviceSecret string) (string, error) {
	uri := fmt.Sprintf(loginUri, productKey, deviceKey)
	req := LoginReq{}
	req.SignMethod = signMethod
	req.LifeTime = defaultLifeTime
	req.Sign = c.genSign(productKey, deviceKey, deviceSecret)

	rawBody, _ := json.Marshal(req)
	data, err := c.post(uri, nil, rawBody)
	if err != nil {
		return "", err
	}

	var loginResp LoginResp
	err = json.Unmarshal(data, &loginResp)
	if err != nil {
		return "", err
	}

	if loginResp.Code != success {
		return "", errors.New(loginResp.Message)
	}

	return loginResp.Data.SessionId, nil
}

func (c Client) genSign(productKey, deviceKey, deviceSecret string) string {
	params := make(map[string]interface{})
	params["deviceKey"] = deviceKey
	params["productKey"] = productKey
	params["signMethod"] = signMethod
	params["lifetime"] = defaultLifeTime
	content := makeContent(deviceSecret, params)
	return utils.SignSha256(content)
}

//sign具体加密方式：(键值对的字典排序拼接，最后加deviceSecret)
func makeContent(deviceSecret string, params map[string]interface{}) string {
	var content string

	var keys []string
	for key := range params {
		keys = append(keys, key)
	}
	sort.Strings(keys)

	for _, key := range keys {
		content += fmt.Sprintf("%v%v", key, params[key])
	}

	content += deviceSecret

	return content
}
