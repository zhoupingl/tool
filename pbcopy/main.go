package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
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

const port = ":9097"

func main() {
	router := gin.Default()

	// cache system
	router.GET("/cache", func(context *gin.Context) {
		for key, values := range context.Request.URL.Query() {
			if !(values[0] == "") {
				set(key, values[0])
			}
			context.String(200, get(key))
			return
		}
	})
	router.StaticFile("/favicon.ico", "./favicon.png")

	// test wxpay return url
	router.Any("/wxpay", func(context *gin.Context) {
		fmt.Println("query", context.Request.URL.Query().Encode())
		buf, _ := ioutil.ReadAll(context.Request.Body)
		fmt.Println("body", string(buf))
	})

	// test alipay return url
	router.Any("/alipay", func(context *gin.Context) {
		fmt.Println("query", context.Request.URL.Query().Encode())
		buf, _ := ioutil.ReadAll(context.Request.Body)
		fmt.Println("body", string(buf))
	})


	fmt.Println("Listen port" + port)
	router.Run(port)

}
