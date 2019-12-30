package main

import (
	"fmt"
	"net/http"

	"github.com/goproxy/goproxy"
)

func main() {
	fmt.Println("binding to localhost:8080")
	http.ListenAndServe("localhost:8080", goproxy.New())
}
