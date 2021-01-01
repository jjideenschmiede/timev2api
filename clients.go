package timev2api

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// To save client data
type ClientsReturn struct {
	Clients []ClientDataReturn
}

type ClientDataReturn struct {
	Name string `json:"name"`
	ID   string `json:"id"`
	Rate string `json:"rate"`
}

// Get a list of all clients
func Clients(domainkey, token string) (ClientsReturn, error) {

	// Define client
	client := &http.Client{}

	// New http request
	request, err := http.NewRequest("GET", fmt.Sprintf("https://%s.timev2.de/api/v1/clients", domainkey), nil)
	if err != nil {
		return ClientsReturn{}, err
	}

	// Set header
	request.Header.Set("authorization", token)

	// Response to timev2
	response, err := client.Do(request)
	if err != nil {
		return ClientsReturn{}, err
	}

	// Save clients
	var decode ClientsReturn

	// Decode json code in struct
	err = json.NewDecoder(response.Body).Decode(&decode)
	if err != nil {
		return ClientsReturn{}, err
	}

	// Return data
	return decode, nil

}
