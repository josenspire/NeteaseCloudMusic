package utils

type Crypto struct {
	SecretKey  string
	OriginData interface{}
	EncodeData string
}

func (ct *Crypto) CreateSecretKey(size int) {
	const keys string = `abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789`
	ct.SecretKey = GenerateRandomString(keys, size)
}

func (ct *Crypto) Encode(originData string) {

}

func (ct *Crypto) Decode(encodeData string) {

}
