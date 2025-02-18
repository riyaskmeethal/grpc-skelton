package crypto

import (
	"bytes"
	"crypto/aes"
	"fmt"
)

func encryptECB(key, plaintext []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	// Ensure the plaintext is a multiple of the block size
	plaintext = PKCS7Padding(plaintext, block.BlockSize())

	ciphertext := make([]byte, len(plaintext))

	// Encrypt each block of plaintext separately
	for i := 0; i < len(plaintext); i += block.BlockSize() {
		block.Encrypt(ciphertext[i:i+block.BlockSize()], plaintext[i:i+block.BlockSize()])
	}

	return ciphertext, nil
}

func decryptECB(key, ciphertext []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	// Ensure the ciphertext is a multiple of the block size
	if len(ciphertext)%block.BlockSize() != 0 {
		return nil, fmt.Errorf("ciphertext length is not a multiple of block size")
	}

	plaintext := make([]byte, len(ciphertext))

	// Decrypt each block of ciphertext separately
	for i := 0; i < len(ciphertext); i += block.BlockSize() {
		block.Decrypt(plaintext[i:i+block.BlockSize()], ciphertext[i:i+block.BlockSize()])
	}

	// Remove padding
	plaintext = PKCS7Unpadding(plaintext)

	return plaintext, nil
}

// PKCS7Padding pads the data using the PKCS7 padding scheme
func PKCS7Padding(data []byte, blockSize int) []byte {
	padding := blockSize - (len(data) % blockSize)
	padText := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(data, padText...)
}

// PKCS7Unpadding removes the PKCS7 padding from the data
func PKCS7Unpadding(data []byte) []byte {

	length := len(data)
	unpadding := int(data[length-1])

	return data[:(length - unpadding)]
}
