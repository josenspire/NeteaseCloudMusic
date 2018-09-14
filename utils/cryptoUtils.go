package utils

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"encoding/json"
	"strings"
)

const (
	modulus = "00e0b509f6259df8642dbc35662901477df22677ec152b5ff68ace615bb7b725152b3ab17a876aea8a5aa76d2e417629ec4ee341f56135fccf695280104e0312ecbda92557c93870114af6c9d05c4f7f0c3685b7a46bee255932575cce10b424d813cfe4875d3e82047b97ddef52741d546b8e289dc6935b3ece0462db0a22b8e7"
	nonce   = "0CoJUm6Qyw8W8jud"
	pubKey  = "010001"
	iv      = "0102030405060708"
)

type Crypto struct {
	SecretKey  string
	OriginData interface{}
}

func (ct *Crypto) CreateSecretKey() {
	const keys string = `abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789`
	ct.SecretKey = GenerateRandomString(keys, 16)
}

func (ct *Crypto) Encrypt(originData interface{}) (string, error) {

	originDataObj, err := json.Marshal(originData)

	if err != nil {
		checkError(err)
	}
	if strings.EqualFold("", ct.SecretKey) {
		ct.CreateSecretKey()
	}

	encTextStr, _ := aesEncrypt(originDataObj, nonce)
	encText, _ := aesEncrypt([]byte(encTextStr), ct.SecretKey)
	// const encSecKey = rsaEncrypt(secKey, pubKey, modulus)
	return encText, nil
}

func (ct *Crypto) Decrypt(decodeStr string) (string, error) {

	if strings.EqualFold("", ct.SecretKey) {
		ct.CreateSecretKey()
	}

	decTextStr, _ := aesDecrypt(decodeStr, ct.SecretKey)
	decText, _ := aesDecrypt(decTextStr, nonce)
	return decText, nil
}

// AES加密的具体算法为: AES-128-CBC，输出格式为 base64
// AES加密时需要指定 iv：0102030405060708
// AES加密时需要 padding
// https://github.com/darknessomi/musicbox/wiki/%E7%BD%91%E6%98%93%E4%BA%91%E9%9F%B3%E4%B9%90%E6%96%B0%E7%99%BB%E5%BD%95API%E5%88%86%E6%9E%90
func aesEncrypt(encodeBytes []byte, secretKeyStr string) (string, error) {
	secretKey := []byte(secretKeyStr)
	block, err := aes.NewCipher(secretKey)
	if err != nil {
		return "", err
	}

	blockSize := block.BlockSize()
	encodeBytes = pKCS5Padding(encodeBytes, blockSize)

	blockMode := cipher.NewCBCEncrypter(block, []byte(iv))
	crypted := make([]byte, len(encodeBytes))
	blockMode.CryptBlocks(crypted, encodeBytes)

	return base64.StdEncoding.EncodeToString(crypted), nil
}

func aesDecrypt(decodeStr string, secretKeyStr string) (string, error) {
	// decode base64
	decodeBytes, err := base64.StdEncoding.DecodeString(decodeStr)
	checkError(err)

	secretKey := []byte(secretKeyStr)
	block, _ := aes.NewCipher(secretKey)
	checkError(err)

	blockMode := cipher.NewCBCDecrypter(block, []byte(iv))
	originData := make([]byte, len(decodeBytes))

	blockMode.CryptBlocks(originData, decodeBytes)
	originData = pKCS5UnPadding(originData)

	var params interface{}
	json.Unmarshal(originData, &params)

	if params != nil {
		return params.(string), nil
	}
	return string(originData[:]), nil
}

func pKCS5Padding(cipherText []byte, blockSize int) []byte {
	padding := blockSize - len(cipherText)%blockSize // 16, 32, 48 etc..
	// 填充
	paddingText := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(cipherText, paddingText...)
}

func pKCS5UnPadding(originData []byte) []byte {
	length := len(originData)
	unPadding := int(originData[length-1])
	return originData[:(length - unPadding)]
}
