package timev2api

import "encoding/json"

// To save user data
type UsersReturn struct {
	Total int              `json:"total"`
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
func Users(domainkey, token string) ([]UserDataReturn, error) {

	// Response to timev2
	response, err := response("users", domainkey, token)
	if err != nil {
		return nil, err
	}

	// Save clients
	var decode UsersReturn

	// Decode json code in struct
	err = json.NewDecoder(response).Decode(&decode)
	if err != nil {
		return nil, err
	}

	// Return data
	return decode.Users, nil

}
