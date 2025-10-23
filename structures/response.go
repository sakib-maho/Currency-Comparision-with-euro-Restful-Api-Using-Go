// Copyright (c) 2025 sakib-maho
// Licensed under the MIT License
// See LICENSE file for details

package structures

type Response struct {
	Base  string             `json:"base"`
	Rates map[string]float64 `json:"rates"`
}

type Rates_analyze map[string]map[string]float64

type AnalyzerResponse struct {
	Base          string                        `json:"base"`
	Rates_analyze map[string]map[string]float64 `json:"rates_analyze"`
}
