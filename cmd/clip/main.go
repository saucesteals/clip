package main

import (
	"flag"
	"log"
)

var (
	op string
)

func init() {
	log.SetFlags(log.Ltime | log.Lmsgprefix)
	log.SetPrefix("| ")
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
		log.Fatalf("invalid operation: %q", op)
	}

}
