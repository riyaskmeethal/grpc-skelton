package crypto

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"encoding/hex"
	"fmt"
)

func unpadPayload(payload []byte, blockSize int) ([]byte, error) {
	padding := int(payload[len(payload)-1])
	if padding > blockSize {
		return nil, fmt.Errorf("invalid padding")
	}

	return payload[:len(payload)-padding], nil
}

// func DecryptPayloadb64(encPayload string, key []byte) ([]byte, error) {
// 	encryptedPayload, err := base64.StdEncoding.DecodeString(encPayload)
// 	if err != nil {
// 		return nil, err
// 	}

//		return DecryptBytes(encryptedPayload, key)
//	}
func DecryptB64Bytes(encPayload string, key []byte) ([]byte, error) {
	encryptedPayload, err := base64.StdEncoding.DecodeString(encPayload)
	if err != nil {
		return nil, err
	}
	return decryptECB(key, encryptedPayload)
}
func DecryptHexBytes(encPayload string, key []byte) ([]byte, error) {
	encryptedPayload, err := hex.DecodeString(encPayload)
	if err != nil {
		return nil, err
	}

	return decryptECB(key, encryptedPayload)
}

func DecryptCBC(encryptedPayload []byte, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	if len(encryptedPayload) < 1 {
		return nil, fmt.Errorf("payload empty")
	}

	iv := encryptedPayload[:aes.BlockSize]
	ciphertext := encryptedPayload[aes.BlockSize:]

	mode := cipher.NewCBCDecrypter(block, iv)

	paddedPayload := make([]byte, len(ciphertext))
	mode.CryptBlocks(paddedPayload, ciphertext)

	payload, err := unpadPayload(paddedPayload, aes.BlockSize)
	if err != nil {
		return nil, err
	}

	return payload, nil
}
