package main

import (
	"fmt"
	"os/exec"
)

func main() {
	fmt.Println("HELLO WORLD")

	name := "sideswipe"
	password := ""

	passed := loginToNetwork(name, password)
	fmt.Println(passed)
}

func loginToNetwork(name, password string) bool {
	// netsh wlan connect name=sideswipe
	return false
}

func deleteProfile(name string) error {
	return nil
}

func addProfile(name string) error {
	exec.Command(`netsh wlan add profile filename="C:\Users\schae\Desktop\Wi-Fi-sideswipe.xml" interface="WI-FI" user=current`)
	return nil
}
