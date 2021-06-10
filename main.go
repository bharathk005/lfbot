package main

import (
	"lfbot/lfserve"
	"log"
	"net/http"
)

func main() {
	log.Printf("starting server")
	http.HandleFunc("/lfbot", lfserve.HandleTelegramWebHook)
	log.Fatal(http.ListenAndServe("0.0.0.0:8443", nil))
	//log.Fatal(http.ListenAndServeTLS("0.0.0.0:8443", "localhost.crt", "localhost.key", nil))
}
