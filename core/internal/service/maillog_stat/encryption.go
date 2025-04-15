package maillog_stat

import (
	"bytes"
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

	keyiv := append(dataAes[:16], dataAes[len(dataAes)-16:]...)

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
	ciphertext := dataAes[16 : len(dataAes)-16]
	plaintext := make([]byte, len(ciphertext))
	mode.CryptBlocks(plaintext, ciphertext)

	unpaddedData := PKCS7Unpadding(plaintext)

	err = json.Unmarshal(unpaddedData, &result)
	if err != nil {
		return
	}

	return
}

func PKCS7Padding(data []byte, blockSize int) []byte {
	padding := blockSize - len(data)%blockSize
	padText := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(data, padText...)
}

func PKCS7Unpadding(data []byte) []byte {
	length := len(data)
	if length == 0 {
		return data
	}
	unpadding := int(data[length-1])
	return data[:(length - unpadding)]
}
