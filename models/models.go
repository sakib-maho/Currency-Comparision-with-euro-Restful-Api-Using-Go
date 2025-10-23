// Copyright (c) 2025 sakib-maho
// Licensed under the MIT License
// See LICENSE file for details

package models

import (
	"Currency-comaprasion/structures"
	"database/sql"
	"encoding/xml"
	"fmt"
	"io"
	"log"
	"net/http"
)

func InsertCurrencyModels(db *sql.DB) {
	var (
		date           string
		currencyName   string
		rate           float64
		error          error
		ResponseObject structures.ResponseStruct
	)
	response, err := http.Get("https://www.ecb.europa.eu/stats/eurofxref/eurofxref-hist-90d.xml")
	if err != nil {
		log.Fatal(err)
	}
	body, responseErr := io.ReadAll(response.Body)
	if responseErr != nil {
		log.Fatal(responseErr)
	}

	unMarshalError := xml.Unmarshal(body, &ResponseObject)
	if err != nil {
		log.Fatal(unMarshalError)
	}

	for _, cubes := range ResponseObject.Cube.Cube {
		date = cubes.Time
		for _, cube := range cubes.Cube {
			currencyName = cube.Currency
			rate = cube.Rate
			sqlStatement := `INSERT INTO currencyrate(ratedate, currencyname,currencyrate)VALUES ( $1, $2, $3)`
			_, error = db.Exec(sqlStatement, date, currencyName, rate)
		}
	}
	if error != nil {
		panic(error)
	} else {
		fmt.Println("Successfully Inserted to database!")
	}
}
