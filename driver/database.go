package driver

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

// DB ...
type DB struct {
	SQL *sql.DB
	// Mgo *mgo.database
}

// DBConn ...
var dbConn = &DB{}

// ConnectSQL ...
func ConnectSQL(host, port, uname, pass, dbname string) (*DB, error) {
	dbSource := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?sslmode=disable",
		uname,
		pass,
		host,
		port,
		dbname,
	)
	d, err := sql.Open("postgres", dbSource)
	if err != nil {
		panic(err)
	}
	dbConn.SQL = d
	return dbConn, err
}

func ConexaoPostgres(host,port,user,password,dbname string) (*DB, error) {
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	dbConn.SQL = db
	return dbConn, err
}
