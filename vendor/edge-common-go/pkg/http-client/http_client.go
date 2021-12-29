/*
 * Copyright © 2020 江苏远景科创智能科技有限公司
 * All rights reserved.
 *
 * Filename : http_client.go
 * Description :
 *
 * Created by Laiyuan YANG at 2020/7/10 17:4:18
 *
 * http://www.envision-group.com/
 */

package httpClient

import (
	"bytes"
	"crypto/tls"
	"crypto/x509"
	"errors"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

const (
	ApiTimeout = 30
)

var (
	timeoutErr = errors.New("time out")
)

type Client struct {
	Host            string
	RootCa          string
	CertificateFile string
	PrivateKeyFile  string
}

type Resp struct {
	Buffer []byte
	Err    error
}

func (hc *Client) GenReq(method, uri string, rawBody []byte, params map[string]string) (*http.Request, error) {
	url := hc.Host + uri
	body := bytes.NewReader(rawBody)

	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, err
	}
	q := req.URL.Query()
	for k, v := range params {
		q.Add(k, v)
	}
	req.URL.RawQuery = q.Encode()

	return req, nil
}

func (hc *Client) request(req *http.Request, ch chan<- Resp) {
	tlsCfg := &tls.Config{}
	if strings.HasPrefix(hc.Host, "https") && hc.RootCa != "" {
		certPool := x509.NewCertPool()
		pemCerts, err := ioutil.ReadFile(hc.RootCa)
		if err == nil {
			certPool.AppendCertsFromPEM(pemCerts)
		} else {
			ch <- Resp{Buffer: nil, Err: err}
			return
		}

		tlsCfg.RootCAs = certPool
		tlsCfg.InsecureSkipVerify = true
	}
	httpc := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: tlsCfg,
		},
	}
	res, err := httpc.Do(req)
	if err != nil {
		ch <- Resp{Buffer: nil, Err: err}
		return
	}
	defer func() {
		err := res.Body.Close()
		if err != nil {
			return
		}
	}()
	if res.StatusCode != http.StatusOK {
		ch <- Resp{Buffer: nil, Err: errors.New(res.Status)}
		return
	}

	respBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		ch <- Resp{Buffer: nil, Err: err}
		return
	}

	ch <- Resp{Buffer: respBody, Err: nil}
}

func (hc *Client) Request(req *http.Request) ([]byte, error) {
	var ch = make(chan Resp)
	go hc.request(req, ch)
	select {
	case <-time.After(time.Duration(ApiTimeout) * time.Second):
		return nil, timeoutErr
	case body := <-ch:
		return body.Buffer, body.Err
	}
}
