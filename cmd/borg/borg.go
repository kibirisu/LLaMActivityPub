package main

import (
	"context"
	"log"
	"net/http"

	"borg/pkg/config"
	"borg/pkg/db"
	"borg/pkg/router"
)

func main() {
	ctx := context.Background()
	conf := config.GetConfig()

	db, err := db.GetDB(ctx, conf.DatabaseUrl)
	if err != nil {
		log.Fatal(err)
	}

	r := router.NewRouter(conf.AppEnv, db)
	if err = http.ListenAndServe(":"+conf.ListenPort, r); err != nil {
		log.Fatal(err)
	}
}
