package main

import (
	"fmt"

	"github.com/ashparshp/podcast/auth"
)

func main() {

	username := "ashparsh"
	password := "ashparsh@1"
	auth.LoginWithCredentials(username,password)
	session := auth.Session()
	fmt.Println(session)
}