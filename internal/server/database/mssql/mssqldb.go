package mssql

import (
	"context"

	_ "github.com/denisenkom/go-mssqldb"
	"osh.com/rps/registrar/config"
	"osh.com/rps/registrar/internal/interfaces"
	"osh.com/rps/registrar/internal/server/database/mssql/operations"
)

type MsSqlDB struct {
	DbOps map[string]operations.DbOps
	conf  interfaces.ConfigInterface
	log   interfaces.LogInterface
}

func GetNewDbConnection(ctx context.Context, conf interfaces.ConfigInterface, mlog interfaces.MultyLogInterface) (spqDbs MsSqlDB) {

	log := mlog.GetLogger()
	config := new(config.SqlDatabaseConfig)
	config.LoadDbConfig(conf, log)

	spqDbs = MsSqlDB{
		conf: conf,
		log:  log,
	}

	dbs := make(map[string]operations.DbOps)

	for institute, Conf := range config.SqlDBconf {
		log := mlog.GetInstituteLogger(institute)
		dbs[institute] = operations.ConnectDB(ctx, Conf, config.SystemConf, log)

		log.Info(institute, "SQL Database connection created successfully.")

	}
	spqDbs.DbOps = dbs
	return
}

func (dbs MsSqlDB) CloseDB() (err error) {
	for _, db := range dbs.DbOps {
		err = db.DBconn.Close()
		if err != nil {
			return
		}
	}
	return
}
