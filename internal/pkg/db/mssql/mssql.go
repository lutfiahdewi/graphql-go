package mssql

import (
	"database/sql"
	_ "github.com/microsoft/go-mssqldb"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/sqlserver"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	
	"log"
	"fmt"
	"context"
)
 
var Db *sql.DB

func OpenDB(connString string, setLimits bool) (*sql.DB, error) {
	var err error
	db, err := sql.Open("sqlserver", connString)
	if err != nil {
		log.Fatal("Error creating connection, message: ", err.Error())
		return nil, err
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
		return nil, err
		// log.Panic(err)
	}
	fmt.Printf("Connected to SQL Server!\n")
	Db = db;
	return db, nil
}

func CloseDB() error {
	return Db.Close()
}

func Migrate() {
	//the cli doesnt work :/
	// fmt.Println("This func is not implemented yet")
	if err := Db.Ping(); err != nil {
		log.Fatal(err)
	}
	driver, err := sqlserver.WithInstance(Db, &sqlserver.Config{})
	if err !=nil {
		log.Fatal(err)
	}
	m, err := migrate.NewWithDatabaseInstance(
		"file://internal/pkg/db/migrations/mssql",
		// "file:///C:/Users/ASUS/Documents/coding etc/graphql-go/internal/pkg/db/migrations/mssql",
		"sqlserver",
		driver,
	)
	if err !=nil {
		log.Fatal(err)
	}
	// m.Force(4)// Use this for versioning
	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatal(err)
	}
}