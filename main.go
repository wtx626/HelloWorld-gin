/**
 * @Author: tianxiongwu
 * @Date: 2020/8/5 4:51 下午
 */
package main

import (
	"fmt"
	//"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	//r := gin.Default()
	//r.GET("/ping", func(c *gin.Context) {
	//	c.JSON(200, gin.H{
	//		"message": "pong",
	//	})
	//})
	//r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
	http.HandleFunc("/", HelloServer)
	http.ListenAndServe(":8080", nil)
}

func HelloServer(w http.ResponseWriter, r *http.Request) {
	//fmt.Fprintf(w, "Hello, %s!", r.URL.Path[1:])
	fmt.Fprintf(w, "Hello, world")
}
