package main

import (
	"context"
	"flag"
	"os"
	"os/signal"
	"syscall"

	"github.com/rs/zerolog/log"
	"osh.com/rps/registrar/gateway/grpc"
	"osh.com/rps/registrar/gateway/rest"
	"osh.com/rps/registrar/internal/interfaces"
	"osh.com/rps/registrar/internal/pkg"
	"osh.com/rps/registrar/internal/server/database"
)

const DEV_CONFIG = "appConfig.yaml"

func main() {

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	confFile, initConf := getConfFile()

	conf, err := pkg.GetConfig(confFile, initConf)
	if err != nil {
		log.Fatal().Err(err).Msgf("Could not load configurations: %v", err)
	}
	log.Logger.Info().Msg("Configurations loaded successfully.")

	ml := pkg.GetMultilogger(conf)

	zlog := ml.GetLogger()

	db, vDb := database.GetDb(ctx, conf, ml)

	osSignalCh := make(chan os.Signal, 1)
	signal.Notify(osSignalCh, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		if err := rest.StartRestServer(ctx, conf, zlog); err != nil {
			zlog.Error("From REST server :", err.Error())
			osSignalCh <- syscall.SIGTERM
		}
	}()

	go func() {
		if err := grpc.StartGrpcServer(ctx, conf, db, vDb, ml); err != nil {
			zlog.Error("From GRPC server :", err.Error())
			osSignalCh <- syscall.SIGTERM
		} else {
			zlog.Info("Rest gateway disabled")
		}
	}()

	signal := <-osSignalCh

	switch signal {
	case syscall.SIGINT:
		zlog.Error("interrupt signal recived.")
	case syscall.SIGTERM:
		zlog.Error("termination signal signal recived.")
	default:
		zlog.Error("Unknown signal recived.")
	}
	CleanUp(ctx, zlog, db, vDb)
}

func CleanUp(ctx context.Context, log interfaces.LogInterface, db interfaces.DbInterface, vDb interfaces.ValidationDbInterface) {
	grpc.StopGrpcServer()
	log.Info("GRPC server stopped.")

	err := db.CloseDB()
	if err != nil {
		log.Error("Closing database :", err.Error())
	} else {
		log.Info("Database connection closed.")
	}

	err = vDb.CloseDB()
	if err != nil {
		log.Error("Closing validation database :", err.Error())
	} else {
		log.Info("validation Database connection closed.")
	}
}

func getConfFile() (file string, initConf bool) {

	initflag := flag.Bool("initConf", false, "Initialise a conf file if not exist")

	confFile := flag.String("config", DEV_CONFIG, "Path to the configuration file")

	flag.Parse()

	if *confFile == "" {
		log.Fatal().Msg("Please provide a configuration file using -config flag")
		os.Exit(1)
	}
	return *confFile, *initflag
}
