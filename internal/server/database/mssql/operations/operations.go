package operations

import (
	"context"
	"database/sql"
	"encoding/hex"
	"fmt"

	"github.com/google/uuid"
	"osh.com/rps/registrar/internal/models"
)

func (do DbOps) GetApiKeyDetails(mid, secretKey, keyVersion string) ([]models.ApiKeyInfo, error) {
	ApiKeyList := []models.ApiKeyInfo{}
	query := `SELECT MID, API_Key, AES_Key, Version FROM merchant_keys WHERE MID = @mid`

	rows, err := do.DBconn.Query(query, sql.Named("mid", mid))
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var row models.ApiKeyInfo
		err := rows.Scan(&row.MID, &row.APIKey, &row.AESKey, &row.Version)
		if err != nil {
			return nil, err
		}
		ApiKeyList = append(ApiKeyList, row)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}
	do.log.Info("api key detaials fetched from db.")

	return ApiKeyList, nil
}

func (do DbOps) CheckClientReregistration(ctx context.Context, mid string, tid string, cid string) (device models.Device, err error) {

	device = models.Device{}

	var rpsPrivateKey, posPublicKey, clientPublicKey, clientAesKey, posAesKey string

	queryString := `SELECT 
						COALESCE(terminal_id,''),
						COALESCE(client_id,''),
						COALESCE(serial_number,''),
						COALESCE(merchant_id,''),
						COALESCE(rps_private_key,''),
						COALESCE(pos_public_key,''),
						COALESCE(client_public_key,''),
						COALESCE(pos_aes_key,''),
						COALESCE(client_aes_key,''),
						COALESCE(pos_aes_uid,''),
						COALESCE(client_aes_uid,''),
						COALESCE(is_terminal_verified,'FALSE'),
						COALESCE(is_client_verified,'FALSE'),
						COALESCE(customer_email,''),
						COALESCE(customer_mobile_number,''),
						COALESCE(institute,''),
						COALESCE(client_key,'')	
					FROM devices
					WHERE terminal_id = @p1 AND merchant_id = @p2 AND client_id = @p3	AND is_rereg = 'FALSE'
	`
	err = do.DBconn.QueryRowContext(ctx, queryString, tid, mid, cid).Scan(
		&device.OmaTerminalID,
		&device.OmaClientID,
		&device.OmaSerialNumber,
		&device.OmaMerchantId,
		&rpsPrivateKey,
		&posPublicKey,
		&clientPublicKey,
		&posAesKey,
		&clientAesKey,
		&device.PosAesUid,
		&device.ClientAesUid,
		&device.IsTerminalVerified,
		&device.IsClientVerified,
		&device.OmaCustomerEmail,
		&device.OmaCustomerMobileNumber,
		&device.OmaInstitute,
		&device.OmaClientKey,
	)

	if err != nil {
		do.log.Error(err.Error())
		return
	}

	decodedPvt, err := hex.DecodeString(rpsPrivateKey)
	if err != nil {
		return
	}
	decodedPosPblc, err := hex.DecodeString(posPublicKey)
	if err != nil {
		return
	}
	decodedCLntPub, err := hex.DecodeString(clientPublicKey)
	if err != nil {
		return
	}
	decodedClntAES, err := hex.DecodeString(clientAesKey)
	if err != nil {
		return
	}
	decodedposAes, err := hex.DecodeString(posAesKey)
	if err != nil {
		return
	}
	device.ClientPublicKey = decodedCLntPub
	device.PosAesKey = decodedposAes
	device.ClientAesKey = decodedClntAES
	device.RpsPrivateKey = decodedPvt
	device.PosPublicKey = decodedPosPblc
	do.log.Info("Terminal found re-registering client : ", device.OmaTerminalID)
	return

}

func (do DbOps) CheckPOSReregistration(ctx context.Context, mid string, tid string, slno string) (device models.Device, err error) {

	device = models.Device{}

	var rpsPrivateKey, posPublicKey, clientPublicKey, clientAesKey, posAesKey string

	queryString := `SELECT 
						COALESCE(terminal_id,''),
						COALESCE(client_id,''),
						COALESCE(serial_number,''),
						COALESCE(merchant_id,''),
						COALESCE(rps_private_key,''),
						COALESCE(pos_public_key,''),
						COALESCE(client_public_key,''),
						COALESCE(pos_aes_key,''),
						COALESCE(client_aes_key,''),
						COALESCE(pos_aes_uid,''),
						COALESCE(client_aes_uid,''),
						COALESCE(is_terminal_verified,'FALSE'),
						COALESCE(is_client_verified,'FALSE'),
						COALESCE(customer_email,''),
						COALESCE(customer_mobile_number,''),
						COALESCE(institute,''),
						COALESCE(client_key,'')	
					FROM devices
					WHERE terminal_id = @p1 AND merchant_id = @p2 AND serial_number = @p3	AND is_rereg = 'FALSE'
	`
	err = do.DBconn.QueryRowContext(ctx, queryString, tid, mid, slno).Scan(
		&device.OmaTerminalID,
		&device.OmaClientID,
		&device.OmaSerialNumber,
		&device.OmaMerchantId,
		&rpsPrivateKey,
		&posPublicKey,
		&clientPublicKey,
		&posAesKey,
		&clientAesKey,
		&device.PosAesUid,
		&device.ClientAesUid,
		&device.IsTerminalVerified,
		&device.IsClientVerified,
		&device.OmaCustomerEmail,
		&device.OmaCustomerMobileNumber,
		&device.OmaInstitute,
		&device.OmaClientKey,
	)

	if err != nil {
		do.log.Error(err.Error())
		return
	}

	decodedPvt, err := hex.DecodeString(rpsPrivateKey)
	if err != nil {
		return
	}
	decodedPosPblc, err := hex.DecodeString(posPublicKey)
	if err != nil {
		return
	}
	decodedCLntPub, err := hex.DecodeString(clientPublicKey)
	if err != nil {
		return
	}
	decodedClntAES, err := hex.DecodeString(clientAesKey)
	if err != nil {
		return
	}
	decodedposAes, err := hex.DecodeString(posAesKey)
	if err != nil {
		return
	}

	device.ClientPublicKey = decodedCLntPub
	device.PosAesKey = decodedposAes
	device.ClientAesKey = decodedClntAES
	device.RpsPrivateKey = decodedPvt
	device.PosPublicKey = decodedPosPblc
	do.log.Info("Terminal found re-registering POS : ", device.OmaTerminalID)
	return

}

func (do DbOps) FindTerminalByTID(ctx context.Context, mid string, tid string) (device models.Device, err error) {

	device = models.Device{}

	var rpsPrivateKey, posPublicKey, clientPublicKey, clientAesKey, posAesKey string

	queryString := `SELECT 
						COALESCE(terminal_id,''),
						COALESCE(client_id,''),
						COALESCE(serial_number,''),
						COALESCE(merchant_id,''),
						COALESCE(rps_private_key,''),
						COALESCE(pos_public_key,''),
						COALESCE(client_public_key,''),
						COALESCE(pos_aes_key,''),
						COALESCE(client_aes_key,''),
						COALESCE(pos_aes_uid,''),
						COALESCE(client_aes_uid,''),
						COALESCE(otp,''),
						COALESCE(otp_expires_at,''),
						COALESCE(is_terminal_verified,'FALSE'),
						COALESCE(is_client_verified,'FALSE'),
						COALESCE(customer_email,''),
						COALESCE(customer_mobile_number,''),
						COALESCE(institute,''),
						COALESCE(client_key,'')
					FROM devices
					WHERE terminal_id = @p1 AND merchant_id = @p2 AND is_rereg = 'FALSE'
	`
	err = do.DBconn.QueryRowContext(ctx, queryString, tid, mid).Scan(
		&device.OmaTerminalID,
		&device.OmaClientID,
		&device.OmaSerialNumber,
		&device.OmaMerchantId,
		&rpsPrivateKey,
		&posPublicKey,
		&clientPublicKey,
		&posAesKey,
		&clientAesKey,
		&device.PosAesUid,
		&device.ClientAesUid,
		&device.OmaOTP,
		&device.OtpExpiresAt,
		&device.IsTerminalVerified,
		&device.IsClientVerified,
		&device.OmaCustomerEmail,
		&device.OmaCustomerMobileNumber,
		&device.OmaInstitute,
		&device.OmaClientKey,
	)

	if err != nil {
		do.log.Error(err.Error())
		return
	}

	decodedPvt, err := hex.DecodeString(rpsPrivateKey)
	if err != nil {
		return
	}
	decodedPosPblc, err := hex.DecodeString(posPublicKey)
	if err != nil {
		return
	}
	decodedCLntPub, err := hex.DecodeString(clientPublicKey)
	if err != nil {
		return
	}
	decodedClntAES, err := hex.DecodeString(clientAesKey)
	if err != nil {
		return
	}
	decodedposAes, err := hex.DecodeString(posAesKey)
	if err != nil {
		return
	}

	device.ClientPublicKey = decodedCLntPub
	device.PosAesKey = decodedposAes
	device.ClientAesKey = decodedClntAES
	device.RpsPrivateKey = decodedPvt
	device.PosPublicKey = decodedPosPblc

	do.log.Info("Found terminal with TID : ", device.OmaTerminalID)
	return

}

func (do DbOps) DelinkTerminalByTIDifExist(ctx context.Context, mid string, tid string) (err error) {

	txn, err := do.DBconn.BeginTx(ctx, nil)
	if err != nil {
		return
	}
	query := `UPDATE devices
			  SET
			    is_rereg = 'TRUE'
			  WHERE 
			  	terminal_id = @terminal_id AND
				merchant_id = @merchant_id AND
				is_rereg = 'FALSE';	
	`
	result, err := txn.ExecContext(ctx, query, sql.Named("merchant_id", mid), sql.Named("terminal_id", tid))
	if err != nil {
		txn.Rollback()
		do.log.Error("failed updating : ", err.Error())
		return err
	}

	// Get the number of affected rows
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		txn.Rollback()
		do.log.Error("failed to get affected rows: ", err.Error())
		return
	}

	if err = txn.Commit(); err != nil {
		do.log.Error("Error committing transaction:", err.Error())
		return
	}

	if rowsAffected != 0 {
		do.log.Info("De-Linked terminal with terminal id : ", tid)
	}

	return
}

func (do DbOps) DelinkTerminalByCIDifExist(ctx context.Context, mid string, cid string) (err error) {

	txn, err := do.DBconn.BeginTx(ctx, nil)
	if err != nil {
		return
	}
	query := `UPDATE devices
			  SET
			    is_rereg = 'TRUE'
			  WHERE 
			  	client_id = @client_id AND
				merchant_id = @merchant_id AND
				is_rereg = 'FALSE';	
	`
	result, err := txn.ExecContext(ctx, query, sql.Named("merchant_id", mid), sql.Named("client_id", cid))
	if err != nil {
		txn.Rollback()
		do.log.Error("failed updating : ", err.Error())
		return err
	}

	// Get the number of affected rows
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		txn.Rollback()
		do.log.Error("failed to get affected rows: ", err.Error())
		return
	}

	if err = txn.Commit(); err != nil {
		do.log.Error("Error committing transaction:", err.Error())
		return
	}

	if rowsAffected != 0 {
		do.log.Info("De-Linked terminal with client id : ", cid)
	}

	return
}

func (do DbOps) DelinkTerminalBySLNOifExist(ctx context.Context, mid string, slno string) (err error) {

	txn, err := do.DBconn.BeginTx(ctx, nil)
	if err != nil {
		return
	}
	query := `UPDATE devices
			  SET
			   is_rereg = 'TRUE'
			  WHERE 
			  	serial_number = @serial_number AND
				merchant_id = @merchant_id AND
				is_rereg = 'FALSE';	
	`
	result, err := txn.ExecContext(ctx, query, sql.Named("merchant_id", mid), sql.Named("serial_number", slno))
	if err != nil {
		txn.Rollback()
		do.log.Error("failed updating : ", err.Error())
		return err
	}

	// Get the number of affected rows
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		txn.Rollback()
		do.log.Error("failed to get affected rows: ", err.Error())
		return
	}

	if err = txn.Commit(); err != nil {
		do.log.Error("Error committing transaction:", err.Error())
		return
	}

	if rowsAffected != 0 {
		do.log.Info("De-Linked terminal with serial number : ", slno)
	}

	return
}

func (do DbOps) UpdateTerminalClient(ctx context.Context, device models.Device) (err error) {
	//CONFIG TIME FORMAT
	// timeFormat := do.conf.TimeFormat
	clientPublicKey := hex.EncodeToString(device.ClientPublicKey)
	clientAesKey := hex.EncodeToString(device.ClientAesKey)

	txn, err := do.DBconn.BeginTx(ctx, nil)
	if err != nil {
		return
	}

	query := `
			UPDATE 
				devices
			SET 
				client_id = @client_id,
				client_public_key = @client_public_key,
				client_aes_key = @client_aes_key,
				customer_email = @customer_email,
				customer_mobile_number = @customer_mobile_number,
				is_client_verified = @is_client_verified,
				is_rereg = @is_rereg
			WHERE 
				terminal_id = @terminal_id AND 
				merchant_id = @merchant_id AND 
				is_rereg = 'FALSE';	
			`
	if _, err = txn.ExecContext(ctx, query,
		sql.Named("client_id", device.OmaClientID),
		sql.Named("client_public_key", clientPublicKey),
		sql.Named("client_aes_key", clientAesKey),
		sql.Named("customer_email", device.OmaCustomerEmail),
		sql.Named("customer_mobile_number", device.OmaCustomerMobileNumber),
		sql.Named("is_client_verified", device.IsClientVerified),
		sql.Named("is_rereg", device.IsRereg),
		sql.Named("terminal_id", device.OmaTerminalID),
		sql.Named("merchant_id", device.OmaMerchantId),
	); err != nil {
		txn.Rollback()
		do.log.Error("failed update : ", err.Error())
		return
	}

	if err = txn.Commit(); err != nil {
		do.log.Error("Error committing transaction:", err.Error())
		return
	}
	do.log.Info("Updated client with client id : ", device.OmaClientID)

	return

}
func (do DbOps) CreateTerminalClient(ctx context.Context, device models.Device) (err error) {
	//CONFIG TIME FORMAT
	// timeFormat := do.conf.TimeFormat
	clientPublicKey := hex.EncodeToString(device.ClientPublicKey)
	clientAesKey := hex.EncodeToString(device.ClientAesKey)

	txn, err := do.DBconn.BeginTx(ctx, nil)
	if err != nil {
		do.log.Error("Failed intiating transaction :", err.Error())
		return
	}

	queryString := `
		INSERT INTO devices (
			id,terminal_id,client_id,merchant_id,client_public_key,client_aes_key,
			is_rereg,is_client_verified, customer_email,customer_mobile_number, institute, client_key
		)
		VALUES (
			@id,@terminal_id, @client_id, @merchant_id, @client_public_key, @client_aes_key,
		    @is_rereg, @is_client_verified, @customer_email,@customer_mobile_number, @institute, @client_key
		);	
	`
	_, err = txn.ExecContext(ctx, queryString,
		sql.Named("id", uuid.New()),
		sql.Named("terminal_id", device.OmaTerminalID),
		sql.Named("client_id", device.OmaClientID),
		sql.Named("merchant_id", device.OmaMerchantId),
		sql.Named("client_public_key", clientPublicKey),
		sql.Named("client_aes_key", clientAesKey),
		sql.Named("is_rereg", device.IsRereg),
		sql.Named("is_client_verified", device.IsClientVerified),
		sql.Named("customer_email", device.OmaCustomerEmail),
		sql.Named("customer_mobile_number", device.OmaCustomerMobileNumber),
		sql.Named("institute", device.OmaInstitute),
		sql.Named("client_key", device.OmaClientKey),
	)

	if err != nil {
		txn.Rollback()
		do.log.Error("transaction failed.", err.Error())
		return
	}

	err = txn.Commit()
	if err != nil {
		txn.Rollback()
		do.log.Error("transaction commit failed.")
		return
	}
	do.log.Info("registered.")

	do.log.Info("Created terminal MID ", device.OmaMerchantId, ", TID", device.OmaTerminalID, "and CID", device.OmaClientID)
	return

}

func (do DbOps) UpdateTerminalPOS(ctx context.Context, device models.Device) (err error) {
	//CONFIG TIME FORMAT
	// timeFormat := do.conf.TimeFormat

	rpsPrivateKey := hex.EncodeToString(device.RpsPrivateKey)
	posPublicKey := hex.EncodeToString(device.PosPublicKey)

	txn, err := do.DBconn.BeginTx(ctx, nil)
	if err != nil {
		return
	}

	query := `UPDATE devices
			SET 
				serial_number =  @serial_number,
				otp =            @otp,
				otp_expires_at =       @otp_expires_at,
				rps_private_key =      @rps_private_key,
				pos_public_key =        @pos_public_key
			WHERE terminal_id = @terminal_id AND is_rereg = 'FALSE';	
			`
	if _, err = txn.ExecContext(ctx, query,
		sql.Named("serial_number", device.OmaSerialNumber),
		sql.Named("otp", device.OmaOTP),
		sql.Named("otp_expires_at", device.OtpExpiresAt),
		sql.Named("rps_private_key", rpsPrivateKey),
		sql.Named("pos_public_key", posPublicKey),
		sql.Named("terminal_id", device.OmaTerminalID),
	); err != nil {
		txn.Rollback()
		do.log.Error("failed update : ", err.Error())
		return
	}

	if err = txn.Commit(); err != nil {
		do.log.Error("Error committing transaction:", err.Error())
		return
	}
	do.log.Info("Updated terminal POS with serial number : ", device.OmaSerialNumber)
	return
}
func (do DbOps) CreateTerminalPOS(ctx context.Context, device models.Device) (err error) {
	//CONFIG TIME FORMAT
	// timeFormat := do.conf.TimeFormat

	rpsPrivateKey := hex.EncodeToString(device.RpsPrivateKey)
	posPublicKey := hex.EncodeToString(device.PosPublicKey)
	posAesKey := hex.EncodeToString(device.PosAesKey)

	txn, err := do.DBconn.BeginTx(ctx, nil)
	if err != nil {
		fmt.Println("Failed intiating transaction :", err.Error())
		return
	}

	queryString := `
		INSERT INTO devices (
			id,terminal_id,serial_number,merchant_id,rps_private_key,pos_public_key,
			pos_aes_key,otp,otp_expires_at, is_rereg, customer_email,
			customer_mobile_number, institute
		)
		VALUES (
			@id,@terminal_id, @serial_number, @merchant_id, @rps_private_key, @pos_public_key,
			@pos_aes_key, @otp,@otp_expires_at, @is_rereg, @customer_email,
			@customer_mobile_number, @institute
		);	
	`
	_, err = txn.ExecContext(ctx, queryString,
		sql.Named("id", uuid.New()),
		sql.Named("terminal_id", device.OmaTerminalID),
		sql.Named("serial_number", device.OmaSerialNumber),
		sql.Named("merchant_id", device.OmaMerchantId),
		sql.Named("rps_private_key", rpsPrivateKey),
		sql.Named("pos_public_key", posPublicKey),
		sql.Named("pos_aes_key", posAesKey),
		sql.Named("otp", device.OmaOTP),
		sql.Named("otp_expires_at", device.OtpExpiresAt),
		sql.Named("is_rereg", device.IsRereg),
		sql.Named("customer_email", device.OmaCustomerEmail),
		sql.Named("customer_mobile_number", device.OmaCustomerMobileNumber),
		sql.Named("institute", device.OmaInstitute),
	)

	if err != nil {
		txn.Rollback()
		do.log.Error("transaction failed.", err.Error())
		return
	}

	err = txn.Commit()
	if err != nil {
		txn.Rollback()
		do.log.Error("transaction commit failed.")
	}

	do.log.Info("Created terminal MID ", device.OmaMerchantId, ", TID", device.OmaTerminalID, "and SLNo", device.OmaSerialNumber)
	return

}

func (do DbOps) UpdatePosVerify(ctx context.Context, mid string, tid string, slno string) (err error) {
	//TODO: CONFIG TIME FORMAT,mid need to make required in VerifyPosOtpRequest proto for removing the condition
	// timeFormat := do.conf.TimeFormat

	txn, err := do.DBconn.BeginTx(ctx, nil)
	if err != nil {
		return
	}
	query := `UPDATE devices
			  SET
			   is_terminal_verified = 'TRUE'
			  WHERE 
			  	terminal_id = @terminal_id AND 
				serial_number = @serial_number AND
				merchant_id = @merchant_id AND 
				is_rereg = 'FALSE';	
			`
	if _, err = txn.ExecContext(ctx, query,
		sql.Named("terminal_id", tid),
		sql.Named("serial_number", slno),
		sql.Named("merchant_id", mid)); err != nil {
		txn.Rollback()
		do.log.Error("failed updating : ", err.Error())
		return
	}

	if err = txn.Commit(); err != nil {
		do.log.Error("Error committing transaction:", err.Error())
		return
	}
	do.log.Info("Terminal POS verified for terminal id : ", tid)
	return

}

func (do DbOps) UpdateSessionKey(ctx context.Context, mid string, tid string, slno string, sessionKey []byte, uid string) (err error) {
	txn, err := do.DBconn.BeginTx(ctx, nil)
	if err != nil {
		return
	}

	query := `UPDATE devices
			SET 
				pos_aes_key = @pos_aes_key,
				pos_aes_uid = @pos_aes_uid				
			WHERE 
				terminal_id = @terminal_id AND 
				merchant_id = @merchant_id AND 
				is_rereg = 'FALSE';	
			`
	if _, err = txn.ExecContext(ctx, query,
		sql.Named("pos_aes_key", sessionKey),
		sql.Named("pos_aes_uid", uid),
		sql.Named("merchant_id", mid),
		sql.Named("terminal_id", tid),
	); err != nil {
		txn.Rollback()
		do.log.Error("failed update : ", err.Error())
		return
	}

	if err = txn.Commit(); err != nil {
		do.log.Error("Error committing transaction:", err.Error())
		return
	}
	do.log.Info("Session key updated for terminal id : ", tid)
	return
}
