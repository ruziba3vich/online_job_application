package main

import (
	"log"

	"github.com/ruziba3vich/countries/api"
	"github.com/ruziba3vich/countries/internal/config"
	"github.com/ruziba3vich/countries/internal/service"
	"github.com/ruziba3vich/countries/internal/storage"
)

func main() {
	configs, err := config.New()
	if err != nil {
		log.Fatal(err)
	}
	storage, err := storage.New(configs)
	if err != nil {
		log.Fatal(err)
	}

	api := api.New(service.New(*storage))

	log.Fatal(api.RUN(configs))
}
