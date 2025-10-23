// Copyright (c) 2025 sakib-maho
// Licensed under the MIT License
// See LICENSE file for details

package main

import (
	"Currency-comaprasion/db"
	"Currency-comaprasion/models"
	"Currency-comaprasion/routers"
	_ "github.com/lib/pq"
	"net/http"
)

func main() {
	models.InsertCurrencyModels(db.Connection())
	mux := routers.Routers()
	s := http.Server{
		Addr:    ":3000",
		Handler: mux,
	}
	s.ListenAndServe()
}
