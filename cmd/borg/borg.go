package main

import (
	"context"
	"log"
	"net/http"

	"borg/pkg/config"
	"borg/pkg/data"
	"borg/pkg/router"
)

func main() {
	ctx := context.Background()
	conf := config.GetConfig()

	ds, err := data.NewDataStore(ctx, conf.DatabaseUrl)
	if err != nil {
		log.Fatal(err)
	}

	r := router.NewRouter(conf.AppEnv, ds)
	if err = http.ListenAndServe(":"+conf.ListenPort, r); err != nil {
		log.Fatal(err)
	}
}
