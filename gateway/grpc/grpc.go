package grpc

import (
	"context"
	"net"

	"google.golang.org/grpc"
	"osh.com/rps/registrar/config"
	"osh.com/rps/registrar/internal/interceptors"
	"osh.com/rps/registrar/internal/interfaces"
	"osh.com/rps/registrar/internal/server/controller"
)

var grpcServer *grpc.Server

func StartGrpcServer(ctx context.Context, conf interfaces.ConfigInterface, db interfaces.DbInterface, vDb interfaces.ValidationDbInterface, mlog interfaces.MultyLogInterface) error {

	log := mlog.GetLogger()

	gateway := new(config.GateWayConfig)
	err := gateway.LoadGateWayConfig(conf, log)
	if err != nil {
		log.Error("failed loading gateway config")
		return err
	}

	_ = controller.GetNewServer(conf, db, vDb, mlog) //SERVER IMPLEMENTATIONS
	log.Debug(" service loaded.")

	interceptorService := interceptors.GetNewInterceptorService(conf, db, mlog)

	serverOpts := grpc.ChainUnaryInterceptor(interceptorService.PanicRecoveryInterceptor())
	log.Debug("Registered interceptors.")

	grpcServer = grpc.NewServer(serverOpts)

	// someappPb.RegistersomAppServer(grpcServer, registrarServer)
	// log.Debug(" service registerd.")

	listener, err := net.Listen("tcp", gateway.GrpcHost)
	if err != nil {
		log.Error("Failed to listen: %s", err.Error())
		return err
	}

	log.Info("GRPC server started and listening on : ", listener.Addr().String())

	return grpcServer.Serve(listener)
}

func StopGrpcServer() {
	grpcServer.GracefulStop()
}
