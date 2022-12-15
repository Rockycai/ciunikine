package main

import (
	"ciunikine/server/service"
	"fmt"
)

func main() {
	// err := app.Run()

	// if err != nil {
	// 	log.Fatal(err)
	// }

	cmd := "whoami"
	ip := "167.88.184.123"
	port := 6000
	username := "wenlinux"
	password := "cwt2823703"
	privateKey := ""
	passphrase := ""

	result, err := service.ExecCommandBySSH(
		cmd,
		ip,
		port,
		username,
		password,
		privateKey,
		passphrase,
	)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Print(result)
}
