package main

import (
	"io/ioutil"
	"log"
	"os"

	"github.com/saucesteals/clip/clipboard"
)

func copy() {
	stdin, err := ioutil.ReadAll(os.Stdin)

	if err != nil {
		log.Fatal(err)
	}

	err = clipboard.Write(stdin)

	if err != nil {
		log.Fatal(err)
	}
}
