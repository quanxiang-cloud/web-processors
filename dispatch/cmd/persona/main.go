package main

import (
	"flag"
	"fmt"
)

var url = flag.String("url", "", "url")

func main() {
	flag.Parse()

	fmt.Println(*url)
}
