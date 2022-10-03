package main

import (
	"log"
	"os"

	"github.com/nadahusin/rental/src/config"
)

func main() {
	if err := config.Run(os.Args[1:]); err != nil {
		log.Fatal(err)
	}

}
