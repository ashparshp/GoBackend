package auth

import "fmt"

func LoginWithCredentials(username, password string) {
	fmt.Printf("Login.. Username: %s Password: %s", username, password)
}