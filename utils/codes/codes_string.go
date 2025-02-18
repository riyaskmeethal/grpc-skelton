package codes

/// Need to change to decimal codes from string codes
func (c ErrorCode) String() string {
	return string(c)
	// switch c {
	// case OK:
	// 	return "000"
	// case Transaction_Processing:
	// 	return "001"
	// case Invalid_Amount_in_Request:
	// 	return "002"
	// case Transaction_reversed:
	// 	return "003"
	// case InvalidArgument:
	// 	return "004"
	// case InvalidRequest:
	// 	return "005"
	// case InvalidVoidRequest:
	// 	return "006"
	// case InvalidRefundRequest:
	// 	return "007"
	// case Invalid_Service_Header:
	// 	return "008"
	// case NotFound:
	// 	return "201"
	// case Duplicate_reff:
	// 	return "202"
	// case AlreadyExists:
	// 	return "203"
	// case No_Sale_Found:
	// 	return "204"
	// case No_Void_Found:
	// 	return "205"
	// case No_Refund_Found:
	// 	return "206"
	// case OmaService_Not_Found:
	// 	return "207"
	// case Already_Voided:
	// 	return "601"
	// case Already_Delined:
	// 	return "602"
	// case Already_Refunded:
	// 	return "603"
	// case Already_Timed_Out:
	// 	return "604"
	// case TimeOut:
	// 	return "701"
	// case Cancelled:
	// 	return "702"
	// case InvalidHeaderDetails:
	// 	return "817"
	// case Invalid_TerminalId:
	// 	return "818"
	// case Invalid_client_Reg:
	// 	return "819"
	// case Unverified_terminalid:
	// 	return "820"
	// case Invalid_aes_key:
	// 	return "821"
	// case Invalid_payload:
	// 	return "822"
	// case Unverified_ClientID:
	// 	return "823"
	// case Invalid_pos_reg:
	// 	return "824"
	// case Encryption_error:
	// 	return "825"
	// case Decryption_error:
	// 	return "826"
	// case Unexpected_err:
	// 	return "999"
	// case ResourceExhausted:
	// 	return "501"
	// case Aborted:
	// 	return "502"
	// case Internal:
	// 	return "503"
	// case Declined:
	// 	return "504"
	// default:
	// 	return c.String()
	// }
}

// var strToCodes = map[string]ErrorCode{
// 	"000": OK,
// 	"001": Transaction_Processing,
// 	"002": Invalid_Amount_in_Request,
// 	"003": Transaction_reversed,
// 	"004": InvalidArgument,
// 	"201": NotFound,
// 	"202": Duplicate_reff,
// 	"203": AlreadyExists,
// 	"601": Already_Voided,
// 	"602": Already_Delined,
// 	"701": TimeOut,
// 	"702": Cancelled,
// 	"817": InvalidHeaderDetails,
// 	"818": Invalid_TerminalId,
// 	"819": Invalid_client_Reg,
// 	"820": Unverified_terminalid,
// 	"821": Invalid_aes_key,
// 	"822": Invalid_payload,
// 	"823": Unverified_ClientID,
// 	"824": Invalid_pos_reg,
// 	"999": Unexpected_err,
// 	"501": ResourceExhausted,
// 	"502": Aborted,
// 	"503": Internal,
// 	"504": Declined,
// }

// func StringToCode(str string) ErrorCode {
// 	return strToCodes[str]
// }
