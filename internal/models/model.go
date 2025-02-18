package models

type Device struct {
	OmaTerminalID           string `json:"omaTerminalID,omitempty"`
	OmaClientID             string `json:"omaClientID,omitempty"`
	OmaSerialNumber         string `json:"omaSerialNumber,omitempty"`
	OmaMerchantId           string `json:"omaMerchantId,omitempty"`
	RpsPrivateKey           []byte `json:"rpsPrivateKey,omitempty"`
	PosPublicKey            []byte `json:"posPublicKey,omitempty"`
	ClientPublicKey         []byte `json:"clientPublicKey,omitempty"`
	PosAesKey               []byte `json:"posAesKey,omitempty"`
	ClientAesKey            []byte `json:"clientAesKey,omitempty"`
	PosAesUid               string `json:"posAesUid,omitempty"`
	ClientAesUid            string `json:"clientAesUid,omitempty"`
	OmaOTP                  string `json:"omaOTP,omitempty"`
	OtpExpiresAt            string `json:"otpExpiresAt,omitempty"`
	IsRereg                 bool   `json:"IsRereg,omitempty"`
	IsTerminalVerified      bool   `json:"isTerminalVerified,omitempty"`
	IsClientVerified        bool   `json:"isClientVerified,omitempty"`
	OmaCustomerEmail        string `json:"omaCustomerEmail,omitempty"`
	OmaCustomerMobileNumber string `json:"omaCustomerMobileNumber,omitempty"`
	OmaInstitute            string `json:"omaInstitute,omitempty"`
	OmaClientKey            string `json:"omaClientKey,omitempty"`
	CreatedAt               string `json:"createdAt,omitempty"`
	ModifiedAt              string `json:"modifiedAt,omitempty"`
}

type ApiKeyInfo struct {
	MID                   int
	APIKey                string
	AESKey                string
	Version               string
	GeneratedEncryptedKey string
}
