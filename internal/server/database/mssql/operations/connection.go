package operations

import (
	"context"
	"database/sql"
	"fmt"

	"osh.com/rps/registrar/config"
	"osh.com/rps/registrar/internal/interfaces"
)

type DbOps struct {
	DBconn *sql.DB
	conf   config.System
	log    interfaces.LogInterface
}

func GetNewDbOperationObject(conf config.System) DbOps {
	return DbOps{
		conf: conf,
	}
}

func ConnectDB(ctx context.Context, dBconf config.SqlDBconf, SystemConf config.System, log interfaces.LogInterface) (db DbOps) {

	err := config.ValidateConfig(dBconf)
	if err != nil {
		log.Fatal("DB config validation failed : ", err.Error())
	}

	err = config.ValidateConfig(SystemConf)
	if err != nil {
		log.Fatal("System config validation failed : ", err.Error())
	}

	db = GetNewDbOperationObject(SystemConf)

	dbHost := dBconf.ServerIp
	dbPort := dBconf.Port
	dbUser := dBconf.User
	dbPassword := dBconf.Password
	DBName := dBconf.DBName

	connString := fmt.Sprintf("server=%s,%s;user id=%s;password=%s;database=%s", dbHost, dbPort, dbUser, dbPassword, DBName)
	// Open the database
	DBconn, err := sql.Open("sqlserver", connString)
	if err != nil {
		log.Fatal("Error opening database: %s", err.Error())
	}

	// Test the database connection
	err = DBconn.Ping()
	if err != nil {
		log.Fatal("Error connecting to database: %s", err.Error())
	}

	db.DBconn = DBconn
	db.log = log

	return
}
