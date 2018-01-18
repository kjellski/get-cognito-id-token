package main

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws/session"
)

func main() {
	sess, err := session.NewSessionWithOptions(session.Options{
		Profile: "dev",
	})

	if err != nil {
		fmt.Printf("Error from session: %v\n", err)
		return
	}

	fmt.Printf("session: %+v", sess)
}
