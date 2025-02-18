package database

import (
	"context"

	"osh.com/rps/registrar/config"
	"osh.com/rps/registrar/internal/interfaces"
	"osh.com/rps/registrar/internal/server/database/mongo"
	"osh.com/rps/registrar/internal/server/database/mssql"
)

func GetDb(ctx context.Context, conf interfaces.ConfigInterface, mlog interfaces.MultyLogInterface) (db interfaces.DbInterface, vDb interfaces.ValidationDbInterface) {

	log := mlog.GetLogger()
	config := new(config.DatabaseConfig)
	config.LoadDbConfig(conf, log)

	switch config.DBSystem {
	case "mongo":
		log.Info("Starting with MongoDB system")
		db = mongo.GetNewDbConnection(ctx, conf, mlog)
		vDb = mssql.GetNewDbConnection(ctx, conf, mlog)
	case "mssql":
		log.Info("Starting with Mssql system")
		sqlCons := mssql.GetNewDbConnection(ctx, conf, mlog)
		db = sqlCons
		vDb = sqlCons
	default:
		log.Info("Starting with defualt database system")
		db = mongo.GetNewDbConnection(ctx, conf, mlog)
		vDb = mssql.GetNewDbConnection(ctx, conf, mlog)

	}
	return
}
