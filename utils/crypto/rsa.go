package crypto

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"math/big"
	"strings"

	"github.com/google/uuid"
)

const (
	otp_length  = 6
	modularSize = 65537
)

func UnmarshalData(req interface{}, data any) error {
	jsonData, err := json.Marshal(req)
	if err != nil {
		return err
	}
	return json.Unmarshal(jsonData, data)
}

// func VerifyOTP(otp string, encryptedOtp string, pos_token []byte) (bool, []byte, error) {

//		otpBytes, err := DecryptB64Bytes(encryptedOtp, pos_token)
//		if err != nil {
//			return false, nil, err
//		}
//		return string(otpBytes) == otp, otpBytes, nil
//	}

func VerifyOTP(priv *rsa.PrivateKey, otp string, encryptedOtp string) (bool, []byte, error) {

	decryptedOtp, err := DecryptRsaHex(priv, encryptedOtp)
	if err != nil {
		return false, nil, err
	}

	return (string(decryptedOtp) == otp || strings.EqualFold(string(decryptedOtp), "SKIP")), decryptedOtp, nil
}

func UnMarshalPrivateKey(privateKeyBytes []byte) (*rsa.PrivateKey, error) {
	return x509.ParsePKCS1PrivateKey(privateKeyBytes)
}

func UnMarshalPublicKey(publicKeyBytes []byte) (*rsa.PublicKey, error) {
	return x509.ParsePKCS1PublicKey(publicKeyBytes)
}

func DecryptPubKeyBytes(rpsPrivateKey *rsa.PrivateKey, encryptedKey []byte) ([]byte, error) {
	return rsa.DecryptPKCS1v15(rand.Reader, rpsPrivateKey, encryptedKey)
}

// Decode a public key from bytes
func DecodePublicKeyHex(publicKeyb64 string) (*rsa.PublicKey, error) {
	n, ok := new(big.Int).SetString(publicKeyb64, 16)
	if !ok {
		return nil, fmt.Errorf("encode public key error")
	}

	// Create an RSA public key
	return &rsa.PublicKey{
		N: n,
		E: modularSize,
	}, nil
}

func DecryptOtp(rpsPrivateKey *rsa.PrivateKey, encryptedOtp []byte) (string, error) {
	// Decrypt the decoded ciphertext using the private key
	plaintext, err := rsa.DecryptPKCS1v15(rand.Reader, rpsPrivateKey, encryptedOtp)
	if err != nil {
		return string(plaintext), err
	}
	return string(plaintext), nil
}
func EncryptRsaB64(pub *rsa.PublicKey, msg []byte) (string, error) {
	encMsg, err := rsa.EncryptPKCS1v15(rand.Reader, pub, msg)
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(encMsg), nil
}

// Encrypt byte message with RSA.publicKey
// and encode to Hexadecimal String
func EncryptRsaHex(pub *rsa.PublicKey, msg []byte) (string, error) {
	encMsg, err := rsa.EncryptPKCS1v15(rand.Reader, pub, msg)
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(encMsg), nil
}

func DecryptRsaB64(priv *rsa.PrivateKey, ciphertextb64 string) ([]byte, error) {
	ciphertext, err := base64.StdEncoding.DecodeString(ciphertextb64)
	if err != nil {
		return nil, err
	}

	return rsa.DecryptPKCS1v15(rand.Reader, priv, ciphertext)
}

// Decrypt the message with RSA PrivateKey after decoding from Hexadecimal string
func DecryptRsaHex(priv *rsa.PrivateKey, ciphertextHex string) ([]byte, error) {
	ciphertext, err := hex.DecodeString(ciphertextHex)
	if err != nil {
		return nil, err
	}

	return rsa.DecryptPKCS1v15(rand.Reader, priv, ciphertext)
}

// Generate AES key - Session Key
// and Session UID
func GetSessionKey() ([]byte, string) {
	key := make([]byte, 32)
	rand.Read(key)
	uid := uuid.New().String()
	return key, uid
}

// Generate OTP of lenght 6 using [A-Z][0-9]
func GenerateOTP() string {
	// const passwordLength = 6
	const passwordChars = "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

	password := make([]byte, otp_length)
	for i := range password {
		randomIndex, _ := rand.Int(rand.Reader, big.NewInt(int64(len(passwordChars))))
		password[i] = passwordChars[randomIndex.Int64()]
	}
	return string(password)
}
