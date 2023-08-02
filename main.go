package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"os"
	"strings"
	"syscall"
	"time"
)

var zoneDirs = []string{
	// Update path according to your OS
	"/usr/share/zoneinfo",
	"/usr/share/lib/zoneinfo",
	"/usr/lib/locale/TZ",
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
	fmt.Fprintf(w, "Hello, you were given id %d\n", id)

	tmpFolder, err := os.MkdirTemp("/tmp/", "test")

	if err != nil {
		panic("An error occured")
	}
	for i := range allTZ {
		filename := fmt.Sprintf("%s/session_data_%s.txt", tmpFolder, allTZ[i])
		flags := syscall.O_WRONLY | syscall.O_CREAT | syscall.O_TRUNC
		mode := uint32(0644)
		fd, err := syscall.Open(filename, flags, mode)
		if err == nil {
			syscall.Write(fd, []byte(time.Now().String()))
		}
	}
}

func main() {
	log.SetOutput(ioutil.Discard)
	http.HandleFunc("/hello", hello)
	fmt.Printf("Starting server!\n")
	http.ListenAndServe("0.0.0.0:8080", nil)
}
