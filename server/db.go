package server

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

const dbms = "mysql"
const mysqlProtocol = "tcp"

var dbConfig DbConfig

func SetDbConfig(config DbConfig) {
	dbConfig = config
}

func GetDbConfig() DbConfig {
	return dbConfig
}

func GetDsnString(config DbConfig) (dsn string) {
	dsn = fmt.Sprintf(
		"%v:%v@%v(%v:%v)/%v",
		config.User,
		config.Pass,
		mysqlProtocol,
		config.Host,
		config.Port,
		config.DbName,
	)
	return dsn
}

func DbConnect() (db *sql.DB, err error) {
	db, _ = sql.Open(dbms, GetDsnString(dbConfig))

	// try to ping the database to check if it is connected
	err = db.Ping()

	return db, err
}
