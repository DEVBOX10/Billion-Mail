package maillog_stat

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/gogf/gf/util/grand"
	"strings"
)

func Encrypt(data interface{}) string {
	dataJSON, _ := json.Marshal(data)

	key := grand.B(16)
	iv := grand.B(16)

	block, _ := aes.NewCipher(key)
	paddedData := PKCS7Padding(dataJSON, block.BlockSize())

	ciphertext := make([]byte, len(paddedData))
	mode := cipher.NewCBCEncrypter(block, iv)
	mode.CryptBlocks(ciphertext, paddedData)

	keyiv := make([]byte, 0, 32)
	for i := 0; i < 32; i++ {
		j := i / 2
		if i%2 == 0 {
			keyiv = append(keyiv, key[j])
		} else {
			keyiv = append(keyiv, iv[j])
		}
	}

	resultBytes := append(keyiv[:16], append(ciphertext, keyiv[16:]...)...)

	result := base64.URLEncoding.EncodeToString(resultBytes)
	return strings.TrimRight(result, "=")
}

func Decrypt(data string, result interface{}) (err error) {
	dataLength := len(data)
	amountToPad := 4 - (dataLength % 4)
	if amountToPad > 0 && amountToPad < 4 {
		data += strings.Repeat("=", amountToPad)
	}

	dataAes, err := base64.URLEncoding.DecodeString(data)
	if err != nil {
		return
	}

	if len(dataAes) < 32 {
		err = fmt.Errorf("invalid data length")
		return
	}

	dataAesLen := len(dataAes)
	keyiv := make([]byte, 0, 32)
	keyiv = append(keyiv, dataAes[:16]...)
	keyiv = append(keyiv, dataAes[dataAesLen-16:]...)

	var key, iv []byte
	for i := 0; i < 32; i++ {
		if i%2 == 0 {
			key = append(key, keyiv[i])
		} else {
			iv = append(iv, keyiv[i])
		}
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return
	}

	mode := cipher.NewCBCDecrypter(block, iv)
	ciphertext := dataAes[16 : dataAesLen-16]

	plaintext := make([]byte, len(ciphertext))
	mode.CryptBlocks(plaintext, ciphertext)

	unpaddedData := PKCS7UnPadding(plaintext, block.BlockSize())

	err = json.Unmarshal(unpaddedData, &result)
	if err != nil {
		return
	}

	return
}

func PKCS7Padding(data []byte, blockSize int) []byte {
	length := len(data)

	amountToPad := blockSize - (length % blockSize)

	if amountToPad == 0 {
		amountToPad = blockSize
	}

	pad := byte(amountToPad)

	for i := 0; i < amountToPad; i++ {
		data = append(data, pad)
	}

	return data
}

func PKCS7UnPadding(data []byte, blockSize int) []byte {
	length := len(data)
	pad := int(data[length-1])

	if pad < 1 || pad > blockSize {
		pad = 0
	}

	return data[:(length - pad)]
}
