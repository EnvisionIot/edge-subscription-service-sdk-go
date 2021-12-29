package httpclient

/**
 * @Author: qianjialin
 * @Date: 2020/7/10 10:56
 */
import (
	"bytes"
	"edge-common-go/pkg/idgen"
	"edge-common-go/pkg/lion"
	"edge-common-go/pkg/log"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/spf13/cast"
	"io/ioutil"
	"net"
	"net/http"
	"net/url"
	"sync"
	"time"
)

var (
	DefaultHeaderMap = map[string]string{
		"Content-Type": "application/json; charset=utf-8",
	}

	//map[timeoutMs]*http.Client
	defaultClientMap = make(map[int64]*http.Client)
	//操作defaultClientMap的锁
	defaultClientMapMutex sync.RWMutex
)

//For Sonar-Scan
var postUrlParam = "postUrl="
var paramsParam = "params="

var HttpClientGlobal *HttpClient

//依赖lion获取打印开关，依赖log获取日志类
type HttpClient struct {
	//log      *log.LogAa
	traceLog *log.LogAa
	//map[timeoutMs]*http.Client
	defaultClientMap map[int64]*http.Client
	//操作defaultClientMap的锁
	defaultClientMapMutex sync.RWMutex
}

func New(logPath string) *HttpClient {
	if HttpClientGlobal != nil {
		return HttpClientGlobal
	}
	var httpClient = HttpClient{}
	var logPathDefault = "/trace/trace.log"
	if logPath != "" {
		logPathDefault = logPath
	}
	httpClient.traceLog, _ = log.GetLogAa(logPathDefault)
	httpClient.defaultClientMap = make(map[int64]*http.Client)

	HttpClientGlobal = &httpClient
	return HttpClientGlobal
}

func (h *HttpClient) HttpDefaultGetWithTimeoutTrace(url string, timeoutMs int64, traceId string) ([]byte, error) {
	needTrace := lion.GetBoolValue("edge.trace.need", false)
	if needTrace {
		if traceId != "" {
			DefaultHeaderMap["X-Trace-ID"] = traceId
		} else {
			var generator = idgen.IDGenerator{}
			traceId, _ = generator.New()
			DefaultHeaderMap["X-Trace-ID"] = traceId
		}
		h.traceLog.Infof("http get url=%s,timeoutMs=%s,traceId=%s", url, timeoutMs, traceId)
	}
	return HttpRequest(nil, http.MethodGet, url, nil, DefaultHeaderMap, timeoutMs)
}

func (h *HttpClient) HttpDefaultPostWithTimeoutTrace(url string, requestBodyContent interface{}, timeoutMs int64, traceId string) ([]byte, error) {
	needTrace := lion.GetBoolValue("edge.trace.need", false)
	if needTrace {
		if traceId != "" {
			DefaultHeaderMap["X-Trace-ID"] = traceId
		} else {
			var generator = idgen.IDGenerator{}
			traceId, _ := generator.New()
			DefaultHeaderMap["X-Trace-ID"] = traceId
		}
		h.traceLog.Infof("http post url=%s,timeoutMs=%s,traceId=%s", url, timeoutMs, traceId)
	}
	return HttpRequest(nil, http.MethodPost, url, requestBodyContent, DefaultHeaderMap, timeoutMs)
}

func (h *HttpClient) HttpDefaultPutWithTimeoutTrace(url string, requestBodyContent interface{}, timeoutMs int64, traceId string) ([]byte, error) {
	needTrace := lion.GetBoolValue("edge.trace.need", false)
	if needTrace {
		if traceId != "" {
			DefaultHeaderMap["X-Trace-ID"] = traceId
		} else {
			var generator = idgen.IDGenerator{}
			traceId, _ := generator.New()
			DefaultHeaderMap["X-Trace-ID"] = traceId
		}
		h.traceLog.Infof("http put url=%s,timeoutMs=%s,traceId=%s", url, timeoutMs, traceId)
	}
	return HttpRequest(nil, http.MethodPut, url, requestBodyContent, DefaultHeaderMap, timeoutMs)
}

func Get(getUrl string) (response string, err error) {
	u, _ := url.Parse(getUrl)
	q := u.Query()
	u.RawQuery = q.Encode() //urlencode
	resp, err := http.Get(u.String())
	if err != nil {
		msg := fmt.Sprintf("Get error url =", getUrl, "res =", resp, "err =", err)
		return "", errors.New(msg)
	}
	defer resp.Body.Close()
	if res, err := ioutil.ReadAll(resp.Body); err != nil {
		msg := fmt.Sprintf("Get resp.Body error url =", getUrl, "res =", res, "err =", err)
		return "", errors.New(msg)
	} else {
		if resp.StatusCode != http.StatusOK {
			return string(res), fmt.Errorf("http status code=%d error", resp.StatusCode)
		}
		return string(res), nil
	}
}

func DoFormPost(postUrl string, paramsMap interface{}, contentType string, encode string) (response string, err error) {
	return DoHttpPostWithFullPara(postUrl, paramsMap, contentType, encode)
}

func DoJsonPost(postUrl string, paramsMap interface{}, encode string) (response string, err error) {
	return DoHttpPostWithFullPara(postUrl, paramsMap, "application/json", encode)
}

func DoHttpPostWithFullPara(postUrl string, paramsMap interface{}, contentType string, encode string) (response string, err error) {
	var m = map[string]interface{}{}

	switch v := paramsMap.(type) {
	case map[interface{}]interface{}:
		for k, val := range v {
			m[cast.ToString(k)] = val
		}
		paramsMap = m
	case map[string]interface{}, map[string]string:
	default:
		return "", fmt.Errorf("unsupported type:%T", paramsMap)
	}

	params, errs := json.Marshal(paramsMap)
	if errs != nil {
		msg := fmt.Sprintf("DoHttpPostWithFullPara paramsMap error =", errs, postUrlParam, postUrl, paramsParam, paramsMap)
		return "", errors.New(msg)
	}
	resp, err := http.Post(postUrl, contentType+";charset="+encode, bytes.NewBuffer(params))
	if err != nil {
		msg := fmt.Sprintf("DoHttpPostWithFullPara error =", err, postUrlParam, postUrl, paramsParam, paramsMap)
		return "", errors.New(msg)
	}
	defer resp.Body.Close()
	if res, err := ioutil.ReadAll(resp.Body); err != nil {
		msg := fmt.Sprintf("DoHttpPostWithFullPara resp.Body error =", err, postUrlParam, postUrl, paramsParam, paramsMap)
		return "", errors.New(msg)
	} else {
		fmt.Printf("DoHttpPostWithFullPara res =", string(res), postUrlParam, postUrl, "paramsMap=", string(params))
		return string(res), nil
	}
}

//默认的带超时的get
func HttpDefaultGetWithTimeout(url string, timeoutMs int64) ([]byte, error) {
	return HttpRequest(nil, http.MethodGet, url, nil, DefaultHeaderMap, timeoutMs)
}

//默认的带超时的post
func HttpDefaultPostWithTimeout(url string, requestBodyContent interface{}, timeoutMs int64) ([]byte, error) {
	return HttpRequest(nil, http.MethodPost, url, requestBodyContent, DefaultHeaderMap, timeoutMs)
}

//默认的带超时的put
func HttpDefaultPutWithTimeout(url string, requestBodyContent interface{}, timeoutMs int64) ([]byte, error) {
	return HttpRequest(nil, http.MethodPut, url, requestBodyContent, DefaultHeaderMap, timeoutMs)
}

//requestBodyContent一般是一个map，httpClient为null时，timeoutMs生效，不为null时，使用httpClient中的timeout
func HttpRequest(httpClient *http.Client, method string, url string, requestBodyContent interface{}, headerMap map[string]string, timeoutMs int64) ([]byte, error) {
	var requestBodyByteArray []byte
	var err error
	if requestBodyContent != nil {
		requestBodyByteArray, err = json.Marshal(requestBodyContent)
		if err != nil {
			return nil, fmt.Errorf("json.Marshal requestBodyContent error, %s", err)
		}
	}

	var requestBody *bytes.Buffer
	requestBody = bytes.NewBuffer(requestBodyByteArray)

	var req *http.Request
	req, err = http.NewRequest(method, url, requestBody)
	if err != nil {
		return nil, fmt.Errorf("http.NewRequest error, %s", err)
	}

	for key, value := range headerMap {
		req.Header.Add(key, value)
	}

	var client *http.Client
	if httpClient != nil {
		client = httpClient
	} else {
		client, _ = getDefaultHttpClientByTimeOutMs(timeoutMs)
	}

	var res *http.Response
	res, err = client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("client.Do error, %s", err)
	}
	defer func() {
		_ = res.Body.Close()
	}()

	var responseBodyByteArray []byte
	responseBodyByteArray, err = ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf("ioutil.ReadAll error, %s", err)
	}

	if res.StatusCode != http.StatusOK {
		return responseBodyByteArray, fmt.Errorf("http status code=%d error", res.StatusCode)
	}

	return responseBodyByteArray, nil
}

//获取推荐配置的*http.Client，相比于原生http库的默认配置，主要修改了MaxIdleConnsPerHost这个参数，防止多线程请求同一个ip的http请求产生过多的timewait
func GetRecommendHttpClient(httpTimeoutMs int64) (*http.Client, error) {
	var client http.Client
	client = http.Client{
		Timeout: time.Duration(httpTimeoutMs) * time.Millisecond,
		Transport: &http.Transport{
			Proxy: http.ProxyFromEnvironment,
			DialContext: (&net.Dialer{
				Timeout:   30 * time.Second,
				KeepAlive: 30 * time.Second,
				DualStack: true,
			}).DialContext,
			MaxIdleConns:          100,
			MaxIdleConnsPerHost:   100, //防止多线程请求同一个ip的http请求产生过多的timewait
			IdleConnTimeout:       90 * time.Second,
			TLSHandshakeTimeout:   10 * time.Second,
			ExpectContinueTimeout: 1 * time.Second,
		},
	}
	return &client, nil
}

func getDefaultHttpClientByTimeOutMs(timeoutMs int64) (*http.Client, error) {
	var httpClient *http.Client
	var ok bool
	defaultClientMapMutex.RLock()
	httpClient, ok = defaultClientMap[timeoutMs]
	if ok {
		defaultClientMapMutex.RUnlock()
		return httpClient, nil
	}
	defaultClientMapMutex.RUnlock()

	defaultClientMapMutex.Lock()
	defer defaultClientMapMutex.Unlock()
	httpClient, ok = defaultClientMap[timeoutMs]
	if ok {
		return httpClient, nil
	}
	httpClient, _ = GetRecommendHttpClient(timeoutMs)
	defaultClientMap[timeoutMs] = httpClient

	return httpClient, nil
}
