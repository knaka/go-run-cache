package main

import (
	"github.com/knaka/gobin/lib"
	"log"
	"os"
)

func main() {
	err := lib.Run(os.Args[1:])
	if err != nil {
		log.Fatalf("%v", err)
	}
}
