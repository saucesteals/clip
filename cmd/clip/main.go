package main

import (
	"flag"
	"log"
)

var (
	op string
)

func init() {
	log.SetFlags(0)
	flag.Parse()
	op = flag.Arg(0)
}

func main() {

	switch op {
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
