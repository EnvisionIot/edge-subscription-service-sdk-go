# edge-common-go

edge common golang library

Table of Contents

* [Installation](#install)
    * [Glide](#glide)
    * [go module](#goModule)
* [Feature List](#feature)
    * [LionUtils](examples/lion/README.md) ---- Provide Lion reading tools on Edge
    * [HttpUtils](examples/httpclient/README.md) ----Provide http request tools on Edge
    * [LogUtils](examples/log/README.md) ----Provides unified log tools on edge
    * NSQClient ----TODO
    * [RedisClient](examples/redisclient/README.md) ----Provide redis client tools on edge
    * [AESUtils](examples/aes/README.md) ----Provide AES encryption and decryption tools on Edge
    * [IdGen](examples/idgen/README.md) ----Provide Random ID generator on Edge
    * ......
* [Others](#others)

<a name="#install"></a>

## Install

<a name="#glide"></a>

### Glide

#### step1 add this repo into your glide.yaml
```yaml
- package: git.envisioncn.com/edge/edge-common-go
  version: tag_edge-common-go_20200609_001
```
#### step2 add mirrors
Before using it on the local machine, you need to map HTTPS to http.

Add the following content to mirrors.yaml：
```yaml
- original: https://git.envisioncn.com/edge/edge-common-go
  repo: http://git.envisioncn.com/edge/edge-common-go
  vcs: git
```

<a name="#goModule"></a>
### go module
* Golang's version 1.12 or higher
* Due to historical reasons, our business module has adopted the following methods to import the edge-common-go. 
  Although there is a better way, please follow the steps below at this stage
### Step1: Run command: `go mod edit -require git.envisioncn.com/edge/edge-common-go.git@commitId` in Terminal:
```yaml
for example run: go mod edit -require git.envisioncn.com/edge/edge-common-go.git@5462187e8905f30398f71844e87bd18940b4f3e2
you will get edge-common-go with version like this:

go.mod:
require (
    git.envisioncn.com/edge/edge-common-go.git v0.0.0-20210310053653-5462187e8905
)
```
### Step2: Edit go.mod with above `version-date-commit` and Add `replace` configuration:
```yaml
go.mod:
require (
	edge-common-go v0.0.0-20210310053653-5462187e8905
)
replace edge-common-go v0.0.0-20210310053653-5462187e8905 => git.envisioncn.com/edge/edge-common-go v0.0.0-20210310053653-5462187e8905
```

### Step3: Import edge-common-go in your bussiness file and Refer to the specific usage README in Feature List:
```yaml
for example log and lion:
import (
  "edge-common-go/pkg/log"
  "edge-common-go/pkg/lion"
)
```

<a name="#feature"></a>
## Feature List
* [LionUtils](examples/lion/README.md) ---- Provide Lion reading tools on Edge
* [HttpUtils](examples/httpclient/README.md) ----Provide http request tools on Edge
* [LogUtils](examples/log/README.md) ----Provides unified log tools on edge
* NSQClient ----TODO
* [RedisClient](examples/redisclient/README.md) ----Provide redis client tools on edge
* [AESUtils](examples/aes/README.md) ----Provide AES encryption and decryption tools on Edge
* [IdGen](examples/idgen/README.md) ----Provide Random ID generator on Edge

    
## Others：
Our first step is to build a Edge general go tool library, so please refer to the Convention for naming new tool packages.https://github.com/golang-standards/project-layout.
When you submit your tool to the project, please attach examples and test.
