package main

import (
	"flag"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	syncQuit = make(chan struct{})
)

var httpAddrFlag = flag.String("http_addr", ":80", "Address to which HTTP server should bind")
var httpsAddrFlag = flag.String("https_addr", ":443", "Address to which HTTPS server should bind")
var httpsCertFlag = flag.String("https_cert", "", "Path to file containing HTTPS certificate")
var httpsKeyFlag = flag.String("https_key", "", "Path to file containing HTTPS key")
var dbTypeFlag = flag.String("db_type", "mysql", "Database type to use for storage")
var dbDsnFlag = flag.String("db_dsn", "root:root@tcp(localhost:3306)/contactee", "DSN to use for connecting to database")

func main() {
	flag.Parse()

	db := mustDatabase(*dbTypeFlag, *dbDsnFlag)
	db.Migrate()

	repositoryManager := NewRepositoryManager(db)

	api := InitApiCore(repositoryManager)
	api.SetHttp(*httpAddrFlag)
	api.SetHttps(*httpsAddrFlag, *httpsCertFlag, *httpsKeyFlag)
	api.StartListen()

	<-syncQuit
}

func mustDatabase(dbType, dbDsn string) *DB {
	db, err := DbInstanceOptional(dbType, dbDsn)
	if err != nil {
		log.Fatalf("Could not connect to database. %v", err)
	}

	return db
}
