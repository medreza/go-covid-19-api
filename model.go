package main

// DataResponse is the model for country/global cases data
type DataResponse struct {
	Ok   bool `json:"ok"`
	Data struct {
		Confirmed    int     `json:"confirmed"`
		Recovered    int     `json:"recovered"`
		Deaths       int     `json:"deaths"`
		Active       int     `json:"active"`
		RecoveryRate float32 `json:"recoveryRate"`
		FatalityRate float32 `json:"fatalityRate"`
	} `json:"data"`
}

func (r *DataResponse) activeCases() int {
	return r.Data.Confirmed - r.Data.Recovered - r.Data.Deaths
}

func (r *DataResponse) recoveryRate() float32 {
	return float32(r.Data.Recovered) / float32(r.Data.Confirmed)
}

func (r *DataResponse) fatalityRate() float32 {
	return float32(r.Data.Deaths) / float32(r.Data.Confirmed)
}

// ErrorResponse loads error message from queries. Always return "ok" : false.
type ErrorResponse struct {
	Ok    bool `json:"ok"`
	Error struct {
		Message string `json:"message"`
	} `json:"error"`
}
