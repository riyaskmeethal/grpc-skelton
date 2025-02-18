package mssql

import (
	"context"
	"errors"

	"osh.com/rps/registrar/internal/models"
)

func (dbs MsSqlDB) GetApiKeyDetails(institute, mid, secretKey, keyVersion string) ([]models.ApiKeyInfo, error) {
	return dbs.DbOps[institute].GetApiKeyDetails(mid, secretKey, keyVersion)
}

func (md MsSqlDB) ServiceAvailable(ctx context.Context, institute string) bool {
	_, ok := md.DbOps[institute]
	return ok
}

func (md MsSqlDB) CheckClientReregistration(ctx context.Context, institute string, mid string, tid string, cid string) (device models.Device, err error) {

	if md.ServiceAvailable(ctx, institute) {
		return md.DbOps[institute].CheckClientReregistration(ctx, mid, tid, cid)
	} else {
		err = errors.New("Service unavailable for " + institute)
	}
	return
}

func (md MsSqlDB) CheckPOSReregistration(ctx context.Context, institute string, mid string, tid string, slno string) (device models.Device, err error) {

	if md.ServiceAvailable(ctx, institute) {
		return md.DbOps[institute].CheckPOSReregistration(ctx, mid, tid, slno)
	} else {
		err = errors.New("Service unavailable for " + institute)
	}
	return
}

func (md MsSqlDB) FindTerminalByTID(ctx context.Context, institute string, mid string, tid string) (device models.Device, err error) {

	if md.ServiceAvailable(ctx, institute) {
		return md.DbOps[institute].FindTerminalByTID(ctx, mid, tid)
	} else {
		err = errors.New("Service unavailable for " + institute)
	}
	return
}

func (md MsSqlDB) DelinkTerminalByTIDifExist(ctx context.Context, institute string, mid string, tid string) error {

	if md.ServiceAvailable(ctx, institute) {
		return md.DbOps[institute].DelinkTerminalByTIDifExist(ctx, mid, tid)
	} else {
		return errors.New("Service unavailable for " + institute)
	}

}
func (md MsSqlDB) DelinkTerminalByCIDifExist(ctx context.Context, institute string, mid string, cid string) error {

	if md.ServiceAvailable(ctx, institute) {
		return md.DbOps[institute].DelinkTerminalByCIDifExist(ctx, mid, cid)
	} else {
		return errors.New("Service unavailable for " + institute)
	}

}

func (md MsSqlDB) DelinkTerminalBySLNOifExist(ctx context.Context, institute string, mid string, slno string) error {

	if md.ServiceAvailable(ctx, institute) {
		return md.DbOps[institute].DelinkTerminalBySLNOifExist(ctx, mid, slno)
	} else {
		return errors.New("Service unavailable for " + institute)
	}

}

func (md MsSqlDB) UpdateTerminalClient(ctx context.Context, institute string, device models.Device) error {

	if md.ServiceAvailable(ctx, institute) {
		return md.DbOps[institute].UpdateTerminalClient(ctx, device)
	} else {
		return errors.New("Service unavailable for " + institute)
	}

}
func (md MsSqlDB) CreateTerminalClient(ctx context.Context, institute string, device models.Device) error {

	if md.ServiceAvailable(ctx, institute) {
		return md.DbOps[institute].CreateTerminalClient(ctx, device)
	} else {
		return errors.New("Service unavailable for " + institute)
	}

}

func (md MsSqlDB) UpdateTerminalPOS(ctx context.Context, institute string, device models.Device) error {

	if md.ServiceAvailable(ctx, institute) {
		return md.DbOps[institute].UpdateTerminalPOS(ctx, device)
	} else {
		return errors.New("Service unavailable for " + institute)
	}

}
func (md MsSqlDB) CreateTerminalPOS(ctx context.Context, institute string, device models.Device) error {

	if md.ServiceAvailable(ctx, institute) {
		return md.DbOps[institute].CreateTerminalPOS(ctx, device)
	} else {
		return errors.New("Service unavailable for " + institute)
	}

}

func (md MsSqlDB) UpdatePosVerify(ctx context.Context, institute string, mid string, tid string, slno string) error {

	if md.ServiceAvailable(ctx, institute) {
		return md.DbOps[institute].UpdatePosVerify(ctx, mid, tid, slno)
	} else {
		return errors.New("Service unavailable for " + institute)
	}

}

func (md MsSqlDB) UpdateSessionKey(ctx context.Context, institute string, mid string, tid string, slno string, sessionKey []byte, uid string) error {

	if md.ServiceAvailable(ctx, institute) {
		return md.DbOps[institute].UpdateSessionKey(ctx, mid, tid, slno, sessionKey, uid)
	} else {
		return errors.New("Service unavailable for " + institute)
	}

}
