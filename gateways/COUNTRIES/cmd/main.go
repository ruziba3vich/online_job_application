package main

import (
	"log"
	"os"

	"github.com/ruziba3vich/ishtopuz/countries/api"
	"github.com/ruziba3vich/ishtopuz/countries/config"
)

func main() {
	config, err := config.New()
	if err != nil {
		log.Fatal(err)
	}

	api := api.New(log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile), config)
	log.Fatal(api.RUN())
}
