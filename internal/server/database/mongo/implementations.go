package mongo

import (
	"context"
	"errors"

	"osh.com/rps/registrar/internal/models"
)

func (md MongoDB) ServiceAvailable(ctx context.Context, institute string) bool {
	_, ok := md.Connect[institute]
	return ok
}

func (md MongoDB) CheckClientReregistration(ctx context.Context, institute string, mid string, tid string, cid string) (device models.Device, err error) {

	if md.ServiceAvailable(ctx, institute) {
		// return md.Connect[institute].CALL DB OPERATIONS(ctx, mid, tid, cid)
		return
	} else {
		err = errors.New("Service unavailable for " + institute)
	}
	return
}

func (md MongoDB) CheckPOSReregistration(ctx context.Context, institute string, mid string, tid string, slno string) (device models.Device, err error) {

	if md.ServiceAvailable(ctx, institute) {
		// return md.Connect[institute].CALL DB OPERATIONS(ctx, mid, tid, cid)
		return
	} else {
		err = errors.New("Service unavailable for " + institute)
	}
	return
}

func (md MongoDB) FindTerminalByTID(ctx context.Context, institute string, mid string, tid string) (device models.Device, err error) {

	if md.ServiceAvailable(ctx, institute) {
		// return md.Connect[institute].CALL DB OPERATIONS(ctx, mid, tid, cid)
		return
	} else {
		err = errors.New("Service unavailable for " + institute)
	}
	return
}

func (md MongoDB) DelinkTerminalByTIDifExist(ctx context.Context, institute string, mid string, tid string) error {

	if md.ServiceAvailable(ctx, institute) {
		// return md.Connect[institute].CALL DB OPERATIONS(ctx, mid, tid, cid)
		return nil
	} else {
		return errors.New("Service unavailable for " + institute)
	}

}
func (md MongoDB) DelinkTerminalByCIDifExist(ctx context.Context, institute string, mid string, cid string) error {

	if md.ServiceAvailable(ctx, institute) {
		// return md.Connect[institute].CALL DB OPERATIONS(ctx, mid, tid, cid)
		return nil
	} else {
		return errors.New("Service unavailable for " + institute)
	}

}

func (md MongoDB) DelinkTerminalBySLNOifExist(ctx context.Context, institute string, mid string, slno string) error {

	if md.ServiceAvailable(ctx, institute) {
		// return md.Connect[institute].CALL DB OPERATIONS(ctx, mid, tid, cid)
		return nil
	} else {
		return errors.New("Service unavailable for " + institute)
	}

}

func (md MongoDB) UpdateTerminalClient(ctx context.Context, institute string, device models.Device) error {

	if md.ServiceAvailable(ctx, institute) {
		// return md.Connect[institute].CALL DB OPERATIONS(ctx, mid, tid, cid)
		return nil
	} else {
		return errors.New("Service unavailable for " + institute)
	}

}
func (md MongoDB) CreateTerminalClient(ctx context.Context, institute string, device models.Device) error {

	if md.ServiceAvailable(ctx, institute) {
		// return md.Connect[institute].CALL DB OPERATIONS(ctx, mid, tid, cid)
		return nil
	} else {
		return errors.New("Service unavailable for " + institute)
	}

}

func (md MongoDB) UpdateTerminalPOS(ctx context.Context, institute string, device models.Device) error {

	if md.ServiceAvailable(ctx, institute) {
		// return md.Connect[institute].CALL DB OPERATIONS(ctx, mid, tid, cid)
		return nil
	} else {
		return errors.New("Service unavailable for " + institute)
	}

}
func (md MongoDB) CreateTerminalPOS(ctx context.Context, institute string, device models.Device) error {

	if md.ServiceAvailable(ctx, institute) {
		// return md.Connect[institute].CALL DB OPERATIONS(ctx, mid, tid, cid)
		return nil
	} else {
		return errors.New("Service unavailable for " + institute)
	}

}

func (md MongoDB) UpdatePosVerify(ctx context.Context, institute string, mid string, tid string, slno string) error {

	if md.ServiceAvailable(ctx, institute) {
		// return md.Connect[institute].CALL DB OPERATIONS(ctx, mid, tid, cid)
		return nil
	} else {
		return errors.New("Service unavailable for " + institute)
	}

}

func (md MongoDB) UpdateSessionKey(ctx context.Context, institute string, mid string, tid string, slno string, sessionKey []byte, uid string) error {

	if md.ServiceAvailable(ctx, institute) {
		// return md.Connect[institute].CALL DB OPERATIONS(ctx, mid, tid, cid)
		return nil
	} else {
		return errors.New("Service unavailable for " + institute)
	}

}
