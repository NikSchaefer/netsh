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

	createFileProfile(name, password)

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
	// netsh wlan connect name=sideswipe
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
	file := fmt.Sprintf("filename=%s\\profile.xml", dir)

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
