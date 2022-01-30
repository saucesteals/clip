package main

import (
	"flag"
	"log"
)

func main() {
	log.SetFlags(0)
	flag.Parse()

	switch flag.Arg(0) {
	case "paste":
		paste()
	case "copy":
		copy()
	case "post":
		post()
	default:
		log.Fatal("Usage: clip <copy | paste | post>")
	}

}
