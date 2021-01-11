package timev2api

import "encoding/json"

// To save project data
type ProjectsReturn struct {
	Projects []ProjectsDataReturn
}

type ProjectsDataReturn struct {
	Name string `json:"name"`
	ID   string `json:"id"`
}

// To get a list of all projects
func Projects(domainkey, token string) ([]ProjectsDataReturn, error) {

	// Response to timev2
	response, err := response(domainkey, token)
	if err != nil {
		return nil, err
	}

	// Save clients
	var decode ProjectsReturn

	// Decode json code in struct
	err = json.NewDecoder(response).Decode(&decode)
	if err != nil {
		return nil, err
	}

	// Return data
	return decode.Projects, nil

}
