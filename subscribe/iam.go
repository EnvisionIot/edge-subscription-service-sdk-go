package subscribe

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"github.com/edge-subscription-service-sdk-go/utils"
	"io/ioutil"
	"net"
	"net/http"
	"strconv"
	"strings"
	"sync"
	"time"
)

var (
	httpOnce sync.Once
	client   *http.Client
)

func getHttpClient() *http.Client {
	httpOnce.Do(func() {
		client = &http.Client{
			Timeout: time.Duration(utils.HttpTimeOut * time.Millisecond),
			Transport: &http.Transport{
				Proxy: http.ProxyFromEnvironment,
				DialContext: (&net.Dialer{
					Timeout:   30 * time.Second,
					KeepAlive: 30 * time.Second,
				}).DialContext,
				MaxIdleConns:          20,
				MaxIdleConnsPerHost:   20,
				IdleConnTimeout:       90 * time.Second,
				TLSHandshakeTimeout:   10 * time.Second,
				ExpectContinueTimeout: 1 * time.Second,
			},
		}
	})

	return client
}

func edgeAuthorityCheck(appKey, appSecret, subId string) (string, bool) {
	param := createAuthParam(appKey, appSecret, subId)
	data, err := json.Marshal(param)
	if err != nil {
		return err.Error(), false
	}
	body := bytes.NewBuffer(data)
	req, err := http.NewRequest(http.MethodPost, utils.IamAddr, body)
	if err != nil {
		return err.Error(), false
	}
	req.Header.Set("Content-Type", "application/json")
	client = getHttpClient()
	rsp, err := client.Do(req)
	if err != nil {
		return err.Error(), false
	}
	defer rsp.Body.Close()
	data, err = ioutil.ReadAll(rsp.Body)
	if err != nil {
		return err.Error(), false
	}
	if len(data) == 0 {
		if err != nil {
			return "iam no data", false
		}
	}
	var iam utils.IamRsp
	err = json.Unmarshal(data, &iam)
	if err != nil {
		return err.Error(), false
	}
	if iam.Status != 0 {
		return iam.Msg, false
	}

	return "", true
}

func createAuthParam(appKey, appSecret, subId string) utils.IamParam {
	param := utils.IamParam{
		AppKey:    appKey,
		Timestamp: time.Now().UnixNano() / 1000000,
		SubId:     subId,
		SkipSubId: false,
	}
	var buff strings.Builder
	buff.WriteString(appKey)
	stamp := strconv.FormatInt(param.Timestamp, 10)
	buff.WriteString(stamp)
	buff.WriteString(appSecret)
	buff.WriteString(subId)
	hashKey := buff.String()
	h := sha256.New()
	h.Write([]byte(hashKey))
	sum := h.Sum(nil)
	code := hex.EncodeToString(sum)
	sumStr := strings.ToLower(code)
	param.Encryption = sumStr

	return param
}
