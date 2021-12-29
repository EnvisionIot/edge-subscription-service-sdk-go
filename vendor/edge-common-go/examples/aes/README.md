# AES-256-CBC
Encrypt and decrypt

## Use Example:
### Step1: add in glide.yaml or go.mod:
```yaml
mirrors.yaml(for glide only):
repos:
- original: https://edge-common-go
  repo: https://git.envisioncn.com/edge/edge-common-go
  vcs: git

glide.yaml:
- package: edge-common-go
  version: 182df3a7d400ce8182b5fb3cfb86cb3f57c49ee5

go.mod:
require (
	edge-common-go v0.0.0-20210309112549-182df3a7d400
)
replace edge-common-go v0.0.0-20210309112549-182df3a7d400 => git.envisioncn.com/edge/edge-common-go v0.0.0-20210309112549-182df3a7d400
```


### Step2: add import:
```go
import "edge-common-go/pkg/aes"
```


### Step3: encrypt and decrypt:
```go
func TestEncryptValueWithP(t *testing.T) {
	var srcString = "aaaaaabbbbbbbcccccddde12643978rrryyyfff"
	var result1 []byte
	var err error
	result1, err = aes.AES256Encrypt([]byte(srcString), aes.GetDefaultP())
	if err != nil {
		t.Fatalf("aes.EncryptValueWithP error, %s", err)
	}

	var result2 []byte
	result2, err = aes.AES256Decrypt(result1, aes.GetDefaultP())
	if err != nil {
		t.Fatalf("aes.DecryptValueWithP error, %s", err)
	}

	fmt.Printf("src string=%s\n", srcString)
	fmt.Printf("encrypt string=%s\n", string(result1))
	fmt.Printf("decrypt string=%s\n", string(result2))

	if srcString != string(result2) {
		t.Fatalf("srcString(%s) != decrypt string(%s)\n", srcString, string(result2))
	}
}
```



