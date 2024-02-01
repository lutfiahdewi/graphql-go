package mssql

import (
	"database/sql"
	_ "github.com/microsoft/go-mssqldb"
	/*"github.com/golang-migrate/migrate"
	"github.com/golang-migrate/migrate/database/mysql"
	_ "github.com/golang-migrate/migrate/source/file"*/
	"log"
	"fmt"
	"context"
)
 
var db *sql.DB

func OpenDB(connString string, setLimits bool) (*sql.DB, error) {
	var err error
	db, err := sql.Open("sqlserver", connString)
	if err != nil {
		log.Fatal("Error creating connection, message: ", err.Error())
		// log.Panic(err)
	}

	if setLimits {
		fmt.Println("setting limits")
		db.SetMaxIdleConns(5)
		db.SetMaxOpenConns(5)
	}

	ctx := context.Background()
	// err = db.Ping()
	err = db.PingContext(ctx)
	if err != nil {
		log.Fatal(err.Error())
		// log.Panic(err)
	}
	fmt.Printf("Connected to SQL Server!\n")
	return db, nil
}

func CloseDB() error {
	return db.Close()
}

func Migrate() {
	fmt.Println("This fun is not implemented yet")

}