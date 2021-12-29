/*
 * Copyright © 2020 江苏远景科创智能科技有限公司
 * All rights reserved.
 *
 * Filename : interface.go
 * Description :
 *
 * Created by Laiyuan YANG at 2020/7/10 16:56:12
 *
 * http://www.envision-group.com/
 */

package apim

type API interface {
	SetRootCa(ca string)
	GetAccessToken() (string, error)
	ApplyCertificate(orgId, productKey, deviceKey, csr string) (CertInfo, error)
	RenewCertificate(orgId, productKey, deviceKey string, certSn int) (CertInfo, error)
	Post(path string, rawBody []byte, params map[string]string, token string) ([]byte, error)
	Get(path string, query map[string]string, token string) ([]byte, error)
}
