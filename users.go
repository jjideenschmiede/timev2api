package timev2api

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// To save user data
type UsersReturn struct {
	Users []UserDataReturn `json:"users"`
}

type UserDataReturn struct {
	Username string `json:"username"`
	Fullname string `json:"fullname"`
	Email    string `json:"email"`
	Role     string `json:"role"`
	Token    string `json:"token"`
}

// Get list of users
func Users(domainkey, token string) (UsersReturn, error) {

	// Define client
	client := &http.Client{}

	// New http request
	request, err := http.NewRequest("GET", fmt.Sprintf("https://%s.timev2.de/api/v1/users", domainkey), nil)
	if err != nil {
		return UsersReturn{}, err
	}

	// Set header
	request.Header.Set("authorization", token)

	// Response ti timev2
	response, err := client.Do(request)
	if err != nil {
		return UsersReturn{}, err
	}

	// Save users
	var decode UsersReturn

	// Decode json conde in struct
	err = json.NewDecoder(response.Body).Decode(&decode)
	if err != nil {
		return UsersReturn{}, err
	}

	// Return data
	return decode, nil

}
