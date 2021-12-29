/*
 * Copyright © 2020 江苏远景科创智能科技有限公司
 * All rights reserved.
 *
 * Filename : requests.go
 * Description :
 *
 * Created by Laiyuan YANG at 2020/7/10 16:56:16
 *
 * http://www.envision-group.com/
 */

package apim

import (
	"edge-common-go/pkg/utils"
	"net/http"
	"sort"
)

//post http post
func (c *Client) Post(path string, rawBody []byte, params map[string]string, token string) ([]byte, error) {
	uri := c.basePath + path
	req, err := c.client.GenReq(http.MethodPost, uri, rawBody, params)
	if err != nil {
		return nil, err
	}
	// override the current token
	if token != "" {
		c.token = token
	}
	if c.token != "" {
		c.genHeader(req, params, string(rawBody))
	}
	c.setJsonHeader(req)
	return c.client.Request(req)
}

//get http get
func (c *Client) Get(path string, query map[string]string, token string) ([]byte, error) {
	uri := c.basePath + path
	req, err := c.client.GenReq(http.MethodGet, uri, nil, query)
	if err != nil {
		return nil, err
	}
	if token != "" {
		c.genHeader(req, query, "")
	}

	c.setJsonHeader(req)

	return c.client.Request(req)
}

func (c *Client) delete(path string, params map[string]string, token string) error {
	uri := c.basePath + path
	req, err := c.client.GenReq(http.MethodDelete, uri, nil, params)
	if err != nil {
		return err
	}
	if token != "" {
		c.genHeader(req, params, "")
	}
	c.setJsonHeader(req)

	_, err = c.client.Request(req)
	return err
}

func (c *Client) genHeader(req *http.Request, params map[string]string, body string) {
	c.setAccessToken(req, c.token)
	timestamp := utils.CurrentMilliSecondStr()
	c.setTimestamp(req, timestamp)
	sign := c.genSign(params, body, timestamp)
	c.setSign(req, sign)
}

func (c *Client) genSign(params map[string]string, body, timestamp string) string {
	content := makeContent(c.token, c.secretKey, timestamp, body, params)
	return utils.SignSha256(content)
}

func makeContent(token, appSecret, timestamp, body string, params map[string]string) string {
	var content string

	content += token
	var keys []string
	for key := range params {
		keys = append(keys, key)
	}
	sort.Strings(keys)

	for _, key := range keys {
		content += key + params[key]
	}
	content += body
	content += timestamp
	content += appSecret

	return content
}
