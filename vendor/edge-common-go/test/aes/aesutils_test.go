package aes

import (
	"edge-common-go/pkg/aes"
	"fmt"
	"testing"
)

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
