package interceptors

import (
	"osh.com/rps/registrar/internal/interfaces"
)

type InterceptorService struct {
	registrarConnection interfaces.DbInterface
	mlog                interfaces.MultyLogInterface
	validationDb        interfaces.ValidationDbInterface
}

func GetNewInterceptorService(conf interfaces.ConfigInterface, db interfaces.DbInterface, mlog interfaces.MultyLogInterface) InterceptorService {

	return InterceptorService{
		registrarConnection: db,
		mlog:                mlog,
	}

}
