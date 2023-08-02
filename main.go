package main

import (
  "fmt"
  "net/http"
  "os"
	"strconv"
	"sync"
)

func main() {
  http.HandleFunc("/test", func(w http.ResponseWriter, r *http.Request) {
		test()
  })
  http.ListenAndServe(":8080", nil)
}

func test(){
	var wg sync.WaitGroup
	for i := 0; i < 4000; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			f, err := os.Create("/tmp/file-" + strconv.Itoa(i))
			if err != nil {
				fmt.Println(err)
			}
			_, _ = f.WriteString("writes\n")
		}(i)
	}
}
