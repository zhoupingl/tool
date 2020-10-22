package main

import (
	"fmt"
	"net/http"
	"sync"
)

var data = map[string]string{}
var lock sync.Mutex

func get(key string) string {
	lock.Lock()
	defer lock.Unlock()
	return data[key]
}

func set(key, val string) {
	lock.Lock()
	defer lock.Unlock()
	data[key] = val
}

func main() {
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		for k, values := range request.URL.Query() {
			for _, v := range values {
				if v == "" {
					writer.Write([]byte(get(k)))
					return
				} else {
					set(k, v)
					writer.Write([]byte(v))
				}
			}
		}
	})

	const port = ":9097"
	fmt.Println("Listen port:" + port)
	http.ListenAndServe(port, nil)
}
