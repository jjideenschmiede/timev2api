package timev2api

import (
	"fmt"
	"io"
	"net/http"
)

// To response the data
func response(key, domainkey, token string) (io.ReadCloser, error) {

	// Define client
	client := &http.Client{}

	// New http request
	request, err := http.NewRequest("GET", fmt.Sprintf("https://%s.timev2.de/api/v1/%s", domainkey, key), nil)
	if err != nil {
		return nil, err
	}

	// Set header
	request.Header.Set("authetication", token)

	// Response to timev2
	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}

	// Return
	return response.Body, nil

}
