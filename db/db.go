// Copyright (c) 2025 sakib-maho
// Licensed under the MIT License
// See LICENSE file for details

package db

import (
	"database/sql"
	"fmt"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "1234"
	dbname   = "currency"
)

func Connection() *sql.DB {

	postgresConfiguration := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	Db, err := sql.Open("postgres", postgresConfiguration)

	err = Db.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Println("Successfully connected!")

	return Db

}
