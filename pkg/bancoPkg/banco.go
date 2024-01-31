package bancopkg

import (
	"database/sql"
	"fmt"
	"log"
)

type Database struct {
	conn *sql.DB
}

func NewDatabase(dataSourceName string) (*Database, error) {
	conn, err := sql.Open("mysql", dataSourceName)

	if err != nil {
		return nil, err
	}

	return &Database{conn: conn}, nil
}

func (db *Database) Close() {
	err := db.conn.Close()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Conexão encerrada com sucesso!")
}
