package timev2api

import "encoding/json"

// To save project data
type ProjectsReturn struct {
	Total    int `json:"total"`
	Projects []ProjectsDataReturn
}

type ProjectsDataReturn struct {
	Name string `json:"name"`
	ID   string `json:"id"`
}

// To get a list of all projects
func Projects(domainkey, token string) (ProjectsReturn, error) {

	// Response to timev2
	response, err := response("projects", domainkey, token)
	if err != nil {
		return ProjectsReturn{}, err
	}

	// Save clients
	var decode ProjectsReturn

	// Decode json code in struct
	err = json.NewDecoder(response).Decode(&decode)
	if err != nil {
		return ProjectsReturn{}, err
	}

	// Return data
	return decode, nil

}
