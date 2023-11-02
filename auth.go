package ioriver

import (
	"encoding/json"
	"fmt"
	"io"
)

type Auth struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type AuthToken struct {
	Token string `json:"auth_token"`
}

// func (client *IORiverClient) register(self, username, email, password) {
// 	options = {"username": username, "password": password,
// 			   "email": email, "firstName": "default", "lastName": "default"}
// 	return self.client.call_api("accounts/", "POST", payload=options, monitor_async=False)
// }

// func (client *IORiverClient) login_jwt(self, username, password) {
// 	options = {"username": username, "password": password}
// 	return self.client.call_api("auth/jwt/create/", "POST", payload=options, monitor_async=False)
// }

func (client *IORiverClient) LoginToken(username string, password string) (string, error) {
	data := Auth{username, password}
	params := CallParams{payload: data}
	resp, err := client.CallApi("auth/token/login/", "POST", params)
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("Failed to read response: %w", err)
	}

	var result AuthToken
	err = json.Unmarshal(respBody, &result)
	if err != nil {
		return "", fmt.Errorf("Error unmarshaling json: %w", err)
	}

	return result.Token, nil
}
