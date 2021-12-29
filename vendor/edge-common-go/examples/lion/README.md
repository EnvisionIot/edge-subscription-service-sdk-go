# lion config parse
This is an Edge lion utils.
Please use this kit follow as:

## Glide:
### Step1: add in glide.yaml:
```yaml
- package: git.envisioncn.com/edge/edge-common-go
  version: tag_edge-common-go_20210107_001
```


### Step2: add import:
```go
import "git.envisioncn.com/edge/edge-common-go/pkg/lion"
```


### Step3: add struct like:
```go
type MyLionConfig struct {
	lion.Default_LionConfig
}
var myLionConfig MyLionConfig
myLionConfig.LoadLionConfig(&myLionConfig)
```


### Step4: add target var:
```go
var Example bool
```


### Step5: overwrite func `InitLionConfig`
```go
func (l *MyLionConfig) InitLionConfig()  {
	Example = lion.GetBoolValue("edge.edge-archive.example.key", false)
}
```

### Step6(optional):get app lion info 
Print out or Provide external query interface

We agree on a unified print out log path: 
`/home/envuser/energy-os/edge-xxx/logs/lionInMemory.log`
```go
myLog.Infof("appLion: %v", lion.GetAppLion())
//or
for key, value := range lion.GetAppLion() {
	myLog.Infof("%s=%v", key, value)
}
//or get app lion info with an ordered keySlice
lionMap, lionKeySlice := lion.GetAppLionWithOrderedSlice()
for _, key := range lionKeySlice {
	myLog.Infof(fmt.Sprintf("%s=%v", key, lionMap[key]))
}
```

### decrypt lion key 
Some key in lion need Decrypt
```
//lion.properties
config-manger.encrpyt.on_off=true
config-manger.encrpyt.file.first=./k/k_1
config-manger.encrpyt.file.second=./k/k_2
config-manger.encrpyt.file.third=./k/k_3

edge.redis.password=1dAGxTAvVAb64Bx5J+QJ8xaCdq740p1VwTrsckEsqMU=
edge.mongodb.password=2Z8kiu5Cp8RT8oOJMOOOu8iloVB9ER4f9D+XD/rXQ/k=
edge.mysql.password=EsWSJMH/6Y+K6voCY/DifXEstD2Vs6N5O6JEmJkysoo=
edge.pushgateway.pass_word=c2Lb5r02JE5Acns5eNUtc2HDTjVu4oalDKikXsQAC28=

edge.need_en.xxx.p=EsWSJMH/6Y+K6voCY/DifXEstD2Vs6N5O6JEmJkysoo=

edge.aaa.bbb=abcd
```

```go
//configure in lion/lionaes.go
var (
	LionKeysNeedEn = map[string]string{
		"edge.redis.password":        "1",
		"edge.mongodb.password":      "1",
		"edge.mysql.password":        "1",
		"edge.pushgateway.pass_word": "1",
	}

    //if key contains these strings, will decrypt
	LionKeysMatchNeedEn = map[string]string{
		".need_en": "1",
		"_need_en": "1",
	}

)

//after LoadLionConfig, these keys will be automatically decrypted into plaintext
var defaultLionConfig MyLionConfig
_ = defaultLionConfig.LoadLionConfig(&defaultLionConfig)

//edge.redis.password is Envisi0nEdge@1
//edge.mongodb.password is Envisi0nEdge@1
//edge.mysql.password is Envisi0nEdge@1
//edge.pushgateway.pass_word is A,12345678
//edge.need_en.xxx.p is Envisi0nEdge@1
//edge.aaa.bbb is abcd
```
