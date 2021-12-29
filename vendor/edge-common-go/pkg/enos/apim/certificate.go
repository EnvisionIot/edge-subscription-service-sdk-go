/*
 * Copyright © 2020 江苏远景科创智能科技有限公司
 * All rights reserved.
 *
 * Filename : certificate.go
 * Description :
 *
 * Created by Laiyuan YANG at 2020/7/10 16:55:57
 *
 * http://www.envision-group.com/
 */

package apim

import (
	"encoding/json"
	"errors"
)

const (
	ApplyCertificateURI = "/connect-service/v2.0/certificates"
	DefaultValidDay     = 730
	DeFaultAuthority    = "RSA"
)

type applyRequestBody struct {
	CSR            string `json:"csr"`
	ValidDay       int    `json:"validDay"`
	IssueAuthority string `json:"issueAuthority"`
}

type CertInfo struct {
	CertChainURL   string `json:"certChainURL"`
	CaCert         string `json:"caCert"`
	Cert           string `json:"cert"`
	CertSN         string `json:"certSN"`
	IssueAuthority string `json:"issueAuthority"`
}
type applyResponseBody struct {
	RespBase
	Data CertInfo `json:"data"`
}

func (c *Client) ApplyCertificate(orgId, productKey, deviceKey, csr string) (CertInfo, error) {
	var certInfo CertInfo
	params := make(map[string]string)
	params["orgId"] = orgId
	params["productKey"] = productKey
	params["deviceKey"] = deviceKey
	params["action"] = "apply"

	req := applyRequestBody{CSR: csr, ValidDay: DefaultValidDay, IssueAuthority: DeFaultAuthority}
	rawBody, err := json.Marshal(req)
	if err != nil {
		return certInfo, err
	}

	data, err := c.Post(ApplyCertificateURI, rawBody, params, c.token)
	if err != nil {
		return certInfo, err
	}
	var resp applyResponseBody
	err = json.Unmarshal(data, &resp)
	if err != nil {
		return certInfo, err
	}

	if resp.Code != 0 {
		return certInfo, errors.New(resp.Message)
	}

	return resp.Data, nil
}

func (c *Client) RenewCertificate(orgId, productKey, deviceKey string, certSn int) (CertInfo, error) {
	var certInfo CertInfo
	params := make(map[string]string)
	params["orgId"] = orgId
	params["productKey"] = productKey
	params["deviceKey"] = deviceKey
	params["action"] = "renew"

	type RenewBody struct {
		CertSn int `json:"certSn"`
	}

	body := RenewBody{CertSn: certSn}
	rawBody, _ := json.Marshal(body)
	data, err := c.Post(ApplyCertificateURI, rawBody, params, c.token)
	if err != nil {
		return certInfo, err
	}
	var resp applyResponseBody
	err = json.Unmarshal(data, &resp)
	if err != nil {
		return certInfo, err
	}

	if resp.Code != 0 {
		return certInfo, errors.New(resp.Message)
	}

	return resp.Data, nil
}

func (c *Client) RevokeCertificate() {

}

func (c *Client) ListCertificate() {

}
