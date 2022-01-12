package main

import (
	"log"
	"os"

	"github.com/saucesteals/clip/clipboard"
)

func paste() {
	buffer, err := clipboard.Read()

	if err != nil {
		log.Fatal(err)
	}

	_, err = os.Stdout.Write(buffer)

	if err != nil {
		log.Fatal(err)
	}
}
