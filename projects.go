package timev2api

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// To save project data
type ProjectsReturn struct {
	Projects []ProjectsDataReturn
}

type ProjectsDataReturn struct {
	Name string `json:"name"`
	ID   string `json:"id"`
}

// To get a list of all projects
func Projects(domainkey, token string) (ProjectsReturn, error) {

	// Define new client
	client := &http.Client{}

	// New http request
	request, err := http.NewRequest("GET", fmt.Sprintf("https://%s.timev2.de/api/v1/projects", domainkey), nil)
	if err != nil {
		return ProjectsReturn{}, err
	}

	// Set header
	request.Header.Set("authorization", token)

	// Response to timev2
	response, err := client.Do(request)
	if err != nil {
		return ProjectsReturn{}, err
	}

	// Save projects
	var decode ProjectsReturn

	// Decode json code in struct
	err = json.NewDecoder(response.Body).Decode(&decode)
	if err != nil {
		return ProjectsReturn{}, err
	}

	// Return
	return decode, nil

}
