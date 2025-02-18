package crypto

// import (
// 	"crypto/rand"
// 	"crypto/rsa"
// 	"crypto/x509"
// 	"encoding/base64"
// 	"encoding/hex"
// 	"encoding/json"
// 	"fmt"
// 	"math/big"
// 	"testing"

// )

// type TransactionInfo struct {
// 	OmaErrorCode              string `json:"omaErrorCode"`
// 	OmaErrorMessage           string `json:"omaErrorMessage"`
// 	OmaIsTransactionSuccess   string `json:"omaIsTransactionSuccess"`
// 	OmaTransactionStatusMsg   string `json:"omaTransactionStatusMsg"`
// 	OmaTxnClientRefNumber     string `json:"omaTxnClientRefNumber"`
// 	OmaTxnMwRequestId         string `json:"omaTxnMwRequestId"`
// 	OmaTerminalId             string `json:"omaTerminalId"`
// 	OmaTidLocation            string `json:"omaTidLocation"`
// 	OmaTxnType                string `json:"omaTxnType"`
// 	OmaTxnTypeDesc            string `json:"omaTxnTypeDesc"`
// 	OmaAmount                 string `json:"omaAmount"`
// 	OmaTotalAmount            string `json:"omaTotalAmount"`
// 	OmaAuthCode               string `json:"omaAuthCode"`
// 	OmaStanNumber             string `json:"omaStanNumber"`
// 	OmaInvoiceNumber          string `json:"omaInvoiceNumber"`
// 	OmaBatchNumber            string `json:"omaBatchNumber"`
// 	OmaRrn                    string `json:"omaRrn"`
// 	OmaMaskedPan              string `json:"omaMaskedPan"`
// 	OmaCardType               string `json:"omaCardType"`
// 	OmaMid                    string `json:"omaMid"`
// 	OmaMerchantName           string `json:"omaMerchantName"`
// 	OmaMerchantAddress        string `json:"omaMerchantAddress"`
// 	OmaMerchantCity           string `json:"omaMerchantCity"`
// 	OmaMerchantZIpCode        string `json:"omaMerchantZIpCode"`
// 	OmaPaymentAuthMode        string `json:"omaPaymentAuthMode"`
// 	OmaPaymentEntryMode       string `json:"omaPaymentEntryMode"`
// 	OmaDeviceTxnStartTime     string `json:"omaDeviceTxnStartTime"`
// 	OmaDeviceTxnEndTime       string `json:"omaDeviceTxnEndTime"`
// 	OmaHostTxnDateEndTime     string `json:"omaHostTxnDateEndTime"`
// 	OmaCurrencyCode           string `json:"omaCurrencyCode"`
// 	OmaDeviceSerial           string `json:"omaDeviceSerial"`
// 	OmaAddtionalInfo          string `json:"omaAddtionalInfo"`
// 	OmaDccInfo                string `json:"omaDccInfo"`
// 	OmaEppInfo                string `json:"omaEppInfo"`
// 	OmaDiscountInfo           string `json:"omaDiscountInfo"`
// 	OmaLoyaltyInfo            string `json:"omaLoyaltyInfo"`
// 	OmaQrInfo                 string `json:"omaQrInfo"`
// 	OmaUrlReceipt             string `json:"omaUrlReceipt"`
// 	OmaCustomerEmail          string `json:"omaCustomerEmail"`
// 	OmaCustomerMobileNumber   string `json:"omaCustomerMobileNumber"`
// 	OmaAppVersion             string `json:"omaAppVersion"`
// 	OmaMRefLabel              string `json:"omaMRefLabel"`
// 	OmaMRefValue              string `json:"omaMRefValue"`
// 	OmaIsDuplicateCopy        string `json:"omaIsDuplicateCopy"`
// 	OmaCustomerSloganMessage1 string `json:"omaCustomerSloganMessage1"`
// 	OmaMerchantSloganMessage1 string `json:"omaMerchantSloganMessage1"`
// }

// func TestEncryption(t *testing.T) {
// 	req := &pos.PosSaleRequest{
// 		OmaTerminalId:         "80040197",
// 		OmaDeviceSerialNumber: "9220207104",
// 	}
// 	aes_key, err := base64.StdEncoding.DecodeString("7FCWVEVVZx98AFaFeVsElvPkb8LdR51pfZBYQjAs2Ts=")
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	fmt.Println(hex.EncodeToString(aes_key))

// 	encdata, _ := EncryptPayloadHex(req, aes_key)

// 	fmt.Println("enc data ", encdata)
// 	// encdata := "88A118BEE4E40D543D659AE219A34F2645C9AC9ECE520B6C8609945ADA04606EA161B82236853AE56C44002430A9C91B543DE2ECD343FD5C03BC837E921108EEFEC49B326D312E488C7341ABC5083E7045F2D8C35C7A35EDE456C4367AA147B4"

// 	decData, err := DecryptHexBytes(encdata, aes_key)
// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	resData := pos.PosSaleRequest{}
// 	fmt.Println(decData)
// 	json.Unmarshal(decData, &resData)
// 	fmt.Println(&resData)
// }
// func TestKeyEnc(t *testing.T) {
// 	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
// 	if err != nil {
// 		t.Fatal(err)
// 		return
// 	}
// 	// posPriv, err := rsa.GenerateKey(rand.Reader, 2048)
// 	// if err != nil {
// 	// 	t.Fatal(err)
// 	// 	return
// 	// }
// 	//	posPubKeyBytes := x509.MarshalPKCS1PublicKey(&posPriv.PublicKey)

// 	aes_key, _ := GetSessionKey()
// 	publicKeyBytes := x509.MarshalPKCS1PublicKey(&privateKey.PublicKey)
// 	encPubKey, err := encryptECB(publicKeyBytes, aes_key)
// 	if err != nil {
// 		t.Fatal(err)
// 		return
// 	}
// 	// token, err := rsa.EncryptPKCS1v15(rand.Reader, &posPriv.PublicKey, aes_key)
// 	// if err != nil {
// 	// 	t.Fatal(err)
// 	// 	return
// 	// }

// 	// rps_token, err := rsa.DecryptPKCS1v15(rand.Reader, posPriv, token)
// 	// if err != nil {
// 	// 	t.Fatal(err)
// 	// 	return
// 	// }
// 	// if bytes.Equal(rps_token, token) {
// 	// 	t.Fatalf("token error, ")
// 	// 	return
// 	// }

// 	rpsPublicKeybytes, err := decryptECB(aes_key, encPubKey)
// 	if err != nil {
// 		t.Fatal(err)
// 		return
// 	}

// 	rpsPublicKey, err := UnMarshalPublicKey(rpsPublicKeybytes)
// 	if err != nil {
// 		t.Fatal(err)
// 		return
// 	}
// 	if !rpsPublicKey.Equal(privateKey.PublicKey) {
// 		t.Fatal("err")
// 	}

// }

// func TestB64(t *testing.T) {
// 	data := "wrK4h3wsYQpDRsWfLIBGq78JOkTDDrvKY3iVyOCZUHDWj/0UBdjjawrTUJ+lEEmYxQAIUVVvKJGj\nXOEYWrHov7Jd4xvKGkPnbpBGCW8yqD0QodpGtU5MXeorltiwA5y5JCqxoRub1EjrXpU9lVLizDGg\nhqqCPrVRZCM+rjwUdWHoktE1bRjzYku6g5eZNXw+v0AFRpnyjp+mkCFg+s+l+KzKHaRggtVXci+V\nhy3kqToDl9DWPg6qfK9KAs7ZxJHXotNHZQqKWvJi9pFwlIyPgbexLNkjZkf4neE8Zy/umfTADPD6\naFx+bg8SXZgWjCHndjLw5ZJax1U/OJrqJNiJnw\u003d\u003d"
// 	b, err := base64.StdEncoding.DecodeString(data)
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	fmt.Println(len(b))
// 	// Convert the bytes to big integers
// 	n := new(big.Int).SetBytes(b)
// 	// e := new(big.Int).SetBytes(eBytes)

// 	moSize := []int{3, 17, 257, 65537}
// 	// Create an RSA public key
// 	for _, v := range moSize {
// 		publicKey := &rsa.PublicKey{
// 			N: n,
// 			E: v,
// 		}
// 		fmt.Println(publicKey.E)
// 		pb := x509.MarshalPKCS1PublicKey(publicKey)
// 		fmt.Println("for", v, "size: ", len(pb))
// 		_, err = x509.ParsePKCS1PublicKey(pb)
// 		if err != nil {
// 			t.Fatalf(err.Error())
// 		}
// 	}
// }

// func TestEncode(t *testing.T) {
// 	publickeyBytes := []int{-59, -119, -53, 21, 64, 1, 92, 104, 15, -84, 50, -101, -92, 57, 30, 8, 11, -72, -58, -20, -45, 71, -110, 113, -120, 45, 36, 102, -26, -65, -88, 81, -100, 58, 122, -94, -10, 34, -38, -8, -61, 76, -16, -57, 119, -38, -47, -4, -6, -115, -98, 64, 13, -25, -37, -25, 46, -117, 110, 109, -75, 39, -105, 66, -87, -33, 112, -101, 97, 35, 14, -14, -22, 6, 84, 0, 39, 16, -104, 24, -4, 90, 50, 117, 61, 71, -93, -37, -24, 25, -23, 57, 116, -66, 10, -97, -15, -40, 81, -8, 27, -16, -101, -94, 36, 53, -80, 78, 70, -4, -70, -47, -58, -15, -94, -125, -95, 123, 44, -30, 61, 82, -74, 117, -118, 1, 4, 109, 54, -39, 36, 58, -33, 55, -31, -78, -46, 6, -118, -1, 17, 35, 125, 123, -2, -106, -15, -67, 75, -20, 98, 95, -63, 85, 114, -126, -83, 114, -21, -28, 28, -82, -61, -125, 77, -1, -104, 108, 101, 10, 41, -24, 2, 94, -94, 109, 103, -108, 33, 57, -72, 63, -62, -47, 13, -18, -74, 119, -116, 29, -7, -110, 127, 20, -42, -111, -104, 43, 125, 17, -83, 52, -123, -60, 75, 40, 81, -22, 118, 7, 37, -17, 74, 39, 125, 59, -75, -36, 52, -87, -26, 117, -87, 74, 120, -119, -52, -51, 3, 37, 72, -89, -27, 52, -55, 89, 28, 92, -51, -9, -104, -5, 113, -128, 126, -2, 11, -96, 8, 100, -128, 93, -117, 76, 39, 39}
// 	pubH := []byte{}
// 	// pub := []byte{}
// 	privateKey, _ := rsa.GenerateKey(rand.Reader, 2048)
// 	fmt.Println(byte(privateKey.PublicKey.E))
// 	for _, v := range publickeyBytes {
// 		pubH = append(pubH, byte(v))
// 	}
// 	fmt.Println(len(privateKey.PublicKey.N.Bytes()))
// 	fmt.Println(privateKey.PublicKey.N.Bytes())
// 	fmt.Println(len(pubH))
// 	// _, err := hex.Decode(pub, pubH)
// 	// if err != nil {
// 	// 	t.Fatal(err)
// 	// }
// 	// key, err := DecodePublicKeyFromBytes(pubH)
// 	// if err != nil {
// 	// 	t.Fatal(err)
// 	// }
// 	// pb := x509.MarshalPKCS1PublicKey(key)
// 	// _, err = x509.ParsePKCS1PublicKey(pb)
// 	// if err != nil {
// 	// 	t.Fatalf(err.Error())
// 	// }
// }
