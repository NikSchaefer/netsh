package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
)

func main() {

	// name := "sideswipe"
	// password := ""

	// passed := loginToNetwork(name, password)
	// fmt.Println(passed)

	addProfile("J")
}

func loginToNetwork(name, password string) bool {
	// netsh wlan connect name=sideswipe
	return false
}

func deleteProfile(name string) error {
	return nil
}

func addProfile(name string) error {
	dir, err := os.Getwd()
	must(err)
	file := fmt.Sprintf("filename=%s\\profile.xml", dir)

	out, err := exec.Command(`netsh`, `wlan`, `add`, `profile`, file, `interface="WI-FI"`, `user=current`).Output()
	must(err)
	fmt.Println(string(out))
	return nil
}

func must(err error) {
	if err != nil {
		log.Fatal(err, "ERR")
	}
}
