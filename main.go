package main

import (
	"github.com/alknopfler/ztp-api-signer/api"
	"github.com/alknopfler/ztp-api-signer/config"
	"log"
	"net/http"
)

func main() {

	err := http.ListenAndServe(config.SrvPort, api.HandlerController())
	if err != nil {
		log.Fatal(err)
	}
}
