package metadata

import (
	"edge-common-go/pkg/httpclient"
	"encoding/json"
	"errors"
	"fmt"
	"strings"
)

/**
 * Description:获取一些模型资产相关的元数据，如模型、资产、资产树等
 *
 * @author yang.zhang4
 * @date 2020/12/3
 */
const (
	MODEL_PATH_SEP              = "/"
	MODEL_ASSET_POINT_SEPARATOR = "::"

	VALUE_TYPE_ARRAY     = "ARRAY"
	VALUE_TYPE_BOOL      = "BOOL"
	VALUE_TYPE_DATE      = "DATE"
	VALUE_TYPE_ENUM      = "ENUM"
	VALUE_TYPE_INT       = "INT"
	VALUE_TYPE_FLOAT     = "FLOAT"
	VALUE_TYPE_DOUBLE    = "DOUBLE"
	VALUE_TYPE_STRUCT    = "STRUCT"
	VALUE_TYPE_TEXT      = "TEXT"
	VALUE_TYPE_STRING    = "STRING"
	VALUE_TYPE_NUMBER    = "NUMBER"
	VALUE_TYPE_TIMESTAMP = "TIMESTAMP"
	VALUE_TYPE_FILE      = "FILE"
	VALUE_TYPE_MAP       = "MAP"
	VALUE_TYPE_UNKNOWN   = "UNKNOWN"

	//模型点类型定义-->
	MODEL_POINT_MEASUREPOINT = "MEASUREPOINT"
	MODEL_POINT_EVENT        = "EVENT"

	MODEL_POINT_MEASUREPOINT_INT32 = 0
	MODEL_POINT_EVENT_INT32        = 1
	MODEL_POINT_ATTR_INT32         = 2
	MODEL_POINT_SERVICE_INT32      = 3
	//<--模型点类型定义
)

//通用结构-->

//多语言结构
type TSLStringI18n struct {
	DefaultValue string            `json:"defaultValue"`
	I18nValue    map[string]string `json:"i18nValue"`
}

//单位
type Unit struct {
	UnitId     string `json:"unitId"`
	Multiplier string `json:"multiplier"`
}

//dataType为ARRAY时，DataDefinition结构，ItemType只能为INT”、“FLOAT”、“DOUBLE”、或“STRING”
type DataDefinitionArray struct {
	ItemType *string `json:"itemType"`
}

//dataType为ENUM时，DataDefinition结构，EnumType只允许INT和STRING
type DataDefinitionEnum struct {
	EnumDesc     map[string]TSLStringI18n `bson:"enumDesc" json:"enumDesc"`
	EnumType     *string                  `bson:"enumType" json:"enumType"`
	DefaultValue interface{}              `bson:"defaultValue" json:"defaultValue"`
}

//里面存放两个常用的模型属性
type SimpleModelData struct {
	ModelId     string `json:"modelId"`
	ModelIdPath string `json:"modelIdPath"`
}

//请求时的body使用的一些通用参数，分页、裁剪、动作等
type EnosPageDataRequestBody struct {
	Pagination Pagination `json:"pagination"`
	Projection []string   `json:"projection"`
	Action     string     `json:"action"`
}

//请求时的body中使用的分页结构
type Pagination struct {
	PageNo   int32    `json:"pageNo"`
	PageSize int32    `json:"pageSize"`
	Sorters  []Sorter `json:"sorters"`
}

//请求返回结构中的Pagination使用的排序相关结构
type Sorter struct {
	Field string `json:"field"`
	Order string `json:"order"`
}

//请求返回结构中的Pagination使用的排序相关结构
type EnosSortedDataV2 struct {
	SortedBy []Sorter `json:"sortedBy"`
}

//请求返回结构中的Pagination结构
type EnosPageDataV2 struct {
	EnosSortedDataV2
	PageNo    int32 `json:"pageNo"`
	PageSize  int32 `json:"pageSize"`
	TotalSize int32 `json:"totalSize"`
}

//<--通用结构

//资产相关-->

//资产请求返回结构
type TSLInstancesResponse struct {
	Code   int32             `json:"code"`
	Status int               `json:"status"`
	Msg    string            `json:"msg"`
	SubMsg string            `json:"submsg"`
	Data   []TSLInstanceData `json:"data"`
}

//资产结构
type TSLInstanceData struct {
	Id           string                 `json:"id"`
	OrgId        string                 `json:"orgId"`
	Name         TSLStringI18n          `json:"name"`
	DeviceKey    string                 `json:"deviceKey"`
	DeviceSecret string                 `json:"deviceSecret"`
	ProductKey   string                 `json:"productKey"`
	Desc         string                 `json:"desc"`
	Timezone     string                 `json:"timezone"`
	Attributes   map[string]interface{} `json:"attributes"`
	ModelId      string                 `json:"modelId"`
	Tags         map[string]string      `json:"tags"`
}

//<--资产相关

//模型相关-->

//模型请求返回结构
type TSLModelsResponse struct {
	Code      int            `json:"code"`
	Status    int            `json:"status"`
	Msg       string         `json:"msg"`
	RequestId string         `json:"requestId"`
	Data      []TSLModelData `json:"data"`
}

//measurepoint类型测点结构
type TSLMeasurePoint struct {
	Identifier     string            `json:"identifier"`
	Name           TSLStringI18n     `json:"name"`
	Desc           string            `json:"desc"`
	Tags           map[string]string `json:"tags"`
	DataType       string            `json:"dataType"`
	InnerType      string            `json:"innerType"`
	DataDefinition string            `json:"dataDefinition"`
	Unit           map[string]string `json:"unit"`
	HasQuality     bool              `json:"hasQuality"`
	SignalType     string            `json:"signalType"`
}

//service类型测点使用的InputData和OutputData结构
type TSLServiceInputOutputData struct {
	Identifier     string            `json:"identifier"`
	Name           TSLStringI18n     `json:"name"`
	Desc           string            `json:"desc"`
	Tags           map[string]string `json:"tags"`
	DataType       string            `json:"dataType"`
	Unit           map[string]string `json:"unit"`
	DataDefinition string            `json:"dataDefinition"`
}

//service类型测点结构
type TSLService struct {
	Identifier string                      `json:"identifier"`
	Name       TSLStringI18n               `json:"name"`
	Desc       string                      `json:"desc"`
	Tags       map[string]string           `json:"tags"`
	InputData  []TSLServiceInputOutputData `json:"inputData"`
	OutputData []TSLServiceInputOutputData `json:"outputData"`
	CallType   string                      `json:"CallType"`
}

//event类型测点使用的OutputData结构
type ThingDataPoint struct {
	Identifier     string            `json:"identifier"`
	Name           TSLStringI18n     `json:"name"`
	Desc           string            `json:"desc"`
	Tags           map[string]string `json:"tags"`
	DataType       string            `json:"dataType"`
	InnerType      string            `json:"innerType"`
	DataDefinition string            `json:"dataDefinition"`
	Unit           Unit              `json:"unit"`
}

//event类型测点结构
type TSLEvent struct {
	Identifier    string                    `json:"identifier"`
	Name          TSLStringI18n             `json:"name"`
	Desc          string                    `json:"desc"`
	Tags          map[string]string         `json:"tags"`
	OutputData    []ThingDataPoint          `json:"outputData"`
	OutputDataMap map[string]ThingDataPoint `json:"outputDataMap"`
	EventType     string                    `json:"eventType"`
}

//模型结构
type TSLModelData struct {
	ModelId       string                     `json:"modelId"`
	ModelIdPath   string                     `json:"modelIdPath"`
	OrgId         string                     `json:"orgId"`
	Name          TSLStringI18n              `json:"name"`
	Desc          string                     `json:"desc"`
	Tags          map[string]string          `json:"tags"`
	Attributes    map[string]interface{}     `json:"attributes"`
	Measurepoints map[string]TSLMeasurePoint `json:"measurepoints"`
	Services      map[string]TSLService      `json:"services"`
	Events        map[string]TSLEvent        `json:"events"`
}

//<--模型相关

//资产树节点相关-->
//资产树节点使用的ExtraInfo结构
type BSLExtraInfo struct {
	CreateBy   string `json:"createBy"`
	CreateTime int64  `json:"createTime"`
	UpdateBy   string `json:"updateBy"`
	UpdateTime int64  `json:"updateTime"`
	Ns         string `json:"ns"`
}

//资产树节点结构
type NodeData struct {
	Id               string        `json:"id"`
	ModelId          string        `json:"modelId"`
	InstanceId       string        `json:"instanceId"`
	Name             TSLStringI18n `json:"name"`
	IsDeviceInstance bool          `json:"isDeviceInstance"`
	TreeId           string        `json:"treeId"`
	ParentId         string        `json:"parentId"`
	HasChild         bool          `json:"hasChild"`
	ExtraInfo        BSLExtraInfo  `json:"extraInfo"`
}

//资产树节点请求返回结构
type NodesResponse struct {
	Code   int        `json:"code"`
	Status int        `json:"status"`
	Msg    string     `json:"msg"`
	SubMsg string     `json:"subMsg"`
	Data   []NodeData `json:"data"`
}

//<--资产树节点相关

//资产树相关-->
//资产树中的Asset结构
type AssetVo struct {
	InValid     bool                   `json:"inValid"`
	AssetId     string                 `json:"assetId"`
	ModelId     string                 `json:"modelId"`
	ModelIdPath string                 `json:"modelIdPath"`
	Name        TSLStringI18n          `json:"name"`
	Timezone    string                 `json:"timezone"`
	Description string                 `json:"description"`
	Label       string                 `json:"label"`
	Attributes  map[string]interface{} `json:"attributes"`
	Tags        map[string]string      `json:"tags"`
}

//资产树结构
type TreeVo struct {
	TreeId string            `json:"treeId"`
	Tags   map[string]string `json:"tags"`
	Asset  AssetVo           `json:"asset"`
}

//资产树请求时的body使用的Filter结构
type TreeSearchVo struct {
	Tags map[string]string `json:"tags"`
}

//资产树请求时的body结构
type SearchAssetTreeRequestBody struct {
	Filter TreeSearchVo `json:"filter"`
	EnosPageDataRequestBody
}

//资产树请求返回结构
type TreeVoListEnosRspV2 struct {
	Code       int32          `json:"code"`
	Status     int            `json:"status"`
	Msg        string         `json:"msg"`
	RequestId  string         `json:"requestId"`
	Data       []TreeVo       `json:"data"`
	Pagination EnosPageDataV2 `json:"pagination"`
}

//<--资产树相关

type ModelAssetBase struct {
	OrgId string
	//key = nodeData.TreeId + "::" + nodeData.InstanceId
	NodeCache map[string]NodeData
	//key = TSLInstanceData.Id(assetId)
	TSLInstanceDataCache map[string]TSLInstanceData
	//key = modelId
	TSLModelDataCache map[string]TSLModelData
	//map[treeId]commonutils.TreeVo
	TreeIdToTreeVoMap map[string]TreeVo
}

type ModelAssetAdvance struct {
	//key = nodeId
	NodeIdToNodeDataMap map[string]NodeData
	//map[assetId]SimpleModelData
	AssetIdToSimpleModelData map[string]SimpleModelData
	//一个model和他所有的子model的map，map[modelId][childModelId]"1"
	ModelIdToChildMap map[string]map[string]string
	//nodeId和这个nodeId对应的assetId的完整路径的map，map[nodeId]assetPath
	NodeIdToAssetPathMap map[string]string
	//assetId, treeId to nodeId, map[assetId][treeId]nodeId
	AssetIdTreeIdToNodeIdMap map[string]map[string]string
	//资产树中，一个NodeId和他的直接子节点NodeId的map
	NodeIdToDirectChildMap map[string]map[string]string
	//一个modelId实例化了多少个assetId，map[modelId][assetId]"1"
	ModelIdToAssetIdMap map[string]map[string]string
	//一个modelId和他的所有子模型的assetId的map（包括modelId自己的assetId）
	ModelIdToChildAssetIdMap map[string]map[string]string
	//map[tagKey::tagValue]map[assetId]string
	AssetTagToAssetIdMap map[string]map[string]string
	//map[treeId]map[assetId]string
	TreeIdToAssetIdMap map[string]map[string]string
	//map[tagKey::tagValue]map[treeId]string
	TreeTagToTreeIdMap map[string]map[string]string
	//一个productKey实例化了多少个assetId，map[productKey][assetId]"1"
	ProductKeyToAssetIdMap map[string]map[string]string
	//productKey+deviceKey至assetId的map
	ProductKeyDeviceKeyToAssetIdMap map[string]string
}

type ModelAsset struct {
	Base    *ModelAssetBase
	Advance *ModelAssetAdvance
}

type ModelAssetConfig struct {
	WebOpTimeoutMs int64

	GetOrgIdUrl        string
	GetTSLInstancesUrl string
	GetThingModelsUrl  string
	GetNodesUrl        string
	GetTreesUrl        string

	//LogArray []*log.LogAa
}

type ModelAssetWork struct {
	Config ModelAssetConfig
}

func CheckModelAssetIsValid(modelAsset *ModelAsset) bool {
	if modelAsset == nil || modelAsset.Base == nil || modelAsset.Advance == nil {
		return false
	}

	return true
}

//根据assetId获取tslInstanceData
func GetTSLInstanceDataByAssetId(modelAsset *ModelAsset, assetId string) (TSLInstanceData, error) {
	if !CheckModelAssetIsValid(modelAsset) {
		return TSLInstanceData{}, ERROR_RESOURCE_IS_NIL
	}

	var data TSLInstanceData
	var ok bool
	data, ok = modelAsset.Base.TSLInstanceDataCache[assetId]
	if !ok {
		return TSLInstanceData{}, ERROR_NOT_FOUND
	}
	return data, nil
}

//根据modelId获取tslModelData
func GetTSLModelDataByModelId(modelAsset *ModelAsset, modelId string) (TSLModelData, error) {
	if !CheckModelAssetIsValid(modelAsset) {
		return TSLModelData{}, ERROR_RESOURCE_IS_NIL
	}

	var data TSLModelData
	var ok bool
	data, ok = modelAsset.Base.TSLModelDataCache[modelId]
	if !ok {
		return TSLModelData{}, ERROR_NOT_FOUND
	}
	return data, nil
}

//根据assetId获取simpleModelData
func GetSimpleModelDataByAssetId(modelAsset *ModelAsset, assetId string) (SimpleModelData, error) {
	if !CheckModelAssetIsValid(modelAsset) {
		return SimpleModelData{}, ERROR_RESOURCE_IS_NIL
	}

	var data SimpleModelData
	var ok bool
	data, ok = modelAsset.Advance.AssetIdToSimpleModelData[assetId]
	if !ok {
		return SimpleModelData{}, ERROR_NOT_FOUND
	}
	return data, nil
}

//根据nodeId获取这个nodeId对应的assetId的完整路径
func GetAssetPathByNodeId(modelAsset *ModelAsset, nodeId string) (string, error) {
	if !CheckModelAssetIsValid(modelAsset) {
		return "", ERROR_RESOURCE_IS_NIL
	}

	var assetPath string
	var ok bool
	assetPath, ok = modelAsset.Advance.NodeIdToAssetPathMap[nodeId]
	if !ok {
		return "", ERROR_NOT_FOUND
	}
	return assetPath, nil
}

//根据assetId, 获取treeIdToNodeIdMap
func GetTreeIdToNodeIdMapByAssetId(modelAsset *ModelAsset, assetId string) (map[string]string, error) {
	if !CheckModelAssetIsValid(modelAsset) {
		return nil, ERROR_RESOURCE_IS_NIL
	}

	var data map[string]string
	var ok bool
	data, ok = modelAsset.Advance.AssetIdTreeIdToNodeIdMap[assetId]
	if !ok {
		return nil, ERROR_NOT_FOUND
	}
	return data, nil
}

//根据assetId, treeId，获取nodeId
func GetNodeIdByAssetIdTreeId(modelAsset *ModelAsset, assetId string, treeId string) (string, error) {
	if !CheckModelAssetIsValid(modelAsset) {
		return "", ERROR_RESOURCE_IS_NIL
	}

	var nodeId string
	var ok bool
	nodeId, ok = modelAsset.Advance.AssetIdTreeIdToNodeIdMap[assetId][treeId]
	if !ok {
		return "", ERROR_NOT_FOUND
	}
	return nodeId, nil
}

//获取一个modelId和他的所有子模型的assetId的map（包括modelId自己的assetId）
func GetChildAssetIdMapByModelId(modelAsset *ModelAsset, modelId string) (map[string]string, error) {
	if !CheckModelAssetIsValid(modelAsset) {
		return nil, ERROR_RESOURCE_IS_NIL
	}

	var childAssetIdMap map[string]string
	var ok bool
	childAssetIdMap, ok = modelAsset.Advance.ModelIdToChildAssetIdMap[modelId]
	if !ok {
		return nil, ERROR_NOT_FOUND
	}
	return childAssetIdMap, nil
}

//获取productKey+deviceKey的key
func GetKeyOfProductKeyDeviceKey(productKey string, deviceKey string) string {
	return productKey + "::" + deviceKey
}

//获取NodeCache的key
func GetNodeCacheKey(treeId string, assetId string) string {
	return treeId + "::" + assetId
}

//获取一个modelId及其所有子模型的modelId
func GetModelIdAllChildModelId(modelAsset *ModelAsset, modelId string) (map[string]string, error) {
	var valueRet map[string]string
	valueRet = make(map[string]string)
	if modelAsset == nil || modelAsset.Base == nil || modelAsset.Advance == nil {
		return valueRet, errors.New("resource is nil error")
	}

	valueRet[modelId] = "1"

	var childMap map[string]string
	var ok bool
	childMap, ok = modelAsset.Advance.ModelIdToChildMap[modelId]
	if !ok {
		return valueRet, nil
	}

	for k := range childMap {
		valueRet[k] = "1"
	}

	return valueRet, nil
}

//获取一个nodeId及其所有子节点的nodeId
func GetNodeIdAllChildNodeId(modelAsset *ModelAsset, nodeId string) (map[string]string, error) {
	var valueRet map[string]string
	valueRet = make(map[string]string)
	if modelAsset == nil || modelAsset.Base == nil || modelAsset.Advance == nil {
		return valueRet, errors.New("resource is nil error")
	}

	valueRet[nodeId] = "1"

	var ok bool
	var directChildMap map[string]string
	directChildMap, ok = modelAsset.Advance.NodeIdToDirectChildMap[nodeId]
	if !ok || len(directChildMap) <= 0 {
		return valueRet, nil
	}

	for k := range directChildMap {
		var directChildMapTemp map[string]string
		directChildMapTemp, _ = GetNodeIdAllChildNodeId(modelAsset, k)

		for k1 := range directChildMapTemp {
			valueRet[k1] = "1"
		}
	}

	return valueRet, nil
}

//获取一个nodeId及其所有子节点的assetId
func GetNodeIdAllChildAssetId(modelAsset *ModelAsset, nodeId string) (map[string]string, error) {
	var valueRet map[string]string
	valueRet = make(map[string]string)
	if modelAsset == nil || modelAsset.Base == nil || modelAsset.Advance == nil {
		return valueRet, errors.New("resource is nil error")
	}

	var allNodeIdMap map[string]string
	var err error
	allNodeIdMap, err = GetNodeIdAllChildNodeId(modelAsset, nodeId)
	if err != nil {
		return valueRet, fmt.Errorf("getNodeIdAllChildNodeId, %s", err)
	}

	var nodeData NodeData
	var ok bool
	for nodeId := range allNodeIdMap {
		nodeData, ok = modelAsset.Advance.NodeIdToNodeDataMap[nodeId]
		if !ok {
			continue
		}
		valueRet[nodeData.InstanceId] = "1"
	}

	return valueRet, nil
}

//请求资产树节点信息
func (f *ModelAssetWork) RequestNodeData(orgId string) (map[string]NodeData, error) {
	var valueRet map[string]NodeData
	var res []byte
	var err error

	valueRet = make(map[string]NodeData)

	//构造Request请求，获取盒子上的所有Node信息
	res, err = httpclient.HttpDefaultGetWithTimeout(f.Config.GetNodesUrl+"&orgId="+orgId, f.Config.WebOpTimeoutMs)
	if err != nil {
		return valueRet, fmt.Errorf("httpDefaultGetWithTimeout, %s", err)
	}

	//解析Node信息
	var nodeResponse NodesResponse
	err = json.Unmarshal(res, &nodeResponse)
	if err != nil {
		return valueRet, fmt.Errorf("json.Unmarshal, %s", err)
	}

	if nodeResponse.Status != 0 || nodeResponse.Code != 0 {
		return valueRet, fmt.Errorf("code or status in response error, response=%s", string(res))
	}

	for _, nodeData := range nodeResponse.Data {
		key := GetNodeCacheKey(nodeData.TreeId, nodeData.InstanceId)
		valueRet[key] = nodeData
	}

	return valueRet, nil
}

//请求资产信息
func (f *ModelAssetWork) RequestTSLInstanceData(orgId string) (map[string]TSLInstanceData, error) {
	var valueRet map[string]TSLInstanceData
	var res []byte
	var err error

	valueRet = make(map[string]TSLInstanceData)

	//构造Request请求，获取TSLInstance信息
	res, err = httpclient.HttpDefaultGetWithTimeout(f.Config.GetTSLInstancesUrl+"&orgId="+orgId, f.Config.WebOpTimeoutMs)
	if err != nil {
		return valueRet, fmt.Errorf("httpDefaultGetWithTimeout, %s", err)
	}

	//解析TSLInstance
	var tSLInstance TSLInstancesResponse
	err = json.Unmarshal(res, &tSLInstance)
	if err != nil {
		return valueRet, fmt.Errorf("json.Unmarshal, %s", err)
	}

	if tSLInstance.Status != 0 || tSLInstance.Code != 0 {
		return valueRet, fmt.Errorf("code or status in response error, response=%s", string(res))
	}

	for _, tSLInstanceData := range tSLInstance.Data {
		valueRet[tSLInstanceData.Id] = tSLInstanceData
	}

	return valueRet, nil
}

//请求所有模型信息
func (f *ModelAssetWork) RequestTSLModelsData(orgId string) (map[string]TSLModelData, error) {
	var valueRet map[string]TSLModelData
	var res []byte
	var err error

	valueRet = make(map[string]TSLModelData)

	//构造Request请求，获取TSLModels信息
	res, err = httpclient.HttpDefaultGetWithTimeout(f.Config.GetThingModelsUrl+"&orgId="+orgId, f.Config.WebOpTimeoutMs)
	if err != nil {
		return valueRet, fmt.Errorf("httpDefaultGetWithTimeout, %s", err)
	}

	//解析TSLInstance
	var tslModelsResponse TSLModelsResponse
	err = json.Unmarshal(res, &tslModelsResponse)
	if err != nil {
		return valueRet, fmt.Errorf("json.Unmarshal, %s", err)
	}

	if tslModelsResponse.Status != 0 || tslModelsResponse.Code != 0 {
		return valueRet, fmt.Errorf("code or status in response error, response=%s", string(res))
	}

	for k := range tslModelsResponse.Data {
		err = f.operateTSLModel(&(tslModelsResponse.Data[k]))
		if err != nil {
			valueRet = make(map[string]TSLModelData)
			return valueRet, fmt.Errorf("operateTSLModel, %s", err)
		}
		valueRet[tslModelsResponse.Data[k].ModelId] = tslModelsResponse.Data[k]
	}

	return valueRet, nil
}

//请求资产树信息
func (f *ModelAssetWork) RequestTreeData(orgId string) (map[string]TreeVo, error) {
	var valueRet map[string]TreeVo
	var res []byte
	var err error

	valueRet = make(map[string]TreeVo)

	var requestBody SearchAssetTreeRequestBody
	requestBody.Action = "search"
	requestBody.Pagination.PageNo = 1
	requestBody.Pagination.PageSize = 20000

	//构造Request请求，获取资产树信息
	res, err = httpclient.HttpDefaultPostWithTimeout(f.Config.GetTreesUrl+"&orgId="+orgId, requestBody, f.Config.WebOpTimeoutMs)
	if err != nil {
		return valueRet, fmt.Errorf("httpDefaultGetWithTimeout, %s", err)
	}

	//解析资产树
	var treeVoListEnosRspV2 TreeVoListEnosRspV2
	err = json.Unmarshal(res, &treeVoListEnosRspV2)
	if err != nil {
		return valueRet, fmt.Errorf("json.Unmarshal, %s", err)
	}

	if treeVoListEnosRspV2.Status != 0 || treeVoListEnosRspV2.Code != 0 {
		return valueRet, fmt.Errorf("code or status in response error, response=%s", string(res))
	}

	for _, treeVo := range treeVoListEnosRspV2.Data {
		valueRet[treeVo.TreeId] = treeVo
	}

	return valueRet, nil
}

//生成base
func (f *ModelAssetWork) GetBaseForModelAsset() (*ModelAssetBase, error) {
	var dataAccess = DataAccessWork{
		Config: DataAccessConfig{
			WebOpTimeoutMs: f.Config.WebOpTimeoutMs,
			GetOrgIdUrl:    f.Config.GetOrgIdUrl,
		},
	}

	var base ModelAssetBase
	var err error
	base.OrgId, err = dataAccess.RequestOrgId()
	if err != nil {
		return nil, fmt.Errorf("dataAccess.RequestOrgId, %s", err)
	}

	base.NodeCache, err = f.RequestNodeData(base.OrgId)
	if err != nil {
		return nil, fmt.Errorf("requestNodeData, %s", err)
	}

	base.TSLInstanceDataCache, err = f.RequestTSLInstanceData(base.OrgId)
	if err != nil {
		return nil, fmt.Errorf("requestTSLInstanceData, %s", err)
	}

	base.TSLModelDataCache, err = f.RequestTSLModelsData(base.OrgId)
	if err != nil {
		return nil, fmt.Errorf("requestTSLModelsData, %s", err)
	}

	base.TreeIdToTreeVoMap, err = f.RequestTreeData(base.OrgId)
	if err != nil {
		return nil, fmt.Errorf("requestTreeData, %s", err)
	}

	return &base, nil
}

//从base生成advance
func (f *ModelAssetWork) GetAdvanceFromBaseForModelAsset(base *ModelAssetBase) (*ModelAssetAdvance, error) {
	if base == nil {
		return nil, errors.New("base is nil error")
	}

	var advance ModelAssetAdvance
	var ok bool

	//构建NodeIdToNodeDataMap
	advance.NodeIdToNodeDataMap = make(map[string]NodeData)
	for _, v := range base.NodeCache {
		advance.NodeIdToNodeDataMap[v.Id] = v
	}

	//构建AssetIdToSimpleModelData
	advance.AssetIdToSimpleModelData = make(map[string]SimpleModelData)
	for _, v := range base.TSLInstanceDataCache {
		var modelData TSLModelData
		modelData, ok = base.TSLModelDataCache[v.ModelId]
		if !ok {
			continue
		}
		var simpleModelData = SimpleModelData{
			ModelId:     modelData.ModelId,
			ModelIdPath: modelData.ModelIdPath,
		}
		advance.AssetIdToSimpleModelData[v.Id] = simpleModelData
	}

	//构建ModelIdToChildMap
	advance.ModelIdToChildMap = make(map[string]map[string]string)
	var childModelIdMap map[string]string
	for _, v := range base.TSLModelDataCache {
		if len(v.ModelIdPath) <= 0 {
			continue
		}

		ss := strings.Split(v.ModelIdPath, MODEL_PATH_SEP)
		var sss []string
		for _, v1 := range ss {
			if len(v1) <= 0 {
				continue
			}
			sss = append(sss, v1)
		}

		if len(sss) <= 1 {
			continue
		}

		for ii := 0; ii < len(sss)-1; ii++ {
			childModelIdMap, ok = advance.ModelIdToChildMap[sss[ii]]
			if !ok {
				childModelIdMap = make(map[string]string)
			}

			for jj := ii + 1; jj < len(sss); jj++ {
				childModelIdMap[sss[jj]] = "1"
			}

			advance.ModelIdToChildMap[sss[ii]] = childModelIdMap
		}
	}

	//构建NodeIdToAssetPathMap
	advance.NodeIdToAssetPathMap = make(map[string]string)
	for _, v := range base.NodeCache {
		if len(v.InstanceId) <= 0 {
			continue
		}
		var assetPath = getAssetPathByNodeId(v.Id, advance.NodeIdToNodeDataMap)

		advance.NodeIdToAssetPathMap[v.Id] = assetPath
	}

	//构建AssetIdTreeIdToNodeIdMap
	advance.AssetIdTreeIdToNodeIdMap = make(map[string]map[string]string)
	for _, v := range base.NodeCache {
		var treeIdNodeIdMap map[string]string
		treeIdNodeIdMap, ok = advance.AssetIdTreeIdToNodeIdMap[v.InstanceId]
		if !ok {
			treeIdNodeIdMap = make(map[string]string)
		}
		treeIdNodeIdMap[v.TreeId] = v.Id
		advance.AssetIdTreeIdToNodeIdMap[v.InstanceId] = treeIdNodeIdMap
	}

	//构建NodeIdToDirectChildMap
	advance.NodeIdToDirectChildMap = make(map[string]map[string]string)
	for _, v := range base.NodeCache {
		if len(v.ParentId) <= 0 {
			continue
		}
		_, ok = advance.NodeIdToNodeDataMap[v.ParentId]
		if !ok {
			continue
		}
		var directChildMap map[string]string
		directChildMap, ok = advance.NodeIdToDirectChildMap[v.ParentId]
		if !ok {
			directChildMap = make(map[string]string)
		}
		directChildMap[v.Id] = "1"

		advance.NodeIdToDirectChildMap[v.ParentId] = directChildMap
	}

	//构建ModelIdToAssetIdMap
	advance.ModelIdToAssetIdMap = make(map[string]map[string]string)
	for _, v := range base.TSLInstanceDataCache {
		var assetIdMap map[string]string
		assetIdMap, ok = advance.ModelIdToAssetIdMap[v.ModelId]
		if !ok {
			assetIdMap = make(map[string]string)
		}
		assetIdMap[v.Id] = "1"

		advance.ModelIdToAssetIdMap[v.ModelId] = assetIdMap
	}

	//构建ModelIdToChildAssetIdMap
	advance.ModelIdToChildAssetIdMap = make(map[string]map[string]string)
	for _, v := range base.TSLModelDataCache {
		//先找出这个模型所有的父模型（包含自己）
		ss := strings.Split(v.ModelIdPath, MODEL_PATH_SEP)
		var sss []string
		for _, v1 := range ss {
			if len(v1) <= 0 {
				continue
			}
			sss = append(sss, v1)
		}

		if len(sss) <= 0 {
			continue
		}

		var assetIdMapFind map[string]string
		assetIdMapFind, ok = advance.ModelIdToAssetIdMap[v.ModelId]
		if !ok {
			continue
		}

		var assetIdMap map[string]string
		for _, v1 := range sss {
			assetIdMap, ok = advance.ModelIdToChildAssetIdMap[v1]
			if !ok {
				assetIdMap = make(map[string]string)
			}

			for v2 := range assetIdMapFind {
				assetIdMap[v2] = "1"
			}

			advance.ModelIdToChildAssetIdMap[v1] = assetIdMap
		}
	}

	//构建AssetTagToAssetIdMap
	advance.AssetTagToAssetIdMap = make(map[string]map[string]string)
	for _, v := range base.TSLInstanceDataCache {
		if len(v.Tags) <= 0 {
			var tagKeyValue = MODEL_ASSET_POINT_SEPARATOR
			var assetIdMap map[string]string
			assetIdMap, ok = advance.AssetTagToAssetIdMap[tagKeyValue]
			if !ok {
				assetIdMap = make(map[string]string)
			}
			assetIdMap[v.Id] = "1"

			advance.AssetTagToAssetIdMap[tagKeyValue] = assetIdMap
			continue
		}

		for tagKey, tagValue := range v.Tags {
			var tagKeyValue = tagKey + MODEL_ASSET_POINT_SEPARATOR + tagValue
			var assetIdMap map[string]string
			assetIdMap, ok = advance.AssetTagToAssetIdMap[tagKeyValue]
			if !ok {
				assetIdMap = make(map[string]string)
			}
			assetIdMap[v.Id] = "1"

			advance.AssetTagToAssetIdMap[tagKeyValue] = assetIdMap
		}
	}

	//构建TreeIdToAssetIdMap
	advance.TreeIdToAssetIdMap = make(map[string]map[string]string)
	var assetIdMap map[string]string
	for _, v := range base.NodeCache {
		assetIdMap, ok = advance.TreeIdToAssetIdMap[v.TreeId]
		if !ok {
			assetIdMap = make(map[string]string)
		}
		assetIdMap[v.InstanceId] = "1"
		advance.TreeIdToAssetIdMap[v.TreeId] = assetIdMap
	}

	//构建TreeTagToAssetIdMap
	advance.TreeTagToTreeIdMap = make(map[string]map[string]string)
	for _, v := range base.TreeIdToTreeVoMap {
		if len(v.Tags) <= 0 {
			var tagKeyValue = MODEL_ASSET_POINT_SEPARATOR
			var treeIdMap map[string]string
			treeIdMap, ok = advance.TreeTagToTreeIdMap[tagKeyValue]
			if !ok {
				treeIdMap = make(map[string]string)
			}
			treeIdMap[v.TreeId] = "1"

			advance.TreeTagToTreeIdMap[tagKeyValue] = treeIdMap
			continue
		}

		for tagKey, tagValue := range v.Tags {
			var tagKeyValue = tagKey + MODEL_ASSET_POINT_SEPARATOR + tagValue
			var treeIdMap map[string]string
			treeIdMap, ok = advance.TreeTagToTreeIdMap[tagKeyValue]
			if !ok {
				treeIdMap = make(map[string]string)
			}
			treeIdMap[v.TreeId] = "1"
			advance.TreeTagToTreeIdMap[tagKeyValue] = treeIdMap
		}
	}

	//构建ProductKeyToAssetIdMap
	advance.ProductKeyToAssetIdMap = make(map[string]map[string]string)
	for _, v := range base.TSLInstanceDataCache {
		if len(v.ProductKey) <= 0 {
			continue
		}
		assetIdMap, ok = advance.ProductKeyToAssetIdMap[v.ProductKey]
		if !ok {
			assetIdMap = make(map[string]string)
		}
		assetIdMap[v.Id] = "1"
		advance.ProductKeyToAssetIdMap[v.ProductKey] = assetIdMap
	}

	//构建ProductKeyDeviceKeyToAssetIdMap
	advance.ProductKeyDeviceKeyToAssetIdMap = make(map[string]string)
	var key string
	for _, v := range base.TSLInstanceDataCache {
		if len(v.ProductKey) <= 0 || len(v.DeviceKey) <= 0 {
			continue
		}

		key = GetKeyOfProductKeyDeviceKey(v.ProductKey, v.DeviceKey)

		advance.ProductKeyDeviceKeyToAssetIdMap[key] = v.Id
	}

	return &advance, nil
}

//生成modelasset
func (f *ModelAssetWork) GetModelAsset() (*ModelAsset, error) {
	var data ModelAsset
	var err error
	data.Base, err = f.GetBaseForModelAsset()
	if err != nil {
		return nil, fmt.Errorf("getBaseForModelAsset, %s", err)
	}

	data.Advance, err = f.GetAdvanceFromBaseForModelAsset(data.Base)
	if err != nil {
		return nil, fmt.Errorf("getAdvanceFromBaseForModelAsset, %s", err)
	}

	return &data, nil
}

//处理模型，提取一些有用的信息
func (f *ModelAssetWork) operateTSLModel(tslModel *TSLModelData) error {
	var errs []error
	var err error
	for k, v := range tslModel.Measurepoints {
		var innerType string
		innerType, err = f.getTSLInnerType(v.DataType, v.DataDefinition)
		if err != nil {
			errs = append(errs, fmt.Errorf("measurepoint, modelId=%s, pointId=%s, dataType=%s, dataDefinition=%s getTSLInnerType, %s", tslModel.ModelId, v.Identifier, v.DataType, v.DataDefinition, err))
			continue
		}
		v.InnerType = innerType
		tslModel.Measurepoints[k] = v
	}

	for k, v := range tslModel.Events {
		v.OutputDataMap = make(map[string]ThingDataPoint)
		for index, eventElement := range v.OutputData {
			var innerType string
			innerType, err = f.getTSLInnerType(eventElement.DataType, eventElement.DataDefinition)
			if err != nil {
				errs = append(errs, fmt.Errorf("event, modelId=%s, pointId=%s, eventKey=%s, dataType=%s, dataDefinition=%s getTSLInnerType, %s", tslModel.ModelId, v.Identifier, eventElement.Identifier, eventElement.DataType, eventElement.DataDefinition, err))
				continue
			}
			eventElement.InnerType = innerType
			v.OutputData[index] = eventElement
			v.OutputDataMap[eventElement.Identifier] = eventElement
		}
		tslModel.Events[k] = v
	}

	if len(errs) > 0 {
		return fmt.Errorf("%+v", errs)
	}

	return nil
}

//获取一些特殊数据类型的内部类型
func (f *ModelAssetWork) getTSLInnerType(dataType string, dataDefinition string) (string, error) {
	var err error
	switch dataType {
	case VALUE_TYPE_ARRAY:
		var dataDef DataDefinitionArray
		err = json.Unmarshal([]byte(dataDefinition), &dataDef)
		if err != nil {
			return "", err
		}
		if dataDef.ItemType == nil || len(*(dataDef.ItemType)) <= 0 {
			return "", fmt.Errorf("no itemType")
		}
		var innerType = strings.ToUpper(*(dataDef.ItemType))
		return innerType, nil
	case VALUE_TYPE_ENUM:
		var dataDef DataDefinitionEnum
		err = json.Unmarshal([]byte(dataDefinition), &dataDef)
		if err != nil {
			return "", err
		}
		if dataDef.EnumType == nil || len(*(dataDef.EnumType)) <= 0 {
			return "", fmt.Errorf("no enumType")
		}
		var innerType = strings.ToUpper(*(dataDef.EnumType))
		return innerType, nil
	default:
		return "", nil
	}
}

//获取一个nodeId的assetId的在资产树中的路径
func getAssetPathByNodeId(nodeId string, nodeIdToNodeDataMap map[string]NodeData) string {
	var nodeData NodeData
	var ok bool
	var assetPathRet string
	nodeData, ok = nodeIdToNodeDataMap[nodeId]
	if !ok {
		return assetPathRet
	}

	var loopCount = 0
	for {
		if len(nodeData.InstanceId) <= 0 {
			return assetPathRet
		}
		assetPathRet = MODEL_PATH_SEP + nodeData.InstanceId + assetPathRet

		if len(nodeData.ParentId) <= 0 {
			return assetPathRet
		}

		nodeData, ok = nodeIdToNodeDataMap[nodeData.ParentId]
		if !ok {
			return assetPathRet
		}

		//防止死循环，最多循环20次
		loopCount++
		if loopCount > 20 {
			return assetPathRet
		}
	}
}
