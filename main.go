package main

import (
	"fmt"
	"net/http"

	"github.com/goproxy/goproxy"
)

func main() {
	fmt.Println("binding to 0.0.0.0:8080")
	http.ListenAndServe("0.0.0.0:8080", goproxy.New())
}
