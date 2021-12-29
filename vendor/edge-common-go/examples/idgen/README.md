# ID Generator
This is a id generator on Edge
## functions list
The id generator support functions list as follows:
### New()
generator a unique uuid by standard google uuid replaceAll '-' with ''

for example 'b5adbe75773311eb8e65a81e84b05e3c'
## Use Example:
### Step1: add in glide.yaml or go.mod:
```yaml
glide.yaml:
- package: git.envisioncn.com/edge/edge-common-go
  version: 1ad8876408767d6545d7d790866707605c6cbbd3

go.mod:
require (
	edge-common-go v0.0.0-20200922033041-1ad8876408767
)
replace edge-common-go v0.0.0-20200922033041-1ad8876408767 => git.envisioncn.com/edge/edge-common-go v0.0.0-20200922033041-1ad8876408767
```


### Step2: add import:
```go
import "git.envisioncn.com/edge/edge-common-go/pkg/idgen"
```


### Step3: add struct and New():
```go
type MyIDGen struct {
	idgen.IDGenerator
}
var myIDGen MyIDGen
myID, err := myIDGen.New()
```



