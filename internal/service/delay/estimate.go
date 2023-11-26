package delay

import (
	"encoding/json"
	"github.com/ashkan-jafarzadeh/delay/pkg/faker"
	"net/http"
)

type EstimateResponse struct {
	Data struct {
		ETA int `json:"eta"`
	} `json:"data"`
}

func findEstimate(url string) (int, error) {
	return faker.RandInt(5, 25), nil

	// TODO: The provided API not working so I had to bypass it and create my own mock
	resp, err := http.Get(url)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	var response EstimateResponse
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		return 0, err
	}

	return response.Data.ETA, nil
}
