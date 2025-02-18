package crypto

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// PKCS7 padding
func padPayload(payload []byte, blockSize int) []byte {
	padding := blockSize - len(payload)%blockSize
	padText := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(payload, padText...)
}

func EncryptPayloadb64(payload interface{}, key []byte) (string, error) {
	encryptedPayload, err := EncryptPayload(payload, key)
	if err != nil {
		return "", err
	}

	return base64.StdEncoding.EncodeToString(encryptedPayload), nil
}
func EncryptPayloadHex(payload interface{}, key []byte) (string, error) {
	encryptedPayload, err := EncryptPayload(payload, key)
	if err != nil {
		return "", err
	}

	return hex.EncodeToString(encryptedPayload), nil
}

func EncryptPayload(payload interface{}, key []byte) ([]byte, error) {
	data, err := json.Marshal(payload)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to marshal request body")
	}

	return encryptECB(key, data)
}
func EncryptBytesb64(data, key []byte) (string, error) {
	encbytes, err := encryptECB(key, data)
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(encbytes), nil
}
func EncryptBytesHex(data, key []byte) (string, error) {
	encbytes, err := encryptECB(key, data)
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(encbytes), nil
}

func EncryptCBC(data, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	iv := make([]byte, aes.BlockSize)
	if _, err := rand.Read(iv); err != nil {
		return nil, err
	}

	mode := cipher.NewCBCEncrypter(block, iv)

	paddedPayload := padPayload(data, aes.BlockSize)

	ciphertext := make([]byte, len(paddedPayload))
	mode.CryptBlocks(ciphertext, paddedPayload)

	return append(iv, ciphertext...), nil
}
