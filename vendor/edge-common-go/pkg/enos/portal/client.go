/*
 * Copyright © 2020 江苏远景科创智能科技有限公司
 * All rights reserved.
 *
 * Filename : client.go
 * Description :
 *
 * Created by Laiyuan YANG at 2020/7/10 16:58:26
 *
 * http://www.envision-group.com/
 */

package portal

import (
	httpClient "edge-common-go/pkg/http-client"
)

type Client struct {
	client   httpClient.Client
	account  string
	password string
	token    string
}

func NewClient(host, account, password string) API {
	httpC := httpClient.Client{
		Host: host,
	}

	c := &Client{client: httpC}
	return c
}
