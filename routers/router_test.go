// Copyright (c) 2025 sakib-maho
// Licensed under the MIT License
// See LICENSE file for details

package routers_test

import (
	_ "Currency-comaprasion/testing_init"
	"fmt"
	. "github.com/smartystreets/goconvey/convey"
	"io/ioutil"
	"log"
	"net/http"
	"testing"
)

func TestRouters(t *testing.T) {
	/* /rates/latest testing */
	var (
		router string
	)
	router = "/rates/latest"
	latestResponse := getResponse(router)
	testResult(t, latestResponse, router)

	/* /rates/{date} testing */
	date := "2022-06-15"
	router = "/rates/{date}"
	dateResponse := getResponse("/rates/" + date)
	testResult(t, dateResponse, router)

	/* /rates/analyze testing */
	router = "/rates/analyze"
	analyzeResponse := getResponse("/rates/analyze")
	testResult(t, analyzeResponse, router)
}

func getResponse(url string) *http.Response {
	URI := "http://localhost:3000" + url
	response, err := http.Get(URI)
	if err != nil {
		panic(err)
	}
	return response
}
func testResult(t *testing.T, response *http.Response, router string) {
	Convey(fmt.Sprintf("%s", "Subject: Testing "+router+"  endpoint"), t, func() {
		Convey("Response Should be 200", func() {
			So(response.StatusCode, ShouldEqual, 200)
			Convey("Body should not be nil", func() {
				body, err := ioutil.ReadAll(response.Body)
				if err != nil {
					log.Fatalln(err)
				}
				So(string(body), ShouldNotBeEmpty)
			})
		})
	})
}
