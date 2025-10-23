// Copyright (c) 2025 sakib-maho
// Licensed under the MIT License
// See LICENSE file for details

package controllers

import (
	"Currency-comaprasion/db"
	"Currency-comaprasion/structures"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strings"
)

func DateRate(w http.ResponseWriter, r *http.Request) {
	var (
		ratedate     string
		currencyname string
		currencyrate float64
		currencyRate map[string]float64
		Response     structures.Response
		m            = make(map[string]float64)
	)
	vars := mux.Vars(r)
	date := vars["date"]
	bool := strings.Contains(date, "-")

	if bool == false {
		fmt.Println(bool)
		w.Write([]byte("Not a valid param"))

	} else {
		db := db.Connection()

		sqlStatement := `SELECT * FROM currencyrate WHERE ratedate=$1`
		result, err := db.Query(sqlStatement, date)
		if err != nil {
			log.Fatal(err)
		}
		for result.Next() {
			result.Scan(&ratedate, &currencyname, &currencyrate)
			currencyRate = map[string]float64{currencyname: currencyrate}
			m[currencyname] = currencyrate
		}
		t, _ := json.Marshal(currencyRate)
		fmt.Println(string(t))
		Response = structures.Response{
			Base:  "EUR",
			Rates: m,
		}
		q, _ := json.Marshal(Response)
		w.Write(q)

	}

}
