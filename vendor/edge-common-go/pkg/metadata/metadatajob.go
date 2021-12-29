package metadata

import (
	"edge-common-go/pkg/log"
	"edge-common-go/pkg/utils"
	"errors"
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

/**
 * Description:
 *
 * @author yang.zhang4
 * @date 2020/12/18
 */

const (
	SUCCESS = "success"
	FAIL    = "fail"
)

var metaDataJobGlobal *MetaDataJob

type MetaDataJobConfig struct {
	//请求http url的超时时间
	WebOpTimeoutMs int64

	//获取设备接入相关的元数据需要的url
	GetOrgIdUrl          string
	GetCollectDevicesUrl string
	GetTemplatesUrl      string
	GetMappingsUrl       string

	//获取模型资产相关的元数据需要的url
	GetTSLInstancesUrl string
	GetThingModelsUrl  string
	GetNodesUrl        string
	GetTreesUrl        string

	//日志
	LogArray []*log.LogAa

	//用于每次更新后，发送更新成功或失败消息
	NotifyMsgChan chan string
}

type MetaDataJob struct {
	configInput  MetaDataJobConfig
	config       MetaDataConfig
	metaDataWork MetaDataWork

	srcFilename string

	needUpdate int32 //是否需要更新订阅缓存，>0时需要更新

	metaDataMutex sync.RWMutex //操作metaData的锁
	metaData      *MetaData
}

func NewMetaDataJob(config MetaDataJobConfig) (*MetaDataJob, error) {
	if metaDataJobGlobal != nil {
		return metaDataJobGlobal, nil
	}

	var err error
	err = CheckConfig(config)
	if err != nil {
		return nil, fmt.Errorf("checkConfig, %s", err)
	}

	info := MetaDataJob{}

	//初始化SrcFilename
	info.srcFilename = utils.GetFileNameLastN(1)

	info.configInput = config
	info.getConfigFromInput(info.configInput)

	metaDataJobGlobal = &info

	return metaDataJobGlobal, nil
}

//获取metadatajob的实例
func GetMetaDataJob() *MetaDataJob {
	return metaDataJobGlobal
}

func (f *MetaDataJob) StartAll() {
	go f.dealUpdateJob()
}

//获取当前断面metadata
func GetMetaData() (*MetaData, error) {
	return metaDataJobGlobal.GetMetaData()
}

//获取当前断面metadata
func (f *MetaDataJob) GetMetaData() (*MetaData, error) {
	f.metaDataMutex.RLock()
	defer f.metaDataMutex.RUnlock()

	return f.metaData, nil
}

//检查输入config是否合法
func CheckConfig(configInput MetaDataJobConfig) error {
	if configInput.WebOpTimeoutMs < 0 {
		return errors.New("webOpTimeoutMs is illegal")
	}

	if len(configInput.GetOrgIdUrl) <= 0 {
		return errors.New("getOrgIdUrl is illegal")
	}

	if len(configInput.GetCollectDevicesUrl) <= 0 {
		return errors.New("getCollectDevicesUrl is illegal")
	}

	if len(configInput.GetTemplatesUrl) <= 0 {
		return errors.New("getTemplatesUrl is illegal")
	}

	if len(configInput.GetMappingsUrl) <= 0 {
		return errors.New("getMappingsUrl is illegal")
	}

	if len(configInput.GetTSLInstancesUrl) <= 0 {
		return errors.New("getTSLInstancesUrl is illegal")
	}

	if len(configInput.GetThingModelsUrl) <= 0 {
		return errors.New("getThingModelsUrl is illegal")
	}

	if len(configInput.GetNodesUrl) <= 0 {
		return errors.New("getNodesUrl is illegal")
	}

	if len(configInput.GetTreesUrl) <= 0 {
		return errors.New("getTreesUrl is illegal")
	}

	return nil
}

//刷新当前断面metadata（同步方式，对外）
func (f *MetaDataJob) RefreshMetaDataSync() error {
	var err error
	err = f.refreshMetaDataSync()
	if err != nil {
		f.sendFail()
		return err
	}

	f.sendSuccess()
	return nil
}

//刷新当前断面metadata直至成功（同步方式）
func (f *MetaDataJob) RefreshMetaDataSyncUntilSuccess() {
	var err error
	var ticker *time.Ticker
	ticker = time.NewTicker(time.Second * 5)

	for {
		select {
		case <-ticker.C:
			err = f.RefreshMetaDataSync()
			if err != nil {
				log.MultiLogErrorf(f.configInput.LogArray, "[%s:%d]:refreshMetaDataSync, %s", f.srcFilename, utils.GetLineNum(), err)
			} else {
				ticker.Stop()
				log.MultiLogInfof(f.configInput.LogArray, "[%s:%d]:refreshMetaDataSync success", f.srcFilename, utils.GetLineNum())
				return
			}
		}
	}
}

//刷新当前断面metadata（异步方式，会汇集一小段时间的更新需求，调用多次这个函数，只做一次更新，主要用于防抖）
func (f *MetaDataJob) RefreshMetaDataASync() {
	atomic.StoreInt32(&f.needUpdate, 1)
}

func (f *MetaDataJob) getConfigFromInput(configInput MetaDataJobConfig) {
	f.config.DataAccessCfg.WebOpTimeoutMs = configInput.WebOpTimeoutMs
	f.config.DataAccessCfg.GetOrgIdUrl = configInput.GetOrgIdUrl
	f.config.DataAccessCfg.GetCollectDevicesUrl = configInput.GetCollectDevicesUrl
	f.config.DataAccessCfg.GetTemplatesUrl = configInput.GetTemplatesUrl
	f.config.DataAccessCfg.GetMappingsUrl = configInput.GetMappingsUrl

	f.config.ModelAssetCfg.WebOpTimeoutMs = configInput.WebOpTimeoutMs
	f.config.ModelAssetCfg.GetOrgIdUrl = configInput.GetOrgIdUrl
	f.config.ModelAssetCfg.GetTSLInstancesUrl = configInput.GetTSLInstancesUrl
	f.config.ModelAssetCfg.GetThingModelsUrl = configInput.GetThingModelsUrl
	f.config.ModelAssetCfg.GetNodesUrl = configInput.GetNodesUrl
	f.config.ModelAssetCfg.GetTreesUrl = configInput.GetTreesUrl

	f.metaDataWork.Config = f.config
}

//刷新当前断面metadata（同步方式，内部使用）
func (f *MetaDataJob) refreshMetaDataSync() error {
	var metaDataTemp *MetaData
	var err error

	metaDataTemp, err = f.metaDataWork.GetMetaData()
	if err != nil {
		return fmt.Errorf("f.metaDataWork.GetMetaData, %s", err)
	}

	f.metaDataMutex.Lock()
	defer f.metaDataMutex.Unlock()

	f.metaData = metaDataTemp

	return nil
}

//处理缓存更新的job
func (f *MetaDataJob) dealUpdateJob() {
	var swapped bool
	var err error

	var ticker *time.Ticker
	ticker = time.NewTicker(time.Second * 5)

	for {
		select {
		case <-ticker.C:
			swapped = atomic.CompareAndSwapInt32(&f.needUpdate, 1, 0)
			if swapped {
				err = f.RefreshMetaDataSync()
				if err != nil {
					log.MultiLogErrorf(f.configInput.LogArray, "[%s:%d]:refreshMetaDataSync, %s", f.srcFilename, utils.GetLineNum(), err)
				} else {
					log.MultiLogInfof(f.configInput.LogArray, "[%s:%d]:refreshMetaDataSync success", f.srcFilename, utils.GetLineNum())
				}
			}
		}
	}
}

//发送成功消息
func (f *MetaDataJob) sendSuccess() {
	if f.configInput.NotifyMsgChan != nil {
		f.configInput.NotifyMsgChan <- SUCCESS
	}
}

//发送失败消息
func (f *MetaDataJob) sendFail() {
	if f.configInput.NotifyMsgChan != nil {
		f.configInput.NotifyMsgChan <- FAIL
	}
}
