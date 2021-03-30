package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"strings"
)

func main() {

	name := "sideswipe"
	password := "myPassword"

	createConnectDelete(name, password)
}

func createConnectDelete(name, pass string) error {
	createFileProfile(name, pass)
	addProfile(name)
	connectToNetwork(name)
	deleteProfile(name)

	return nil
}

func createFileProfile(name, pass string) error {
	filename := fmt.Sprintf("_%s.xml", name)
	file, err := os.Create(filename)
	must(err)

	data, err := ioutil.ReadFile("profile.xml")
	must(err)

	str := string(data)
	str = strings.Replace(str, "{password}", pass, 1)
	str = strings.Replace(str, "{SSID}", name, 2)

	file.Write([]byte(str))
	return nil
}

func connectToNetwork(name string) bool {
	str := fmt.Sprintf("name=%s", name)
	out, err := exec.Command("netsh", "wlan", "connect", str).Output()
	must(err)
	fmt.Print(out)
	return false
}

func deleteProfile(name string) error {
	out, err := exec.Command(`netsh`, `wlan`, `delete`, `profile`, name).Output()
	must(err)
	fmt.Print(string(out))
	return nil
}

func addProfile(name string) error {
	dir, err := os.Getwd()
	must(err)
	file := fmt.Sprintf("filename=%s\\_%s.xml", dir, name)
	out, err := exec.Command(`netsh`, `wlan`, `add`, `profile`, file, `interface="WI-FI"`, `user=current`).Output()
	must(err)
	fmt.Print(string(out))
	return nil
}

func must(err error) {
	if err != nil {
		log.Fatal(err, "ERR")
	}
}
