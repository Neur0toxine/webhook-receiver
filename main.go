package main

import (
    "log"
    "net/http"
	"os"
)

func main() {
	go processSignals()
	addr := ":8080"

	if len(os.Args) > 1 {
		addr = os.Args[1]
	}

    log.Printf("Listening %s...\n\n", addr)
    err := http.ListenAndServe(addr, NewHandler())
    if err != nil {
        log.Fatal(err)
    }
}
