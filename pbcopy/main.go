package main

import (
	"fmt"
	"io/ioutil"
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

	// static file dir
	h := http.FileServer(http.Dir("/Users/ann/www/tool/pbcopy"))

	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		for k, values := range request.URL.Query() {
			for _, v := range values {
				if v == "" {
					writer.Write([]byte(get(k)))
					return
				} else {
					set(k, v)
					writer.Write([]byte(v))
					return
				}
			}
		}
	})

	http.HandleFunc("/wechat", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Println("query", request.URL.Query().Encode())
		buf, _ := ioutil.ReadAll(request.Body)
		fmt.Println("body", string(buf))
	})

	const port = ":9097"
	fmt.Println("Listen port:" + port)
	http.ListenAndServe(port, h)
}
