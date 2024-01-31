package database

import (
	"database/sql"
	"fmt"
	"log"
	"sync"
)

var (
	db   *sql.DB
	once sync.Once
)

func Init(dataSourceName string) (*sql.DB, error) {
	var err error

	once.Do(func() {
		db, err = sql.Open("mysql", dataSourceName)
		if err != nil {
			log.Fatal(err)
		}

		err = db.Ping()
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println("Conexão estabelecida com sucesso!")
	})

	return db, err
}

func Close() {
	if db != nil {
		db.Close()
		fmt.Println("Conexão encerrada com sucesso!")
	}
}
