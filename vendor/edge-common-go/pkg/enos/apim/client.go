/*
 * Copyright © 2020 江苏远景科创智能科技有限公司
 * All rights reserved.
 *
 * Filename : client.go
 * Description :
 *
 * Created by Laiyuan YANG at 2020/7/10 16:56:2
 *
 * http://www.envision-group.com/
 */

package apim

import (
	httpClient "edge-common-go/pkg/http-client"
)

type Client struct {
	client    httpClient.Client
	accessKey string
	secretKey string
	basePath  string
	token     string
}

//GetClientInstance 新建一个HTTP客户端
func NewClient(host, basePath, accessKey, secretKey string) (API, error) {
	httpC := httpClient.Client{Host: host}
	client := Client{client: httpC, basePath: basePath, accessKey: accessKey, secretKey: secretKey}

	token, err := client.GetAccessToken()
	if err != nil {
		return nil, err
	}

	client.token = token

	return &client, nil
}

func (c *Client) SetRootCa(ca string) {
	c.client.RootCa = ca
}
