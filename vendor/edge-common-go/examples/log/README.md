# edge log utils
This is an Edge log utils.
Please use this kit follow as:

## Glide:
### Step1: add in glide.yaml:
```yaml
- package: git.envisioncn.com/edge/edge-common-go
  version: tag_edge-common-go_20200916_001
```


### Step2: add import:
```go
import "git.envisioncn.com/edge/edge-common-go/pkg/log"
```


### Step3: init logAGlobal:
```go
log.LogAGlobal.Init(GetBaseDir())

func GetBaseDir() string {
	ex, err := os.Executable()
	if err != nil {
		panic(err)
	}
	sp := string(filepath.Separator)
	baseDir := strings.TrimSuffix(filepath.Dir(ex), sp+"bin")
	return baseDir
}
```


### Step4: init log like:
```go
EdgeLog, _ := log.GetLogAa("log/path/edge.log")
```


### Step5: print log like:
```go
EdgeLog.Infof("edge log %s", "info")
EdgeLog.Info("edge log")
EdgeLog.Debugf("edge log %s", "debug")
EdgeLog.Debug("edge log")
EdgeLog.Errorf("edge log %s", "error")
EdgeLog.Error("edge log")
```


