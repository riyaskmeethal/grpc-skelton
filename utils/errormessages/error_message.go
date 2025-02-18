package errormessages

type ValidationError string
type OmaErrorMessage string
type RegErrorMessage string
type OmaServiceError string

const (
	Invalid_Arguments ValidationError = "Invalid Argument"
	Invalid_Amount    ValidationError = "Amount is invalid"
	// Amount_Not_Found        ValidationError = "Amount not provided"
	// Invalid_TerminalID      ValidationError = "Terminal ID is invalid"
	TerminalId_Not_Found    ValidationError = "Terminal ID not provided"
	TXN_MW_Req_ID_Not_Found ValidationError = "Transaction middleware request ID not provided"
	Cli_Ref_Not_Found       ValidationError = "Client Refference Id not provided"
	Miss_Match_Cred         ValidationError = "Miss match in given details"
	Invalid_Request         ValidationError = "Invalid request to the server"
	Invalid_Void_request    ValidationError = "Invalid void request on the sale record"
	Invalid_Refund_request  ValidationError = "Invalid refund request on the sale record"
	Invalid_Error_Code      ValidationError = "Invalid Error Code"
	// Invalid_Success_Flag    ValidationError = "Invalid IsTransactionSuccess Field"
	Invalid_Sale_Details   ValidationError = "Invalid Sale Details"
	SerialNumber_Not_Found ValidationError = "SerialNumber not provided"
	MerchantId_Not_Found   ValidationError = "Merchant ID not provided"
)

const (
	// Sale_Created       OmaErrorMessage = "Sale record created successfully"
	Sale_Updated          OmaErrorMessage = "Transaction record updated successfully"
	Sale_Not_Found        OmaErrorMessage = "Transaction record not found"
	Sale_Timeout          OmaErrorMessage = "Transaction record found with TimeOut"
	Sale_Success          OmaErrorMessage = "Transaction record found with Success"
	Sale_Declined         OmaErrorMessage = "Transaction request has been declined"
	Sale_DB_Error         OmaErrorMessage = "Transaction request has declined, had error on DB side"
	Sale_Pending_Pos      OmaErrorMessage = "Transaction initiated, Pending action from POS"
	Sale_pending_Host     OmaErrorMessage = "Payment initiated by POS, Pending action from HOST"
	Already_Declined      OmaErrorMessage = "Transaction already declined"
	Already_Voided        OmaErrorMessage = "Transaction already Voided"
	Already_Authorized    OmaErrorMessage = "Transaction already authorized"
	Already_Refunded      OmaErrorMessage = "Transaction already refunded"
	Already_Timed_Out     OmaErrorMessage = "Transaction already Timed out"
	Void_Approved         OmaErrorMessage = "Transaction voided"
	Void_Pending_Pos      OmaErrorMessage = "Void sale initiated, Pending action from POS"
	Void_Pending_Host     OmaErrorMessage = "Void Initiated by POS, pending from HOST"
	Void_Declined         OmaErrorMessage = "Void request has delined"
	Void_Timeout          OmaErrorMessage = "Void request has timed out"
	Refund_Approved       OmaErrorMessage = "Transaction refunded"
	Refund_Pending_Pos    OmaErrorMessage = "Refund sale initiated, Pending action from POS"
	Refund_Pending_Host   OmaErrorMessage = "Refund Initiated by POS, pending from HOST"
	Refund_Timeout        OmaErrorMessage = "Refund request has timed out"
	Refund_Declined       OmaErrorMessage = "Refund request has Diclined"
	Internal_Error        OmaErrorMessage = "Server has met with an internal error"
	Sale_In_Progress      OmaErrorMessage = "Sale in progress"
	Void_In_Progress      OmaErrorMessage = "Void in progress"
	Refund_In_Progress    OmaErrorMessage = "Refund in progress"
	Void_Not_Found        OmaErrorMessage = "No new void found"
	Refund_Not_Found      OmaErrorMessage = "No new refund found"
	InvalidHeaderDetails  OmaErrorMessage = "Header details missing!"
	Invalid_Uid           OmaErrorMessage = "Invalid UID value in header !"
	Invalid_Institute     OmaErrorMessage = "Invalid Insitute value in Header !"
	Invalid_Mid           OmaErrorMessage = "Invalid Mid in Header !"
	Invalid_ReregVal      OmaErrorMessage = "Invalid omaRereg value in Header !"
	Invalid_KeyVersion    OmaErrorMessage = "Invalid key version in Header !"
	Invalid_SecretKey     OmaErrorMessage = "Invalid secret key in Header !"
	Invalid_Tid           OmaErrorMessage = "Invalid Termianl ID in Header !"
	Mid_Mismatch          OmaErrorMessage = "Mid in header does not match with Mid in request"
	Invalid_TerminalId    OmaErrorMessage = "Could not find any device for the credentials"
	Invalid_client_Reg    OmaErrorMessage = "Corresponding client is not verified"
	Unverified_terminalid OmaErrorMessage = "Pos terminal is not verified"
	Invalid_aes_key       OmaErrorMessage = "Found aes key miss match"
	Invalid_rsa_key       OmaErrorMessage = "Found invalid rsa key"
	Invalid_client_key    OmaErrorMessage = "Invalid Client Secret key. Key outdated, please re-register with updated key. Client validation failed !"
	Empty_payload         OmaErrorMessage = "Empty payload"
	Invalid_Enc_payload   OmaErrorMessage = "Invalid encrypted payload. Key outdated, please re-register for updated key."
	Invalid_payload       OmaErrorMessage = "Invalid payload"
	Unverified_ClientID   OmaErrorMessage = "Client ID unverified. Registration failed."
	Invalid_pos_reg       OmaErrorMessage = "Corresponding pos not verified"
	Encryption_error      OmaErrorMessage = "Encryption Error"
	Decryption_error      OmaErrorMessage = "Decryption Error"
	Terminal_Missmatch    OmaErrorMessage = "Terminal id given in header does not match with request body"
	TerminalId_Mapped     OmaErrorMessage = "Client is already registered to another TID. If required, please delink and re-register."
	Invalid_Txn_Type      OmaErrorMessage = "Transaction type value is invalid"
	No_InvoiceNo          OmaErrorMessage = "Invoice number required for this transaction."
	No_RRN                OmaErrorMessage = "RRN number required for this transaction."
	No_PaperRoll          OmaErrorMessage = "Empty Paper roll, please replace paper roll and initiate."
	Client_Txn_Cancel     OmaErrorMessage = "Transaction cancelled by Client. Please proceed with the transaction in cash."
	RPS_Timeout           OmaErrorMessage = "Client Timeout"
	MongoDB_UpdateErr     OmaErrorMessage = "MongoDB Update Error !"
	No_TipAmount          OmaErrorMessage = "Tip Amount required for this transaction."
)

const (
	DelinkRegister                   RegErrorMessage = "DeLink and Re-registration successful"
	Client_Reregister                RegErrorMessage = "Existing client successfully re-registered"
	Register_Success                 RegErrorMessage = "Registered successfully"
	Terminal_Exist                   RegErrorMessage = "Pos terminal already registered. If required, please delink and re-register during POS registration."
	ClientId_exist                   RegErrorMessage = "ECR client-ID in use already. If required, please delink and re-register during POS registration."
	SerialNumber_Already_Exist       RegErrorMessage = "Found Duplicate Terminal SerialNumber. If required, please delink and re-register during POS registration"
	SerialNumber_Already_Exist_NoReg RegErrorMessage = "Found serial number mapped with another Terminal without client registration. If required, please delink and re-register during POS registration"
	Sno_MissMatch                    RegErrorMessage = "Terminal Id already registerd with another serial number"
	Duplicate_Cid                    RegErrorMessage = "Found Duplicate client ID"
	Cid_MissMatch                    RegErrorMessage = "Terminal Id already registered with another Client ID. If required, please delink and re-register the Client."
	No_Pos_Found                     RegErrorMessage = "No Pos terminal found for the given credentials"
	Otp_expired                      RegErrorMessage = "OTP has timed out"
	Otp_missmatch                    RegErrorMessage = "OTP does not match"
	Otp_verified                     RegErrorMessage = "OTP verified"
	Pos_not_Auth                     RegErrorMessage = "Session request from unauthorized terminal"
	Session_key_generated            RegErrorMessage = "Session key created"
	Terminal_Already_Exists          RegErrorMessage = "Terminal ID already registered to another serial number. If required, please delink and re-register."
	Terminal_Client_Exist            RegErrorMessage = "Terminal already registered to client. If required, please delink and re-register."
	Invalid_Rereg                    RegErrorMessage = "OmaRereg cannot be true for first time registration"
	ClientId_Already_Exist           RegErrorMessage = "Client ID already exists and is mapped to another TID. If required, please delink and re-register."
)

const (
	Invalid_Service_Header OmaServiceError = "Invalid service header, need omaInstitute header"
	OmaService_Not_Found   OmaServiceError = "Service not found for the given omaInstitute"
	Invalid_Service        OmaServiceError = "Service not found/currently unavailable"
	MerchantId_MissMatch   OmaServiceError = "Merchant ID missmatch"
)

// var statusToErrMsg = map[status.Status]string{
// 	0: string(Sale_Pending_Pos),
// 	1: string(Sale_pending_Host),
// 	2: string(Sale_Success),
// 	3: string(Sale_Delined),
// 	4: string(Sale_Timeout),
// 	5: string(Sale_Not_Found),
// 	6: string(Void_Approved),
// 	7: string(Void_Declined),
// 	// `"REFUNDPENDING"`:   ,
// 	// `"REFUNDSUCCESS"`:   REFUNDSUCCESS,
// 	// `"REFUNDDECLINED"`:  REFUNDDECLINED,
// }

// func GetErrMsgByStatus(status status.Status) string {
// 	if msg, ok := statusToErrMsg[status]; ok {
// 		return msg
// 	}
// 	// ToDo :- make clear err msg are ok
// 	return "Error msg not found"
// }
