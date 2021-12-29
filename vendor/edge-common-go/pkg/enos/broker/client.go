/*
 * Copyright © 2020 江苏远景科创智能科技有限公司
 * All rights reserved.
 *
 * Filename : client.go
 * Description :
 *
 * Created by Laiyuan YANG at 2020/7/10 16:59:33
 *
 * http://www.envision-group.com/
 */

package broker

import (
	httpClient "edge-common-go/pkg/http-client"
	"fmt"
)

type Client struct {
	client       httpClient.Client
	productKey   string
	deviceKey    string
	deviceSecret string
}

func NewClient(host, productKey, deviceKey, deviceSecret string) Api {
	httpC := httpClient.Client{
		Host: host,
	}

	c := &Client{
		client:       httpC,
		productKey:   productKey,
		deviceKey:    deviceKey,
		deviceSecret: deviceSecret,
	}
	return c
}

func (c Client) GetOtaPkg(fileUri string) ([]byte, error) {
	return c.DownloadFile(fileUri, "ota")
}

//https://support-ppe1.envisioniot.com/docs/device-connection/zh_CN/latest/reference/http/downstream/download_file.html
func (c Client) DownloadFile(fileURI, category string) ([]byte, error) {
	sessionId, err := c.login(c.productKey, c.deviceKey, c.deviceSecret)
	if err != nil {
		return nil, err
	}
	query := make(map[string]string)
	query["sessionId"] = sessionId
	query["fileUri"] = fileURI
	query["category"] = category
	uri := fmt.Sprintf("/multipart/sys/%v/%v/file/download", c.productKey, c.deviceKey)
	return c.get(uri, query)
}
