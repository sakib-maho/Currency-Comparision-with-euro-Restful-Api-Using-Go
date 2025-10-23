// Copyright (c) 2025 sakib-maho
// Licensed under the MIT License
// See LICENSE file for details

package structures

type ResponseStruct struct {
	Cube struct {
		Cube []struct {
			Text string `xml:",chardata"`
			Time string `xml:"time,attr"`
			Cube []struct {
				Text     string  `xml:",chardata"`
				Currency string  `xml:"currency,attr"`
				Rate     float64 `xml:"rate,attr"`
			} `xml:"Cube"`
		} `xml:"Cube"`
	} `xml:"Cube"`
}
