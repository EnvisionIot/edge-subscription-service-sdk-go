/*
 * Copyright © 2020 江苏远景科创智能科技有限公司
 * All rights reserved.
 *
 * Filename : api_test.go
 * Description :
 *
 * Created by Laiyuan YANG at 2020/7/14 13:51:14
 *
 * http://www.envision-group.com/
 */

package broker

import (
	httpClient "edge-common-go/pkg/http-client"
	"testing"
)

const (

//productKey   = "uNclxctt"
//deviceKey    = "trinity_matirx"
//deviceSecret = "W0jC75k2ZijO8KvlFEOT"
)

//open the bidirectinal authentication
func TestHttpsLogin(t *testing.T) {
	host := "https://iot-http-broker.beta-k8s-cn4.eniot.io"
	productKey := "uNclxctt"
	deviceKey := "trinity_matirx"
	deviceSecret := "W0jC75k2ZijO8KvlFEOT"

	httpC := httpClient.Client{
		Host:            host,
		RootCa:          "D:\\workspace\\envision\\edge-client-go\\build\\data\\apps\\config\\cert\\ca_root.crt",
		CertificateFile: "D:\\workspace\\envision\\edge-client-go\\build\\data\\apps\\config\\cert\\enos_edge.crt",
		PrivateKeyFile:  "D:\\workspace\\envision\\edge-client-go\\build\\data\\apps\\config\\cert\\enos_edge.key",
	}

	c := &Client{client: httpC}
	sessionId, err := c.login(productKey, deviceKey, deviceSecret)
	if err != nil {
		t.Error(err)
		return
	}

	t.Log("sessionId = ", sessionId)
}

func TestHttpLogin(t *testing.T) {
	host := "http://iot-http-broker.beta-k8s-cn4.eniot.io"
	productKey := "uNclxctt"
	deviceKey := "trinity_matirx"
	deviceSecret := "W0jC75k2ZijO8KvlFEOT"
	httpC := httpClient.Client{
		Host: host,
	}

	c := &Client{client: httpC}
	sessionId, err := c.login(productKey, deviceKey, deviceSecret)
	if err != nil {
		t.Error(err)
		return
	}

	t.Log("sessionId = ", sessionId)
}

func TestGetPkg(t *testing.T) {
	host := "http://iot-http-broker.beta-k8s-cn4.eniot.io"
	productKey := "uNclxctt"
	deviceKey := "trinity_matirx"
	deviceSecret := "W0jC75k2ZijO8KvlFEOT"

	c := NewClient(host, productKey, deviceKey, deviceSecret)

	data, err := c.GetOtaPkg("enos-connect://23f5ff270b800000.zip")
	if err != nil {
		t.Error(err)
		return
	}

	t.Log(data)

}
