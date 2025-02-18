package operations

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"osh.com/rps/registrar/config"
	"osh.com/rps/registrar/internal/interfaces"
)

var osh_db = "osh"

type dbDetails struct {
	dbName     string
	collection string
}

type DbOps struct {
	DBconn      *mongo.Client
	RegistrarDB dbDetails
	conf        config.System
	log         interfaces.LogInterface
}

func GetNewDbOperationObject(conf config.System) DbOps {
	return DbOps{
		conf: conf,
	}
}

func ConnectDB(ctx context.Context, Conf config.MongoDBconf, systemConf config.System, log interfaces.LogInterface) (db DbOps) {

	err := config.ValidateConfig(Conf)
	if err != nil {
		log.Fatal("DB config validation failed : ", err.Error())
	}

	err = config.ValidateConfig(systemConf)
	if err != nil {
		log.Fatal("System config validation failed : ", err.Error())
	}

	var mongoDbName, mongoHost string

	db = GetNewDbOperationObject(systemConf)

	mongoHost = fmt.Sprintf("mongodb://%s", Conf.Host)

	clientOptions := options.Client().ApplyURI(mongoHost).
		SetMaxPoolSize(Conf.MaxConn).
		SetMaxConnIdleTime(30 * time.Second)

	conn, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err.Error())
		return
	}
	mongoDbName = osh_db + Conf.Name
	//TODO: CONFIGURE DATABASE
	db.DBconn = conn
	db.RegistrarDB.dbName = mongoDbName
	db.RegistrarDB.collection = Conf.Collection
	db.log = log
	return
}
