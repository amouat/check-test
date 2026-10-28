package main

import (
	"flag"
	"fmt"

	"github.com/amouat/check-test/greet"
)

func main() {
	name := flag.String("name", "", "name to greet")
	flag.Parse()
	fmt.Println(greet.Hello(*name))
}
