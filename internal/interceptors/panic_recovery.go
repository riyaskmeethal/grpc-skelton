package interceptors

import (
	"context"

	"google.golang.org/grpc"
	"osh.com/rps/registrar/utils/codes"
	ErrorMsg "osh.com/rps/registrar/utils/errormessages"
)

func (is InterceptorService) PanicRecoveryInterceptor() grpc.UnaryServerInterceptor {

	return func(
		ctx context.Context,
		req interface{},
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (resp interface{}, err error) {

		log := is.mlog.GetLogger()

		methodName := info.FullMethod

		defer func() {
			if r := recover(); r != nil {
				resp = interceptionError(req, codes.ResourceExhausted.String(), string(ErrorMsg.Internal_Error))
				log.Error("Recovered from Error occured in : ", methodName)
				return
			}
		}()
		return handler(ctx, req)
	}
}
