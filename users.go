package timev2api

import "encoding/json"

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

	// Response to timev2
	response, err := response(domainkey, token)
	if err != nil {
		return UsersReturn{}, err
	}

	// Save clients
	var decode UsersReturn

	// Decode json code in struct
	err = json.NewDecoder(response).Decode(&decode)
	if err != nil {
		return UsersReturn{}, err
	}

	// Return data
	return decode, nil

}
