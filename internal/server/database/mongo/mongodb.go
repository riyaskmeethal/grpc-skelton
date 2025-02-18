package mongo

import (
	"context"

	"osh.com/rps/registrar/config"
	"osh.com/rps/registrar/internal/interfaces"
	"osh.com/rps/registrar/internal/server/database/mongo/operations"
)

type MongoDB struct {
	Connect map[string]operations.DbOps
	conf    interfaces.ConfigInterface
	log     interfaces.MultyLogInterface
}

func GetNewDbConnection(ctx context.Context, conf interfaces.ConfigInterface, mlog interfaces.MultyLogInterface) (mdbs MongoDB) {

	log := mlog.GetLogger()
	config := new(config.MongoDatabaseConfig)
	config.LoadDbConfig(conf, log)

	mdbs = MongoDB{
		conf: conf,
		log:  mlog,
	}

	dbs := make(map[string]operations.DbOps)

	for institute, Conf := range config.MongoDBconf {
		log := mlog.GetInstituteLogger(institute)
		dbs[institute] = operations.ConnectDB(ctx, Conf, config.SystemConf, log)

		log.Info(institute, "MONGO Database connection created successfully.")

	}
	mdbs.Connect = dbs
	return
}

func (db MongoDB) CloseDB() error {
	for _, db := range db.Connect {
		if err := db.DBconn.Disconnect(context.TODO()); err != nil {
			return err
		}
	}
	return nil
}
