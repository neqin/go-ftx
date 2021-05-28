package futures

import (
	"fmt"
	"net/http"
	"time"
)

type RequestForStats struct {
	ProductCode string `url:"-"`
}

type ResponseForStats Stats

/*
{"impliedVolatility":1.3356387615203857,"delta":-0.09478103596922183,"gamma":0.00012176595605819894},"openInterest":260.2345}
*/

type StatsList []Stats
type Stats struct {
	Name                     string  `json:"name,omitempty"`
	Volume                   float64 `json:"volume"`
	NextFundingRate          float64 `json:"nextFundingRate"`
	ExpirationPrice          float64 `json:"expirationPrice"`
	PredictedExpirationPrice float64 `json:"predictedExpirationPrice"`
	StrikePrice              float64 `json:"strikePrice"`
	OpenInterest             float64 `json:"openInterest"`
	Greeks                   struct {
		ImpliedVolatility float64 `json:"impliedVolatility"`
		Delta             float64 `json:"delta"`
		Gamma             float64 `json:"gamma"`
		OpenInterest      float64 `json:"openInterest"`
	} `json:"greeks"`
	NextFundingTime time.Time `json:"nextFundingTime"`
}

func (req *RequestForStats) Path() string {
	return fmt.Sprintf("/futures/%s/stats", req.ProductCode)
}

func (req *RequestForStats) Method() string {
	return http.MethodGet
}

func (req *RequestForStats) Query() string {
	return ""
}

func (req *RequestForStats) Payload() []byte {
	return nil
}

func (a StatsList) Len() int      { return len(a) }
func (a StatsList) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a StatsList) Less(i, j int) bool {
	return a[i].NextFundingRate < a[j].NextFundingRate
}
