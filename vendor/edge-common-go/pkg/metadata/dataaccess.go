package metadata

import (
	"edge-common-go/pkg/httpclient"
	"encoding/json"
	"errors"
	"fmt"
)

/**
 * Description:获取一些设备接入相关的元数据，如orgId，mapping，模板，collectDevice等信息
 *
 * @author yang.zhang4
 * @date 2020/12/3
 */

const (
	STATE_DISCONNECTED = 0
	STATE_CONNECTED    = 1

	MODEL_ID_PATH       = "modelIdPath"
	HAS_QUALITY         = "hasQuality"
	SIGNAL_TYPE         = "signalType"
	DATA_INNER_TYPE     = "dataInnerType"
	USING_OME_TIMESTAMP = "usingOMETimestamp"

	DEVICE_LOGICAL_ID     = "deviceLogicalId"
	TEMPLATE_TYPE_NORMAL  = 0
	TEMPLATE_TYPE_FORWARD = 1
	TEMPLATE_TYPE_DYNAMIC = 2

	MAPPING_CONTROL_SET             = "CONTROL_SET"
	MAPPING_CONTROL_SET_1           = "CONTROL_SET_1"
	MAPPING_NO_MAPPING              = "NO_MAPPING"
	MAPPING_INVALID                 = "INVALID"
	MAPPING_EQUAL                   = "EQUAL"
	MAPPING_MULTICHANNEL            = "MULTICHANNEL"
	MAPPING_MULTICHANNEL_SHORTTOINT = "MULTICHANNEL_SHORTTOINT"
	MAPPING_MULTIBIT                = "MULTIBIT"
	MAPPING_SUM                     = "SUM"
	MAPPING_PRODUCT                 = "PRODUCT"
	MAPPING_CROSS_PRODUCT           = "CROSS_PRODUCT"
	MAPPING_RATIO                   = "RATIO"
	MAPPING_LOGICAL_OR              = "LOGICAL_OR"
	MAPPING_RATIO_AGAINST_SUM       = "RATIO_AGAINST_SUM"
	MAPPING_BIT_N                   = "BIT_N"
	MAPPING_BITS_M_TO_N             = "BITS_M_TO_N"
	MAPPING_IF_EQUAL                = "IF_EQUAL"
	MAPPING_IF_RANGE                = "IF_RANGE"
	MAPPING_ENUM_N                  = "ENUM_N"
	CONNECT_STATE                   = "CONNECT_STATE"
	EVENT_EQUAL                     = "EVENT_EQUAL"
)

type CollectDevice struct {
	Id           string            `json:"id"`
	AssetId      string            `json:"assetId"`
	TemplateId   int64             `json:"templateId"`
	CollectId    string            `json:"collectId"`
	ProductKey   string            `json:"productKey"`
	DeviceKey    string            `json:"deviceKey"`
	DeviceSecret string            `json:"deviceSecret"`
	BoxId        string            `json:"boxId"`
	Attributes   map[string]string `json:"attributes"`
	CreateDate   int64             `json:"createDate"`
	CollectType  int64             `json:"collectType"`
}

type CollectDeviceResponse struct {
	Code int             `json:"code"`
	Data []CollectDevice `json:"data"`
}

type Template struct {
	Id           int64             `json:"id"`
	OrgId        string            `json:"orgId"`
	ModelId      string            `json:"modelId"`
	ProtocolId   int64             `json:"protocolId"`
	TemplateName string            `json:"templateName"`
	TemplateType int64             `json:"templateType"`
	Brand        string            `json:"brand"`
	DeviceType   string            `json:"deviceType"`
	InnerVersion string            `json:"innerVersion"`
	ScriptIdList []string          `json:"scriptIdList"`
	PointCsvPath string            `json:"pointCsvPath"`
	Creator      string            `json:"creator"`
	CreateDate   int64             `json:"createDate"`
	UpdateData   int64             `json:"updateData"`
	Attributes   map[string]string `json:"attributes"`
}

type TemplateResponse struct {
	Code int        `json:"code"`
	Data []Template `json:"data"`
}

type Mapping struct {
	Id               int64             `json:"id"`
	TemplateId       int64             `json:"templateId"`
	ModelId          string            `json:"modelId"`
	FeatureName      string            `json:"featureName"`
	FeatureType      string            `json:"featureType"`
	FeatureValueType string            `json:"featureValueType"`
	Control          bool              `json:"control"`
	Multiple         bool              `json:"multiple"`
	UploadType       string            `json:"uploadType"`
	UploadInterval   int               `json:"uploadInterval"`
	MappingRule      MappingRule       `json:"mappingRule"`
	CreateDate       int64             `json:"createDate"`
	UpdateDate       int64             `json:"updateDate"`
	Attributes       map[string]string `json:"attributes"`
}

type MappingRule struct {
	MappingFxStr string   `json:"mappingFxStr"`
	Arguments    []string `json:"arguments"`
	Operands     []string `json:"operands"`
}

type MappingResponse struct {
	Code int       `json:"code"`
	Data []Mapping `json:"data"`
}

type OrgIdResponse struct {
	Code   int       `json:"code"`
	Status int       `json:"status"`
	Msg    string    `json:"msg"`
	SubMsg string    `json:"subMsg"`
	Data   OrgIdData `json:"data"`
}

type OrgIdData struct {
	OrgId string `json:"orgId"`
}

type DataAccessBase struct {
	OrgId                             string
	CollectDeviceIdToCollectDeviceMap map[string]CollectDevice
	TemplateIdToTemplateMap           map[int64]Template
	MappingIdToMappingMap             map[int64]Mapping
}

type DataAccessAdvance struct {
	//map[CollectDeviceId]CollectId
	CollectDeviceIdToCollectId map[string]string
	//map[CollectDeviceId]TemplateId
	CollectDeviceIdToTemplateIdMap map[string]int64
	//map[AssetId]map[CollectDeviceId]"1"
	AssetIdToCollectDeviceIdMap map[string]map[string]string
	//map[TemplateId]map[MappingId]"1"
	TemplateIdToMappingIdMap map[int64]map[int64]string
	//map[TemplateId]map[模型点标识符]map[deviceLogicalId]MappingId,这里只考虑采集mapping
	TemplateIdFeatureNameDeviceLogicalIdToMappingIdNonControlMap map[int64]map[string]map[string]int64
	//map[TemplateId]map[模型点标识符]map[deviceLogicalId]MappingId,这里只考虑控制mapping
	TemplateIdFeatureNameDeviceLogicalIdToMappingIdControlMap map[int64]map[string]map[string]int64
	//map[TemplateId]map[DevicePointName][]MappingId，如map[12345]map[ai.1][12345,12346,12347]
	TemplateIdToDevicePointNameMappingMap map[int64]map[string][]int64
}

type DataAccess struct {
	Base    *DataAccessBase
	Advance *DataAccessAdvance
}

type DataAccessConfig struct {
	WebOpTimeoutMs int64

	GetOrgIdUrl          string
	GetCollectDevicesUrl string
	GetTemplatesUrl      string
	GetMappingsUrl       string

	//LogArray []*log.LogAa
}

type DataAccessWork struct {
	Config DataAccessConfig
}

func CheckDataAccessIsValid(dataAccess *DataAccess) bool {
	if dataAccess == nil || dataAccess.Base == nil || dataAccess.Advance == nil {
		return false
	}

	return true
}

//从assetId获取collectDeviceIdMap
func GetCollectDeviceIdMapByAssetId(dataAccess *DataAccess, assetId string) (map[string]string, error) {
	if !CheckDataAccessIsValid(dataAccess) {
		return nil, ERROR_RESOURCE_IS_NIL
	}

	var collectDeviceIdMap map[string]string
	var ok bool
	collectDeviceIdMap, ok = dataAccess.Advance.AssetIdToCollectDeviceIdMap[assetId]
	if !ok {
		return nil, ERROR_NOT_FOUND
	}
	return collectDeviceIdMap, nil
}

//请求orgId
func (f *DataAccessWork) RequestOrgId() (string, error) {
	var valueRet string
	var res []byte
	var err error

	//构造Request请求
	res, err = httpclient.HttpDefaultGetWithTimeout(f.Config.GetOrgIdUrl, f.Config.WebOpTimeoutMs)
	if err != nil {
		return "", fmt.Errorf("httpDefaultGetWithTimeout, %s", err)
	}

	//解析OrgId信息
	var orgIdResponse OrgIdResponse
	err = json.Unmarshal(res, &orgIdResponse)
	if err != nil {
		return "", fmt.Errorf("json.Unmarshal, %s", err)
	}

	if orgIdResponse.Status != 0 || orgIdResponse.Code != 0 {
		return "", fmt.Errorf("code or status in response error, response=%s", string(res))
	}

	if len(orgIdResponse.Data.OrgId) <= 0 {
		return "", errors.New("orgId is empty")
	}

	valueRet = orgIdResponse.Data.OrgId

	return valueRet, nil
}

//请求orgId直到成功为止
//func (f *DataAccessWork) RequestOrgIdUntilSuccess() string {
//	var ticker *time.Ticker
//	ticker = time.NewTicker(time.Second * 3)
//
//	var err error
//	var valueRet string
//
//	for {
//		select {
//		case <-ticker.C:
//			valueRet, err = f.RequestOrgId()
//			if err == nil {
//				ticker.Stop()
//				return valueRet
//			} else {
//				log.MultiLogErrorf(f.Config.LogArray, "requestOrgId error, %s", err)
//			}
//		}
//	}
//}

//请求CollectDevice
func (f *DataAccessWork) RequestCollectDeviceData(orgId string) (map[string]CollectDevice, error) {
	var valueRet map[string]CollectDevice
	var res []byte
	var err error

	valueRet = make(map[string]CollectDevice)

	//构造Request请求，获取CollectDevice信息
	res, err = httpclient.HttpDefaultGetWithTimeout(f.Config.GetCollectDevicesUrl+"&orgId="+orgId, f.Config.WebOpTimeoutMs)
	if err != nil {
		return valueRet, fmt.Errorf("httpDefaultGetWithTimeout, %s", err)
	}

	//解析CollectDevice
	var collectDeviceResponse CollectDeviceResponse
	err = json.Unmarshal(res, &collectDeviceResponse)
	if err != nil {
		return valueRet, fmt.Errorf("json.Unmarshal, %s", err)
	}

	if collectDeviceResponse.Code != 0 {
		return valueRet, fmt.Errorf("code in response error, response=%s", string(res))
	}

	for _, collectDevice := range collectDeviceResponse.Data {
		valueRet[collectDevice.Id] = collectDevice
	}

	return valueRet, nil
}

//请求Template
func (f *DataAccessWork) RequestTemplateData(orgId string) (map[int64]Template, error) {
	var valueRet map[int64]Template
	var res []byte
	var err error

	valueRet = make(map[int64]Template)

	//构造Request请求，获取Template信息
	res, err = httpclient.HttpDefaultGetWithTimeout(f.Config.GetTemplatesUrl+"&orgId="+orgId, f.Config.WebOpTimeoutMs)
	if err != nil {
		return valueRet, fmt.Errorf("httpDefaultGetWithTimeout, %s", err)
	}

	//解析Template
	var templateResponse TemplateResponse
	err = json.Unmarshal(res, &templateResponse)
	if err != nil {
		return valueRet, fmt.Errorf("json.Unmarshal, %s", err)
	}

	if templateResponse.Code != 0 {
		return valueRet, fmt.Errorf("code in response error, response=%s", string(res))
	}

	for _, template := range templateResponse.Data {
		valueRet[template.Id] = template
	}

	return valueRet, nil
}

//请求Mapping
func (f *DataAccessWork) RequestMappingData(orgId string) (map[int64]Mapping, error) {
	var valueRet map[int64]Mapping
	var res []byte
	var err error

	valueRet = make(map[int64]Mapping)

	//构造Request请求，获取Mapping信息
	res, err = httpclient.HttpDefaultGetWithTimeout(f.Config.GetMappingsUrl+"&orgId="+orgId, f.Config.WebOpTimeoutMs)
	if err != nil {
		return valueRet, fmt.Errorf("httpDefaultGetWithTimeout, %s", err)
	}

	//解析Mapping
	var mappingResponse MappingResponse
	err = json.Unmarshal(res, &mappingResponse)
	if err != nil {
		return valueRet, fmt.Errorf("json.Unmarshal, %s", err)
	}

	if mappingResponse.Code != 0 {
		return valueRet, fmt.Errorf("code in response error, response=%s", string(res))
	}

	for _, mapping := range mappingResponse.Data {
		valueRet[mapping.Id] = mapping
	}

	return valueRet, nil
}

//生成base
func (f *DataAccessWork) GetBaseForDataAccess() (*DataAccessBase, error) {
	var base DataAccessBase
	var err error
	base.OrgId, err = f.RequestOrgId()
	if err != nil {
		return nil, fmt.Errorf("requestOrgId, %s", err)
	}

	base.CollectDeviceIdToCollectDeviceMap, err = f.RequestCollectDeviceData(base.OrgId)
	if err != nil {
		return nil, fmt.Errorf("requestCollectDeviceData, %s", err)
	}

	base.TemplateIdToTemplateMap, err = f.RequestTemplateData(base.OrgId)
	if err != nil {
		return nil, fmt.Errorf("requestTemplateData, %s", err)
	}

	base.MappingIdToMappingMap, err = f.RequestMappingData(base.OrgId)
	if err != nil {
		return nil, fmt.Errorf("requestTemplateData, %s", err)
	}

	return &base, nil
}

//从base生成advance
func (f *DataAccessWork) GetAdvanceFromBaseForDataAccess(base *DataAccessBase) (*DataAccessAdvance, error) {
	if base == nil {
		return nil, errors.New("base is nil error")
	}

	var advance DataAccessAdvance
	var ok bool

	advance.CollectDeviceIdToCollectId = make(map[string]string)
	for _, v := range base.CollectDeviceIdToCollectDeviceMap {
		advance.CollectDeviceIdToCollectId[v.Id] = v.CollectId
	}

	advance.CollectDeviceIdToTemplateIdMap = make(map[string]int64)
	for _, v := range base.CollectDeviceIdToCollectDeviceMap {
		advance.CollectDeviceIdToTemplateIdMap[v.Id] = v.TemplateId
	}

	advance.AssetIdToCollectDeviceIdMap = make(map[string]map[string]string)
	for _, v := range base.CollectDeviceIdToCollectDeviceMap {
		var collectDeviceIdMapTemp map[string]string
		collectDeviceIdMapTemp, ok = advance.AssetIdToCollectDeviceIdMap[v.AssetId]
		if ok {
			collectDeviceIdMapTemp[v.Id] = "1"
			advance.AssetIdToCollectDeviceIdMap[v.AssetId] = collectDeviceIdMapTemp
		} else {
			collectDeviceIdMapTemp = make(map[string]string)
			collectDeviceIdMapTemp[v.Id] = "1"
			advance.AssetIdToCollectDeviceIdMap[v.AssetId] = collectDeviceIdMapTemp
		}
	}

	advance.TemplateIdToMappingIdMap = make(map[int64]map[int64]string)
	for _, v := range base.MappingIdToMappingMap {
		var MappingIdMapTemp map[int64]string
		MappingIdMapTemp, ok = advance.TemplateIdToMappingIdMap[v.TemplateId]
		if !ok {
			MappingIdMapTemp = make(map[int64]string)
		}
		MappingIdMapTemp[v.Id] = "1"
		advance.TemplateIdToMappingIdMap[v.TemplateId] = MappingIdMapTemp
	}

	advance.TemplateIdFeatureNameDeviceLogicalIdToMappingIdNonControlMap = make(map[int64]map[string]map[string]int64)
	for k, v := range advance.TemplateIdToMappingIdMap {
		var featureNameDeviceLogicalIdToMappingIdMap map[string]map[string]int64
		featureNameDeviceLogicalIdToMappingIdMap, ok = advance.TemplateIdFeatureNameDeviceLogicalIdToMappingIdNonControlMap[k]
		if !ok {
			featureNameDeviceLogicalIdToMappingIdMap = make(map[string]map[string]int64)
		}

		for mappingId := range v {
			var mappingTemp Mapping
			mappingTemp, ok = base.MappingIdToMappingMap[mappingId]
			if !ok {
				continue
			}

			if mappingTemp.Control {
				//控制mapping不考虑
				continue
			}

			var deviceLogicalId string
			//动态模板下，这个是assetId，其他情况，这个是空串
			deviceLogicalId, _ = mappingTemp.Attributes[DEVICE_LOGICAL_ID]

			var deviceLogicalIdToMappingIdMap map[string]int64
			deviceLogicalIdToMappingIdMap, ok = featureNameDeviceLogicalIdToMappingIdMap[mappingTemp.FeatureName]
			if !ok {
				deviceLogicalIdToMappingIdMap = make(map[string]int64)
			}

			deviceLogicalIdToMappingIdMap[deviceLogicalId] = mappingId
			featureNameDeviceLogicalIdToMappingIdMap[mappingTemp.FeatureName] = deviceLogicalIdToMappingIdMap
		}

		advance.TemplateIdFeatureNameDeviceLogicalIdToMappingIdNonControlMap[k] = featureNameDeviceLogicalIdToMappingIdMap
	}

	advance.TemplateIdToDevicePointNameMappingMap = make(map[int64]map[string][]int64)
	for templateId, mappingIdMap := range advance.TemplateIdToMappingIdMap {
		if len(mappingIdMap) <= 0 {
			continue
		}

		var devicePointNameMappingMap = make(map[string][]int64)
		var mappingIdArray []int64
		var mappingId int64
		for mappingId = range mappingIdMap {
			var mappingTemp Mapping
			mappingTemp, ok = base.MappingIdToMappingMap[mappingId]
			if !ok {
				continue
			}
			for _, devicePointName := range mappingTemp.MappingRule.Arguments {
				mappingIdArray, ok = devicePointNameMappingMap[devicePointName]
				if !ok {
					mappingIdArray = nil
					mappingIdArray = append(mappingIdArray, mappingId)
					devicePointNameMappingMap[devicePointName] = mappingIdArray
				} else {
					mappingIdArray = append(mappingIdArray, mappingId)
					devicePointNameMappingMap[devicePointName] = mappingIdArray
				}
			}
		}
		if len(devicePointNameMappingMap) <= 0 {
			continue
		}

		advance.TemplateIdToDevicePointNameMappingMap[templateId] = devicePointNameMappingMap
	}

	advance.TemplateIdFeatureNameDeviceLogicalIdToMappingIdControlMap = make(map[int64]map[string]map[string]int64)
	for k, v := range advance.TemplateIdToMappingIdMap {
		var featureNameDeviceLogicalIdToMappingIdMap map[string]map[string]int64
		featureNameDeviceLogicalIdToMappingIdMap, ok = advance.TemplateIdFeatureNameDeviceLogicalIdToMappingIdControlMap[k]
		if !ok {
			featureNameDeviceLogicalIdToMappingIdMap = make(map[string]map[string]int64)
		}

		for mappingId := range v {
			var mappingTemp Mapping
			mappingTemp, ok = base.MappingIdToMappingMap[mappingId]
			if !ok {
				continue
			}

			if !(mappingTemp.Control) {
				continue
			}

			var deviceLogicalId string
			//动态模板下，这个是assetId，其他情况，这个是空串
			deviceLogicalId, _ = mappingTemp.Attributes[DEVICE_LOGICAL_ID]

			var deviceLogicalIdToMappingIdMap map[string]int64
			deviceLogicalIdToMappingIdMap, ok = featureNameDeviceLogicalIdToMappingIdMap[mappingTemp.FeatureName]
			if !ok {
				deviceLogicalIdToMappingIdMap = make(map[string]int64)
			}

			deviceLogicalIdToMappingIdMap[deviceLogicalId] = mappingId
			featureNameDeviceLogicalIdToMappingIdMap[mappingTemp.FeatureName] = deviceLogicalIdToMappingIdMap
		}

		advance.TemplateIdFeatureNameDeviceLogicalIdToMappingIdControlMap[k] = featureNameDeviceLogicalIdToMappingIdMap
	}

	return &advance, nil
}

//生成dataAccess
func (f *DataAccessWork) GetDataAccess() (*DataAccess, error) {
	var data DataAccess
	var err error
	data.Base, err = f.GetBaseForDataAccess()
	if err != nil {
		return nil, fmt.Errorf("getBase %s", err)
	}

	data.Advance, err = f.GetAdvanceFromBaseForDataAccess(data.Base)
	if err != nil {
		return nil, fmt.Errorf("getAdvanceFromBase %s", err)
	}

	return &data, nil
}
