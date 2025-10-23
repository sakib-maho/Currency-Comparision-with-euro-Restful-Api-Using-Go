// Copyright (c) 2025 sakib-maho
// Licensed under the MIT License
// See LICENSE file for details

package controllers

import (
	"Currency-comaprasion/db"
	"Currency-comaprasion/structures"
	"encoding/json"
	"log"
	"net/http"
	time2 "time"
)

func Latest(w http.ResponseWriter, r *http.Request) {

	var (
		ratedate     string
		currencyname string
		currencyrate float64
		Response     structures.Response
		m            = make(map[string]float64)
	)

	db := db.Connection()
	time := time2.Now().AddDate(0, 0, -2).Format("2006-01-02")

	sqlStatement := `SELECT * FROM currencyrate WHERE ratedate=$1`
	result, err := db.Query(sqlStatement, time)
	if err != nil {
		log.Fatal(err)
	}
	for result.Next() {
		result.Scan(&ratedate, &currencyname, &currencyrate)
		m[currencyname] = currencyrate
	}
	Response = structures.Response{
		Base:  "EUR",
		Rates: m,
	}
	q, _ := json.Marshal(Response)
	w.Write(q)
}
