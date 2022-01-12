package main

import (
	"bytes"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func post() {
	buffer, err := ioutil.ReadAll(os.Stdin)

	if err != nil {
		log.Fatal(err)
	}

	resp, err := http.Post("https://paste.rs", "text/plain", bytes.NewBuffer(buffer))

	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		log.Fatal(err)
	}

	switch resp.StatusCode {
	case 201:
		log.Print(string(body))
	case 206:
		log.Printf("Partially posted: %s", body)
	default:
		log.Fatalf("Got bad response: %q", resp.Status)
	}

}
