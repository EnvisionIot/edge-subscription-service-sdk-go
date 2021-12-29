package metadata

import (
	"errors"
	"fmt"
	"strconv"
)

/**
 * Description:
 *
 * @author yang.zhang4
 * @date 2020/12/17
 */

//单个mapping的一些信息
type MappingInfo struct {
	Mapping           Mapping         //来源于DataAccess
	UsingOMETimestamp bool            //来源于DataAccess
	ModelPointType    string          //表明这个mapping是测点还是事件
	MeasurePoint      TSLMeasurePoint //来源于ModelAsset
	Event             TSLEvent        //来源于ModelAsset
	ModelIdPath       string          //来源于ModelAsset
}

type MetaDataBase struct {
	DataAccess *DataAccess
	ModelAsset *ModelAsset
}

type MetaDataAdvance struct {
	MappingIdToMappingInfoMap map[int64]*MappingInfo
	//map[CollectDeviceId]map[模型点标识符]MappingId,这里只考虑采集mapping
	CollectDeviceIdPointIdToMappingIdNonControlMap map[string]map[string]int64
	//map[AssetId]map[模型点标识符]CollectDeviceId,这里只考虑采集mapping
	AssetIdPointIdToCollectDeviceIdNonControlMap map[string]map[string]string
	//map[CollectDeviceId]map[模型点标识符]MappingId,这里只考虑控制mapping
	CollectDeviceIdPointIdToMappingIdControlMap map[string]map[string]int64
	//map[AssetId]map[模型点标识符]CollectDeviceId,这里只考虑控制mapping
	AssetIdPointIdToCollectDeviceIdControlMap map[string]map[string]string
	//map[assetId]map[ConnectStatPointId]"1"
	AssetIdToConnectStatPointIdMap map[string]map[string]string
	//map[collectDeviceId]map[ConnectStatPointId]"1"
	CollectDeviceIdToConnectStatPointIdMap map[string]map[string]string
}

type MetaData struct {
	Base    *MetaDataBase
	Advance *MetaDataAdvance
}

type MetaDataConfig struct {
	DataAccessCfg DataAccessConfig
	ModelAssetCfg ModelAssetConfig
}

type MetaDataWork struct {
	Config MetaDataConfig
}

//生成base
func (f *MetaDataWork) GetBaseForMetaData() (*MetaDataBase, error) {
	var base MetaDataBase
	var err error
	var dataAccess = DataAccessWork{
		Config: f.Config.DataAccessCfg,
	}

	var modelAsset = ModelAssetWork{
		Config: f.Config.ModelAssetCfg,
	}

	base.DataAccess, err = dataAccess.GetDataAccess()
	if err != nil {
		return nil, fmt.Errorf("getDataAccess, %s", err)
	}

	base.ModelAsset, err = modelAsset.GetModelAsset()
	if err != nil {
		return nil, fmt.Errorf("getModelAsset, %s", err)
	}

	return &base, nil
}

//从base生成advance
func (f *MetaDataWork) GetAdvanceFromBaseForMetaData(base *MetaDataBase) (*MetaDataAdvance, error) {
	if base == nil || base.DataAccess == nil || base.ModelAsset == nil {
		return nil, errors.New("base is nil error")
	}

	var ok bool
	var advance MetaDataAdvance
	var err error
	var stringTemp string

	advance.MappingIdToMappingInfoMap = make(map[int64]*MappingInfo)
	for _, mapping := range base.DataAccess.Base.MappingIdToMappingMap {
		var mappingInfo MappingInfo
		mappingInfo.Mapping = mapping

		stringTemp, _ = mapping.Attributes[USING_OME_TIMESTAMP]
		mappingInfo.UsingOMETimestamp, _ = strconv.ParseBool(stringTemp)

		var tslModelData TSLModelData
		tslModelData, ok = base.ModelAsset.Base.TSLModelDataCache[mappingInfo.Mapping.ModelId]
		if !ok {
			return nil, fmt.Errorf("no tslModelData, modelId=%s", mappingInfo.Mapping.ModelId)
		}

		mappingInfo.MeasurePoint, ok = tslModelData.Measurepoints[mappingInfo.Mapping.FeatureName]
		if ok {
			mappingInfo.ModelPointType = MODEL_POINT_MEASUREPOINT
		} else {
			mappingInfo.Event, ok = tslModelData.Events[mappingInfo.Mapping.FeatureName]
			if ok {
				mappingInfo.ModelPointType = MODEL_POINT_EVENT
			} else {
				return nil, fmt.Errorf("no such point, modelId=%s, pointId=%s", mappingInfo.Mapping.ModelId, mappingInfo.Mapping.FeatureName)
			}
		}

		mappingInfo.ModelIdPath = tslModelData.ModelIdPath

		advance.MappingIdToMappingInfoMap[mappingInfo.Mapping.Id] = &mappingInfo
	}

	advance.CollectDeviceIdPointIdToMappingIdNonControlMap = make(map[string]map[string]int64)
	for collectDeviceId, collectDevice := range base.DataAccess.Base.CollectDeviceIdToCollectDeviceMap {
		var pointIdToMappingIdNonControlMap map[string]int64
		pointIdToMappingIdNonControlMap, ok = advance.CollectDeviceIdPointIdToMappingIdNonControlMap[collectDeviceId]
		if !ok {
			pointIdToMappingIdNonControlMap = make(map[string]int64)
		}

		var simpleModelData SimpleModelData
		simpleModelData, ok = base.ModelAsset.Advance.AssetIdToSimpleModelData[collectDevice.AssetId]
		if !ok {
			return nil, fmt.Errorf("can not find model, assetId=%s", collectDevice.AssetId)
		}

		var tslModelData TSLModelData
		tslModelData, ok = base.ModelAsset.Base.TSLModelDataCache[simpleModelData.ModelId]
		if !ok {
			return nil, fmt.Errorf("can not find model, modelId=%s", simpleModelData.ModelId)
		}
		for measurepointId := range tslModelData.Measurepoints {
			var mappingId int64
			mappingId, err = getMappingIdByCollectDeviceIdAndPointIdForInit(collectDeviceId, measurepointId, base)
			if err != nil {
				continue
			}
			pointIdToMappingIdNonControlMap[measurepointId] = mappingId
		}
		for eventId := range tslModelData.Events {
			var mappingId int64
			mappingId, err = getMappingIdByCollectDeviceIdAndPointIdForInit(collectDeviceId, eventId, base)
			if err != nil {
				continue
			}
			pointIdToMappingIdNonControlMap[eventId] = mappingId
		}
		advance.CollectDeviceIdPointIdToMappingIdNonControlMap[collectDeviceId] = pointIdToMappingIdNonControlMap
	}

	advance.AssetIdPointIdToCollectDeviceIdNonControlMap = make(map[string]map[string]string)
	for assetId, simpleModelData := range base.ModelAsset.Advance.AssetIdToSimpleModelData {
		var pointIdToCollectDeviceIdNonControlMap map[string]string
		pointIdToCollectDeviceIdNonControlMap, ok = advance.AssetIdPointIdToCollectDeviceIdNonControlMap[assetId]
		if !ok {
			pointIdToCollectDeviceIdNonControlMap = make(map[string]string)
		}

		var tslModelData TSLModelData
		tslModelData, ok = base.ModelAsset.Base.TSLModelDataCache[simpleModelData.ModelId]
		if !ok {
			return nil, fmt.Errorf("can not find model, modelId=%s", simpleModelData.ModelId)
		}
		for measurepointId := range tslModelData.Measurepoints {
			var collectDeviceId string
			collectDeviceId, err = getCollectDeviceIdByAssetIdAndPointIdForInit(assetId, measurepointId, base)
			if err != nil {
				continue
			}
			pointIdToCollectDeviceIdNonControlMap[measurepointId] = collectDeviceId
		}
		for eventId := range tslModelData.Events {
			var collectDeviceId string
			collectDeviceId, err = getCollectDeviceIdByAssetIdAndPointIdForInit(assetId, eventId, base)
			if err != nil {
				continue
			}
			pointIdToCollectDeviceIdNonControlMap[eventId] = collectDeviceId
		}
		advance.AssetIdPointIdToCollectDeviceIdNonControlMap[assetId] = pointIdToCollectDeviceIdNonControlMap
	}

	advance.CollectDeviceIdPointIdToMappingIdControlMap = make(map[string]map[string]int64)
	for collectDeviceId, collectDevice := range base.DataAccess.Base.CollectDeviceIdToCollectDeviceMap {
		var pointIdToMappingIdControlMap map[string]int64
		pointIdToMappingIdControlMap, ok = advance.CollectDeviceIdPointIdToMappingIdControlMap[collectDeviceId]
		if !ok {
			pointIdToMappingIdControlMap = make(map[string]int64)
		}

		var simpleModelData SimpleModelData
		simpleModelData, ok = base.ModelAsset.Advance.AssetIdToSimpleModelData[collectDevice.AssetId]
		if !ok {
			return nil, fmt.Errorf("can not find model, assetId=%s", collectDevice.AssetId)
		}

		var tslModelData TSLModelData
		tslModelData, ok = base.ModelAsset.Base.TSLModelDataCache[simpleModelData.ModelId]
		if !ok {
			return nil, fmt.Errorf("can not find model, modelId=%s", simpleModelData.ModelId)
		}
		for measurepointId := range tslModelData.Measurepoints {
			var mappingId int64
			mappingId, err = getMappingIdByCollectDeviceIdAndPointIdForControlForInit(collectDeviceId, measurepointId, base)
			if err != nil {
				continue
			}
			pointIdToMappingIdControlMap[measurepointId] = mappingId
		}
		for serviceId := range tslModelData.Services {
			var mappingId int64
			mappingId, err = getMappingIdByCollectDeviceIdAndPointIdForControlForInit(collectDeviceId, serviceId, base)
			if err != nil {
				continue
			}
			pointIdToMappingIdControlMap[serviceId] = mappingId
		}
		for eventId := range tslModelData.Events {
			var mappingId int64
			mappingId, err = getMappingIdByCollectDeviceIdAndPointIdForControlForInit(collectDeviceId, eventId, base)
			if err != nil {
				continue
			}
			pointIdToMappingIdControlMap[eventId] = mappingId
		}
		advance.CollectDeviceIdPointIdToMappingIdControlMap[collectDeviceId] = pointIdToMappingIdControlMap
	}

	advance.AssetIdPointIdToCollectDeviceIdControlMap = make(map[string]map[string]string)
	for assetId, simpleModelData := range base.ModelAsset.Advance.AssetIdToSimpleModelData {
		var pointIdToCollectDeviceIdControlMap map[string]string
		pointIdToCollectDeviceIdControlMap, ok = advance.AssetIdPointIdToCollectDeviceIdControlMap[assetId]
		if !ok {
			pointIdToCollectDeviceIdControlMap = make(map[string]string)
		}

		var tslModelData TSLModelData
		tslModelData, ok = base.ModelAsset.Base.TSLModelDataCache[simpleModelData.ModelId]
		if !ok {
			return nil, fmt.Errorf("can not find model, modelId=%s", simpleModelData.ModelId)
		}
		for measurepointId := range tslModelData.Measurepoints {
			var collectDeviceId string
			collectDeviceId, err = getCollectDeviceIdByAssetIdAndPointIdForControlForInit(assetId, measurepointId, base)
			if err != nil {
				continue
			}
			pointIdToCollectDeviceIdControlMap[measurepointId] = collectDeviceId
		}
		for serviceId := range tslModelData.Services {
			var collectDeviceId string
			collectDeviceId, err = getCollectDeviceIdByAssetIdAndPointIdForControlForInit(assetId, serviceId, base)
			if err != nil {
				continue
			}
			pointIdToCollectDeviceIdControlMap[serviceId] = collectDeviceId
		}
		for eventId := range tslModelData.Events {
			var collectDeviceId string
			collectDeviceId, err = getCollectDeviceIdByAssetIdAndPointIdForControlForInit(assetId, eventId, base)
			if err != nil {
				continue
			}
			pointIdToCollectDeviceIdControlMap[eventId] = collectDeviceId
		}
		advance.AssetIdPointIdToCollectDeviceIdControlMap[assetId] = pointIdToCollectDeviceIdControlMap
	}

	advance.AssetIdToConnectStatPointIdMap = make(map[string]map[string]string)
	for assetId, simpleModelData := range base.ModelAsset.Advance.AssetIdToSimpleModelData {
		var connectStatPointIdMap map[string]string
		connectStatPointIdMap, ok = advance.AssetIdToConnectStatPointIdMap[assetId]
		if !ok {
			connectStatPointIdMap = make(map[string]string)
		}

		var tslModelData TSLModelData
		tslModelData, ok = base.ModelAsset.Base.TSLModelDataCache[simpleModelData.ModelId]
		if !ok {
			return nil, fmt.Errorf("can not find model, modelId=%s", simpleModelData.ModelId)
		}

		var pointIdList []string
		for measurepointId := range tslModelData.Measurepoints {
			pointIdList = append(pointIdList, measurepointId)
		}
		for event := range tslModelData.Events {
			pointIdList = append(pointIdList, event)
		}
		for _, pointId := range pointIdList {
			var mapping Mapping
			mapping, err = getMappingByAssetIdAndPointIdForInit(assetId, pointId, base)
			if err != nil {
				continue
			}
			if mapping.MappingRule.MappingFxStr != CONNECT_STATE {
				continue
			}
			connectStatPointIdMap[pointId] = "1"
		}

		advance.AssetIdToConnectStatPointIdMap[assetId] = connectStatPointIdMap
	}

	advance.CollectDeviceIdToConnectStatPointIdMap = make(map[string]map[string]string)
	for collectDeviceId, collectDevice := range base.DataAccess.Base.CollectDeviceIdToCollectDeviceMap {
		var connectStatPointIdMap map[string]string
		connectStatPointIdMap, ok = advance.CollectDeviceIdToConnectStatPointIdMap[collectDeviceId]
		if !ok {
			connectStatPointIdMap = make(map[string]string)
		}

		var simpleModelData SimpleModelData
		simpleModelData, ok = base.ModelAsset.Advance.AssetIdToSimpleModelData[collectDevice.AssetId]
		if !ok {
			return nil, fmt.Errorf("can not find model, assetId=%s", collectDevice.AssetId)
		}

		var tslModelData TSLModelData
		tslModelData, ok = base.ModelAsset.Base.TSLModelDataCache[simpleModelData.ModelId]
		if !ok {
			return nil, fmt.Errorf("can not find model, modelId=%s", simpleModelData.ModelId)
		}

		var pointIdList []string
		for measurepointId := range tslModelData.Measurepoints {
			pointIdList = append(pointIdList, measurepointId)
		}
		for event := range tslModelData.Events {
			pointIdList = append(pointIdList, event)
		}
		for _, pointId := range pointIdList {
			var mapping Mapping
			mapping, err = getMappingByCollectDeviceIdAndPointIdForInit(collectDeviceId, pointId, base)
			if err != nil {
				continue
			}
			if mapping.MappingRule.MappingFxStr != CONNECT_STATE {
				continue
			}
			connectStatPointIdMap[pointId] = "1"
		}

		advance.CollectDeviceIdToConnectStatPointIdMap[collectDeviceId] = connectStatPointIdMap
	}

	return &advance, nil
}

//生成metadata
func (f *MetaDataWork) GetMetaData() (*MetaData, error) {
	var data MetaData
	var err error
	data.Base, err = f.GetBaseForMetaData()
	if err != nil {
		return nil, fmt.Errorf("GetBaseForMetaData, %s", err)
	}

	data.Advance, err = f.GetAdvanceFromBaseForMetaData(data.Base)
	if err != nil {
		return nil, fmt.Errorf("getAdvanceFromBaseForMetaData, %s", err)
	}

	return &data, nil
}

func CheckMetaDataIsValid(metaData *MetaData) bool {
	if metaData == nil || metaData.Base == nil || metaData.Advance == nil {
		return false
	}

	if !CheckDataAccessIsValid(metaData.Base.DataAccess) {
		return false
	}

	if !CheckModelAssetIsValid(metaData.Base.ModelAsset) {
		return false
	}

	return true
}

func GetOrgId(metaData *MetaData) (string, error) {
	if !CheckMetaDataIsValid(metaData) {
		return "", ERROR_RESOURCE_IS_NIL
	}

	return metaData.Base.DataAccess.Base.OrgId, nil
}

//根据assetId得到配置了CONNECT_STATE公式的点的map
func GetConnectStateMapByAssetId(assetId string, metaData *MetaData) (map[string]string, error) {
	if !CheckMetaDataIsValid(metaData) {
		return nil, errors.New("metaData is invalid")
	}

	var ok bool
	var connectStatPointIdMap map[string]string
	connectStatPointIdMap, ok = metaData.Advance.AssetIdToConnectStatPointIdMap[assetId]
	if !ok {
		return nil, errors.New("no connectStatPointIdMap")
	}

	return connectStatPointIdMap, nil
}

//根据collectDeviceId得到配置了CONNECT_STATE公式的点的map
func GetConnectStateMapByCollectDeviceId(collectDeviceId string, metaData *MetaData) (map[string]string, error) {
	if !CheckMetaDataIsValid(metaData) {
		return nil, errors.New("metaData is invalid")
	}

	var ok bool
	var connectStatPointIdMap map[string]string
	connectStatPointIdMap, ok = metaData.Advance.CollectDeviceIdToConnectStatPointIdMap[collectDeviceId]
	if !ok {
		return nil, errors.New("no connectStatPointIdMap")
	}

	return connectStatPointIdMap, nil
}

//根据collectDeviceId和pointId判断这个点需不需要oem方式上送
func CollectDeviceIdPointIdNeedOem(collectDeviceId string, pointId string, metaData *MetaData) bool {
	if len(collectDeviceId) <= 0 {
		return true
	}

	if !CheckMetaDataIsValid(metaData) {
		return false
	}

	var mappingId int64
	var ok bool
	var err error
	mappingId, err = GetMappingIdByCollectDeviceIdAndPointId(collectDeviceId, pointId, metaData)
	if err != nil {
		return true
	}

	var mappingInfo *MappingInfo
	mappingInfo, ok = metaData.Advance.MappingIdToMappingInfoMap[mappingId]
	if !ok {
		return true
	}

	if mappingInfo.UsingOMETimestamp {
		return true
	} else {
		return false
	}
}

//根据collectDeviceId和pointId判断这个点需不需要oem方式上送
func CollectDeviceIdPointIdNeedOemForInt32(collectDeviceId string, pointId string, metaData *MetaData) int32 {
	var useOemBool = CollectDeviceIdPointIdNeedOem(collectDeviceId, pointId, metaData)
	if useOemBool {
		return 1
	} else {
		return 0
	}
}

func GetDataTypeAndSubDataType(assetId string, pointId string, metaData *MetaData) (int32, string, string, error) {
	if metaData == nil || metaData.Base == nil || metaData.Advance == nil {
		return MODEL_POINT_MEASUREPOINT_INT32, "", "", fmt.Errorf("resource is nil error")
	}

	var tslInstanceData TSLInstanceData
	var ok bool
	tslInstanceData, ok = metaData.Base.ModelAsset.Base.TSLInstanceDataCache[assetId]
	if !ok {
		return MODEL_POINT_MEASUREPOINT_INT32, "", "", fmt.Errorf("can not find asset, assetId=%s", assetId)
	}

	var tslModelData TSLModelData
	tslModelData, ok = metaData.Base.ModelAsset.Base.TSLModelDataCache[tslInstanceData.ModelId]
	if !ok {
		return MODEL_POINT_MEASUREPOINT_INT32, "", "", fmt.Errorf("can not find model, modelId=%s", tslInstanceData.ModelId)
	}

	var tslMeasurePoint TSLMeasurePoint
	tslMeasurePoint, ok = tslModelData.Measurepoints[pointId]
	if ok {
		return MODEL_POINT_MEASUREPOINT_INT32, tslMeasurePoint.DataType, tslMeasurePoint.InnerType, nil
	} else {
		_, ok = tslModelData.Events[pointId]
		if !ok {
			return MODEL_POINT_MEASUREPOINT_INT32, "", "", fmt.Errorf("can not find point, pointId=%s", pointId)
		}
		return MODEL_POINT_EVENT_INT32, VALUE_TYPE_STRUCT, "", nil
	}
}

//根据assetId和pointId得到这个点的MappingRule
func GetMappingRuleByAssetIdAndPointId(assetId string, pointId string, metaData *MetaData) (MappingRule, error) {
	var mappingRule MappingRule
	if !CheckMetaDataIsValid(metaData) {
		return mappingRule, errors.New("metaData is invalid")
	}

	var collectDeviceId string
	collectDeviceId, _ = GetCollectDeviceIdByAssetIdAndPointId(assetId, pointId, metaData)
	if len(collectDeviceId) <= 0 {
		return mappingRule, errors.New("can not find collectDeviceId")
	}

	var mappingId int64
	var ok bool
	var err error
	mappingId, err = GetMappingIdByCollectDeviceIdAndPointId(collectDeviceId, pointId, metaData)
	if err != nil {
		return mappingRule, err
	}

	var mapping Mapping
	mapping, ok = metaData.Base.DataAccess.Base.MappingIdToMappingMap[mappingId]
	if !ok {
		return mappingRule, errors.New("can not find Mapping")
	}

	return mapping.MappingRule, nil
}

//根据assetId和pointId得到collectDeviceId，采集点
func GetCollectDeviceIdByAssetIdAndPointId(assetId string, pointId string, metaData *MetaData) (string, error) {
	if !CheckMetaDataIsValid(metaData) {
		return "", errors.New("metaData is invalid")
	}
	var collectDeviceId string
	var collectDeviceIdMap map[string]string
	var ok bool
	collectDeviceId, ok = metaData.Advance.AssetIdPointIdToCollectDeviceIdNonControlMap[assetId][pointId]
	if !ok {
		//这个点没有做过采集mapping，或者这个点就不存在
		collectDeviceIdMap, ok = metaData.Base.DataAccess.Advance.AssetIdToCollectDeviceIdMap[assetId]
		if (!ok) || (len(collectDeviceIdMap) <= 0) {
			collectDeviceId = ""
		} else {
			for collectDeviceId = range collectDeviceIdMap {
				break
			}
		}
	}
	return collectDeviceId, nil
}

//根据assetId和pointId得到collectDeviceId，控制点
func GetCollectDeviceIdByAssetIdAndPointIdForControl(assetId string, pointId string, metaData *MetaData) (string, error) {
	if !CheckMetaDataIsValid(metaData) {
		return "", errors.New("metaData is invalid")
	}
	var collectDeviceId string
	var collectDeviceIdMap map[string]string
	var ok bool
	collectDeviceId, ok = metaData.Advance.AssetIdPointIdToCollectDeviceIdControlMap[assetId][pointId]
	if !ok {
		//这个点没有做过控制mapping，或者这个点就不存在
		collectDeviceIdMap, ok = metaData.Base.DataAccess.Advance.AssetIdToCollectDeviceIdMap[assetId]
		if (!ok) || (len(collectDeviceIdMap) <= 0) {
			collectDeviceId = ""
		} else {
			for collectDeviceId = range collectDeviceIdMap {
				break
			}
		}
	}
	return collectDeviceId, nil
}

//根据collectDeviceId和devicePointName（设备点，如ai.1 ai.2）得到mappingId的数组
func GetMappingIdsByCollectDeviceIdAndPointName(collectDeviceId string, pointName string, metaData *MetaData) ([]int64, error) {
	if !CheckMetaDataIsValid(metaData) {
		return nil, errors.New("metaData is invalid")
	}

	var templateId int64
	var ok bool
	templateId, ok = metaData.Base.DataAccess.Advance.CollectDeviceIdToTemplateIdMap[collectDeviceId]
	if !ok {
		return nil, errors.New("can not find templateId")
	}

	var mappingIdArray []int64
	mappingIdArray, ok = metaData.Base.DataAccess.Advance.TemplateIdToDevicePointNameMappingMap[templateId][pointName]
	if !ok || len(mappingIdArray) <= 0 {
		return nil, errors.New("can not find mapping")
	}

	return mappingIdArray, nil
}

//根据collectDeviceId和pointId找到mappingId，采集点
func GetMappingIdByCollectDeviceIdAndPointId(collectDeviceId string, pointId string, metaData *MetaData) (int64, error) {
	if !CheckMetaDataIsValid(metaData) {
		return -1, errors.New("metaData is invalid")
	}

	var mappingId int64
	var ok bool
	mappingId, ok = metaData.Advance.CollectDeviceIdPointIdToMappingIdNonControlMap[collectDeviceId][pointId]
	if !ok {
		return -1, errors.New("can not find mappingId")
	}

	return mappingId, nil
}

//根据collectDeviceId和pointId找到mappingId，控制点
func GetMappingIdByCollectDeviceIdAndPointIdForControl(collectDeviceId string, pointId string, metaData *MetaData) (int64, error) {
	if !CheckMetaDataIsValid(metaData) {
		return -1, errors.New("metaData is invalid")
	}

	var mappingId int64
	var ok bool
	mappingId, ok = metaData.Advance.CollectDeviceIdPointIdToMappingIdControlMap[collectDeviceId][pointId]
	if !ok {
		return -1, errors.New("can not find mappingId")
	}

	return mappingId, nil
}

//根据assetId和pointId找到mappingId，控制点
func GetMappingIdByAssetIdAndPointIdForControl(assetId string, pointId string, metaData *MetaData) (int64, error) {
	if !CheckMetaDataIsValid(metaData) {
		return -1, errors.New("metaData is invalid")
	}

	var collectDeviceId string
	var err error
	collectDeviceId, err = GetCollectDeviceIdByAssetIdAndPointIdForControl(assetId, pointId, metaData)
	if err != nil {
		return -1, err
	}

	var mappingId int64
	mappingId, err = GetMappingIdByCollectDeviceIdAndPointIdForControl(collectDeviceId, pointId, metaData)
	if err != nil {
		return -1, err
	}

	return mappingId, nil
}

//初始化时使用，根据assetId和pointId得到mapping，采集点
func getMappingByAssetIdAndPointIdForInit(assetId string, pointId string, metaDataBase *MetaDataBase) (Mapping, error) {
	if metaDataBase == nil || metaDataBase.DataAccess == nil || metaDataBase.ModelAsset == nil {
		return Mapping{}, errors.New("resource nil error")
	}

	var collectDeviceId string
	var err error
	collectDeviceId, err = getCollectDeviceIdByAssetIdAndPointIdForInit(assetId, pointId, metaDataBase)
	if err != nil {
		return Mapping{}, fmt.Errorf("getCollectDeviceIdByAssetIdAndPointIdForInit, %s", err)
	}

	return getMappingByCollectDeviceIdAndPointIdForInit(collectDeviceId, pointId, metaDataBase)
}

//初始化时使用，根据collectDeviceId和pointId得到mapping，采集点
func getMappingByCollectDeviceIdAndPointIdForInit(collectDeviceId string, pointId string, metaDataBase *MetaDataBase) (Mapping, error) {
	if metaDataBase == nil || metaDataBase.DataAccess == nil || metaDataBase.ModelAsset == nil {
		return Mapping{}, errors.New("resource nil error")
	}

	var err error
	var mappingId int64
	mappingId, err = getMappingIdByCollectDeviceIdAndPointIdForInit(collectDeviceId, pointId, metaDataBase)
	if err != nil {
		return Mapping{}, fmt.Errorf("getMappingIdByCollectDeviceIdAndPointIdForInit, %s", err)
	}
	var mapping Mapping
	var ok bool
	mapping, ok = metaDataBase.DataAccess.Base.MappingIdToMappingMap[mappingId]
	if !ok {
		return Mapping{}, errors.New("can not find Mapping")
	}

	return mapping, nil
}

//初始化时使用，根据assetId和pointId得到collectDeviceId，采集点
func getCollectDeviceIdByAssetIdAndPointIdForInit(assetId string, pointId string, metaDataBase *MetaDataBase) (string, error) {
	if metaDataBase == nil || metaDataBase.DataAccess == nil || metaDataBase.ModelAsset == nil {
		return "", fmt.Errorf("resource nil error")
	}

	var collectDeviceIdMap map[string]string
	var ok bool
	var err error
	collectDeviceIdMap, ok = metaDataBase.DataAccess.Advance.AssetIdToCollectDeviceIdMap[assetId]
	if !ok {
		return "", fmt.Errorf("can not find collectDeviceIdMap")
	}

	for collectDeviceId := range collectDeviceIdMap {
		var mappingId int64
		mappingId, err = getMappingIdByCollectDeviceIdAndPointIdForInit(collectDeviceId, pointId, metaDataBase)
		if err != nil {
			continue
		}
		var mapping Mapping
		mapping, ok = metaDataBase.DataAccess.Base.MappingIdToMappingMap[mappingId]
		if !ok {
			continue
		}
		if mapping.MappingRule.MappingFxStr == MAPPING_NO_MAPPING || mapping.MappingRule.MappingFxStr == MAPPING_INVALID {
			continue
		}
		return collectDeviceId, nil
	}

	return "", fmt.Errorf("can not find collectDeviceId")
}

//初始化时使用，根据assetId和pointId得到collectDeviceId，控制点
func getCollectDeviceIdByAssetIdAndPointIdForControlForInit(assetId string, pointId string, metaDataBase *MetaDataBase) (string, error) {
	if metaDataBase == nil || metaDataBase.DataAccess == nil || metaDataBase.ModelAsset == nil {
		return "", fmt.Errorf("resource nil error")
	}

	var collectDeviceIdMap map[string]string
	var ok bool
	var err error
	collectDeviceIdMap, ok = metaDataBase.DataAccess.Advance.AssetIdToCollectDeviceIdMap[assetId]
	if !ok {
		return "", fmt.Errorf("can not find collectDeviceIdMap")
	}

	for collectDeviceId := range collectDeviceIdMap {
		var mappingId int64
		mappingId, err = getMappingIdByCollectDeviceIdAndPointIdForControlForInit(collectDeviceId, pointId, metaDataBase)
		if err != nil {
			continue
		}
		var mapping Mapping
		mapping, ok = metaDataBase.DataAccess.Base.MappingIdToMappingMap[mappingId]
		if !ok {
			continue
		}
		if mapping.MappingRule.MappingFxStr == MAPPING_NO_MAPPING || mapping.MappingRule.MappingFxStr == MAPPING_INVALID {
			continue
		}
		return collectDeviceId, nil
	}

	return "", fmt.Errorf("can not find collectDeviceId")
}

//初始化时使用，根据collectDeviceId和pointId找到mappingId，采集点
func getMappingIdByCollectDeviceIdAndPointIdForInit(collectDeviceId string, pointId string, metaDataBase *MetaDataBase) (int64, error) {
	if metaDataBase == nil || metaDataBase.DataAccess == nil || metaDataBase.ModelAsset == nil {
		return -1, errors.New("resource nil error")
	}

	var templateId int64
	var ok bool
	templateId, ok = metaDataBase.DataAccess.Advance.CollectDeviceIdToTemplateIdMap[collectDeviceId]
	if !ok {
		return -1, errors.New("can not find templateId")
	}

	var template Template
	template, ok = metaDataBase.DataAccess.Base.TemplateIdToTemplateMap[templateId]
	if !ok {
		return -1, errors.New("can not find template")
	}

	var deviceLogicalId string
	if template.TemplateType == TEMPLATE_TYPE_DYNAMIC {
		var collectDevice CollectDevice
		collectDevice, ok = metaDataBase.DataAccess.Base.CollectDeviceIdToCollectDeviceMap[collectDeviceId]
		if !ok {
			return -1, errors.New("can not find collectDevice")
		}
		deviceLogicalId = collectDevice.AssetId
	}

	var mappingId int64
	mappingId, ok = metaDataBase.DataAccess.Advance.TemplateIdFeatureNameDeviceLogicalIdToMappingIdNonControlMap[templateId][pointId][deviceLogicalId]
	if !ok {
		return -1, errors.New("can not find mappingId")
	}

	return mappingId, nil
}

//初始化时使用，根据collectDeviceId和pointId找到mappingId，控制点
func getMappingIdByCollectDeviceIdAndPointIdForControlForInit(collectDeviceId string, pointId string, metaDataBase *MetaDataBase) (int64, error) {
	if metaDataBase == nil || metaDataBase.DataAccess == nil || metaDataBase.ModelAsset == nil {
		return -1, errors.New("resource nil error")
	}

	var templateId int64
	var ok bool
	templateId, ok = metaDataBase.DataAccess.Advance.CollectDeviceIdToTemplateIdMap[collectDeviceId]
	if !ok {
		return -1, errors.New("can not find templateId")
	}

	var template Template
	template, ok = metaDataBase.DataAccess.Base.TemplateIdToTemplateMap[templateId]
	if !ok {
		return -1, errors.New("can not find template")
	}

	var deviceLogicalId string
	if template.TemplateType == TEMPLATE_TYPE_DYNAMIC {
		var collectDevice CollectDevice
		collectDevice, ok = metaDataBase.DataAccess.Base.CollectDeviceIdToCollectDeviceMap[collectDeviceId]
		if !ok {
			return -1, errors.New("can not find collectDevice")
		}
		deviceLogicalId = collectDevice.AssetId
	}

	var mappingId int64
	mappingId, ok = metaDataBase.DataAccess.Advance.TemplateIdFeatureNameDeviceLogicalIdToMappingIdControlMap[templateId][pointId][deviceLogicalId]
	if !ok {
		return -1, errors.New("can not find mappingId")
	}

	return mappingId, nil
}
