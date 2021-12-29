/*
 * Copyright © 2020 江苏远景科创智能科技有限公司
 * All rights reserved.
 *
 * Filename : requtests_test.go
 * Description :
 *
 * Created by Laiyuan YANG at 2020/7/28 19:50:5
 *
 * http://www.envision-group.com/
 */

package apim

import (
	httpClient "edge-common-go/pkg/http-client"
	"testing"
)

var (
	host = "https://beta-apim-cn4.eniot.io"
	//edge_test
	accessKey = "635986bc-f6cb-4812552310a8-da8c-462a"
	secretKey = "552310a8-da8c-462a-a8ca-f60dbdd3a97a"

	caFileName = "ca_root.crt"
	//publicFileName = "cert.pem"
	privateFileName = "edge.key"
	//applyedFileName = "edge.pem"
	csrFileName = "enos_edge_csr"
	orgId       = "o15427722038191"
	productKey  = "bQBK85tM"
	deviceKey   = "test_ca_dev002"
	boxId       = "edge-client-go-common-name"
)

func TestSign(t *testing.T) {
	basePath := ""

	params := make(map[string]string)
	params["orgId"] = orgId
	params["productKey"] = productKey
	params["deviceKey"] = deviceKey
	params["action"] = "apply"
	httpC := httpClient.Client{Host: host}
	client := Client{
		client:    httpC,
		accessKey: accessKey,
		secretKey: secretKey,
		basePath:  basePath,
	}
	token, err := client.GetAccessToken()
	if err != nil {
		t.Error(err)
	}
	t.Log("token =", token)
	client.token = token

	//client, _ := NewClient(host, "", accessKey, secretKey)
	body := "{\n        \"csr\":  \"-----BEGIN CERTIFICATE REQUEST-----\\nMIICrTCCAZUCAQIwaDELMAkGA1UEBhMCQ04xETAPBgNVBAgMCFNoYW5naGFpMREw\\nDwYDVQQHDAhTaGFuZ2hhaTENMAsGA1UECgwERW5PUzESMBAGA1UECwwJRW5PUyBF\\nZGdlMRAwDgYDVQQDDAdmanh5Y2htMIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIB\\nCgKCAQEA3nv3CEl5adtyNamGwav95Ng/ZPysWfhusiZQaLJnA92KGZoXF1M430G+\\nr1Aq9z7RiMXlgiq9V3jVn0/alJMJFNOj1bNvjnBEIvJjuXA8XTcrv1cn9d7dtjwD\\nRYXGoznAHrwsyyKWHf5C/Q3ivHqKoJWA3Z/g2ZVML8YIrkOsWg5HJN/My0VRBcqi\\nJRkCvqoTJxJ2RP8gW4R0aCEfI89i/SY1RL7sEki1clu4wJJ1Mj5VwqLTE2IcqCxf\\naTvMxXhwQIsDVcUDrX8qdqx2fk7QKCM7qX/z4WiSMzXEgGwNAhGWa+X6F6KiqxUV\\nlMtNehDOqamL06PsX3DwhreC474zVQIDAQABoAAwDQYJKoZIhvcNAQELBQADggEB\\nANc0z90yTEmgVSwWPfttNN4089lcFqt03/tNQpD9Q3TpHaGS/PJSiJuz1F/xOtZl\\nNOagRiyeJPv5pqtl0ItP9rdZflHsZF3tYyvEni0J8xAtitqRZtkwhNR6+7JY/lcK\\nAE5RD+Z+4vC9BE5yZ99JDA3QKAEsf9MmHPqmMSXMudknp5zyup12FH6t4eHn6cym\\nKW1Hyid0wObDMaJ3dRyLtIO2wMXwcU0ClkXx+7qekx0oDayZCbxcF5zieQR1/iUM\\nOjcpzAKeQOfls+mgRG2lI7nvociR89NC8j/HaBoeA+Lu8kAxhvSFbaltqnM+5Hdr\\n/sKRrjs0MEtKrfVA8CO1MHs=\\n-----END CERTIFICATE REQUEST-----\\n\",\n        \"validDay\":     5,\n        \"issueAuthority\":       \"RSA\"\n}"
	timeStamp := "1595936431565"
	sign := client.genSign(params, body, timeStamp)
	t.Log("sign=", sign)
}
