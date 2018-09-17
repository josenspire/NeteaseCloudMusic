package utils

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"math/big"
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

func (ct *Crypto) RSAEncrypt(originData string) string {
	encSecKey := rsaEncrypt(originData, pubKey, modulus)
	return encSecKey
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

func rsaEncrypt(plainText string, pubKey string, modulus string) string {
	// 倒序 key
	rKey := ""
	for i := len(plainText) - 1; i >= 0; i-- {
		rKey += plainText[i : i+1]
	}
	// 将 key 转 ascii 编码 然后转成 16 进制字符串
	hexRKey := ""
	for _, char := range []rune(rKey) {
		hexRKey += fmt.Sprintf("%x", int(char))
	}
	// 将 16进制 的 三个参数 转为10进制的 bigint
	bigRKey, _ := big.NewInt(0).SetString(hexRKey, 16)
	bigPubKey, _ := big.NewInt(0).SetString(pubKey, 16)
	bigModulus, _ := big.NewInt(0).SetString(modulus, 16)
	// 执行幂乘取模运算得到最终的bigint结果
	bigRs := bigRKey.Exp(bigRKey, bigPubKey, bigModulus)
	// 将结果转为 16进制字符串
	hexRs := fmt.Sprintf("%x", bigRs)
	// 可能存在不满256位的情况，要在前面补0补满256位
	return addRSAPadding(hexRs, modulus)
}

// 补0步骤
func addRSAPadding(encText string, modulus string) string {
	ml := len(modulus)
	for i := 0; ml > 0 && modulus[i:i+1] == "0"; i++ {
		ml--
	}
	num := ml - len(encText)
	prefix := ""
	for i := 0; i < num; i++ {
		prefix += "0"
	}
	return prefix + encText
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
