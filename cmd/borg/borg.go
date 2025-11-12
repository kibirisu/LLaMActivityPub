package main

import (
	"context"
	"log"
	"net/http"

	"borg/internal/config"
	"borg/internal/domain"
	"borg/internal/router"
)

func main() {
	ctx := context.Background()
	conf := config.GetConfig()

	ds, err := domain.NewDataStore(ctx, conf.DatabaseURL)
	if err != nil {
		log.Fatal(err)
	}

	r := router.NewRouter(ds)
	_ = http.ListenAndServe(":"+conf.ListenPort, r)
}
