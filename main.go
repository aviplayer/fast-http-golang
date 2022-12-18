package main

import (
	"fmt"
	"github.com/valyala/fasthttp"
)

func main() {
	fmt.Println("Starting server on ", "8089")
	err := fasthttp.ListenAndServe(":8089", Handle)
	if err != nil {
		panic(err)
	}
}
