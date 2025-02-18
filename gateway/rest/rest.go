package rest

import (
	"context"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"osh.com/rps/registrar/config"
	"osh.com/rps/registrar/internal/interfaces"
)

func StartRestServer(ctx context.Context, conf interfaces.ConfigInterface, log interfaces.LogInterface) error {

	gateway := new(config.GateWayConfig)
	err := gateway.LoadGateWayConfig(conf, log)
	if err != nil {
		log.Error("failed loading gateway config")
		return err
	}

	if gateway.RestEnable {

		log.Info("Setting up REST gateway for grpc server.")

		mux := runtime.NewServeMux(
			runtime.WithIncomingHeaderMatcher(CustomMatcher),
		)

		var dialOption []grpc.DialOption
		transportOption := grpc.WithTransportCredentials(insecure.NewCredentials())
		dialOption = append(dialOption, transportOption)

		// err = Pb.RegisterFromEndpoint(ctx, mux, gateway.GrpcHost, dialOption)
		// if err != nil {
		// 	log.Error("Failed to register GRPC gateway.")
		// 	return err
		// }

		server := http.Server{Addr: gateway.RestHost, Handler: mux}
		log.Info("Rest server listening on ", gateway.RestHost)

		if err := server.ListenAndServe(); err != nil {
			log.Error("Failed serve rest serever on :.", gateway.RestHost)
			return err
		}

	}

	return nil
}

func CustomMatcher(key string) (string, bool) {
	switch key {
	case "Omaterminalid":
		return "omaTerminalId", true
	case "Omaclientuid":
		return "omaClientUid", true
	case "Omaclientkey":
		return "omaClientKey", true
	case "Omauniqueuid":
		return "omaUniqueUid", true
	case "Omainstitute":
		return "omaInstitute", true
	case "Omamid":
		return "omaMid", true
	case "Omakeyversion":
		return "omaKeyVersion", true
	case "Omasecretkey":
		return "omaSecretKey", true
	case "Omarereg":
		return "omaRereg", true
	default:
		return key, false
	}
}
