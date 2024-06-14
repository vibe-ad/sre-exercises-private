package main

import (
	"bufio"
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"os"
	"strings"
)

var zoneDirs = []string{
	// Update path according to your OS
	"/usr/share/zoneinfo",
	"/usr/share/lib/zoneinfo",
	"/usr/lib/locale/TZ",
}

func LoadEnvFile(envFilepath string) bool {
	file, err := os.Open(envFilepath)
	if err != nil {
		return false
	}
	defer func() {
		_ = file.Close()
	}()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()

		if strings.HasPrefix(line, "#") {
			continue
		}

		separatorIndex := strings.Index(line, "=")

		if separatorIndex == -1 {
			continue
		}

		key, value := line[0:separatorIndex], line[separatorIndex+1:]

		_ = os.Setenv(key, value)
	}

	return true
}

func Setup() {
	f, err := os.OpenFile("/etc/hosts", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		panic(err)
	}

	defer f.Close()

	if _, err = f.WriteString("127.0.1.2 httpbin.org"); err != nil {
		panic(err)
	}
}

func GetAllTZ() []string {
	result := make([]string, 0, 500)
	for _, zoneDir := range zoneDirs {
		result = append(result, listTZ(zoneDir, "")...)
	}
	return result
}

func listTZ(dir string, path string) []string {
	result := make([]string, 0)
	files, _ := os.ReadDir(dir + path)
	for _, f := range files {
		if f.Name() != strings.ToUpper(f.Name()[:1])+f.Name()[1:] {
			continue
		}
		if f.IsDir() {
			result = append(result, listTZ(dir, path+"/"+f.Name())...)
		} else {
			tz := path + "/" + f.Name()
			if tz[0] == '/' {
				tz = tz[1:]
			}
			result = append(result, tz)
		}
	}
	return result
}

var tmpFolder string

var allTZ = GetAllTZ()

func hello(w http.ResponseWriter, req *http.Request) {

	id := rand.Int63()
	tz := allTZ[rand.Intn(len(allTZ))]
	fmt.Fprintf(w, "Hello, you were given id %d and timezone %s\n", id, tz)

	client := http.Client{}
	req, err := http.NewRequest("GET", fmt.Sprintf("http://httpbin.org/anything?id=%d&tz=%s", id, tz), nil)
	if err != nil {
		fmt.Println(err)
	}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	fmt.Fprintf(w, "Server responded with:")
	fmt.Fprintf(w, string(body))
}

func main() {
	LoadEnvFile(".env")
	Setup()
	http.HandleFunc("/hello", hello)
	fmt.Printf("Starting server!\n")
	http.ListenAndServe("0.0.0.0:8080", nil)
}
