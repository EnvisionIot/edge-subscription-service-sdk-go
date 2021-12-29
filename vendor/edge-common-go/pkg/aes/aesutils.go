package aes

/**
 * @Author: meng.chai
 * @Date: 2021/3/03 10:55
 */

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"errors"
	"fmt"
	"io"
)

// en_K must be 32 chars long for aes256
const en_K = "1ri86clp@KRGBWL60xgdueqmgd*sycal"

// Encrypt encrypts plain text string into cipher text string, 会将iv放在结果的前16字节，iv是随机数，所以每次加密结果都不一样
func AES256Encrypt(unencrypted []byte, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	plainText, err := PKCS7Pad(unencrypted, block.BlockSize())
	if err != nil {
		return nil, fmt.Errorf(`PKCS7Pad, %s`, err)
	}
	if len(plainText)%block.BlockSize() != 0 {
		err := fmt.Errorf(`plainText: "%s" has the wrong block size`, plainText)
		return nil, err
	}

	cipherText := make([]byte, block.BlockSize()+len(plainText))
	iv := cipherText[:block.BlockSize()]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return nil, err
	}

	mode := cipher.NewCBCEncrypter(block, iv)
	mode.CryptBlocks(cipherText[block.BlockSize():], plainText)

	var base64String = base64.StdEncoding.EncodeToString(cipherText)
	return []byte(base64String), nil
}

// Decrypt decrypts cipher text string into plain text string，会先从前16字节取出iv
func AES256Decrypt(encrypted []byte, key []byte) ([]byte, error) {
	cipherText, _ := base64.StdEncoding.DecodeString(string(encrypted))

	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	if len(cipherText) < block.BlockSize() {
		return nil, errors.New("cipherText too short")
	}
	iv := cipherText[:block.BlockSize()]
	cipherText = cipherText[block.BlockSize():]
	if len(cipherText)%block.BlockSize() != 0 {
		return nil, errors.New("cipherText is not a multiple of the block size")
	}

	mode := cipher.NewCBCDecrypter(block, iv)
	mode.CryptBlocks(cipherText, cipherText)

	cipherText, _ = PKCS7UnPad(cipherText, block.BlockSize())
	return cipherText, nil
}

func PKCS7Pad(buf []byte, size int) ([]byte, error) {
	bufLen := len(buf)
	padLen := size - bufLen%size
	padded := make([]byte, bufLen+padLen)
	copy(padded, buf)
	for i := 0; i < padLen; i++ {
		padded[bufLen+i] = byte(padLen)
	}
	return padded, nil
}

func PKCS7UnPad(padded []byte, size int) ([]byte, error) {
	if len(padded)%size != 0 {
		return nil, errors.New("pkcs7: Padded value wasn't in correct size.")
	}

	bufLen := len(padded) - int(padded[len(padded)-1])
	buf := make([]byte, bufLen)
	copy(buf, padded[:bufLen])
	return buf, nil
}

//获取默认内置密码（预定义的）
func GetDefaultP() []byte {
	return []byte(en_K)
}
