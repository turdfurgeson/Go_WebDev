package main

import (
	"fmt"
	"net/http"
)

type batman string

func (b batman) ServeHTTP(a http.ResponseWriter, c *http.Request) {
	fmt.Fprintln(a, "Holy fuckablls batman. I am printing to my browser through my own router!")
}

func main() {
	var x batman

	http.ListenAndServe(":8080", x)
}
