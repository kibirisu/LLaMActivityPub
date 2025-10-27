package main

import (
	"context"
	"io/fs"
	"log"

	"borg/pkg/config"
	"borg/pkg/db"
	"borg/pkg/router"
)

type contextKey string

var (
	assets fs.FS
	dbKey  contextKey = "db"
)

func main() {
	ctx := context.Background()
	conf := config.GetConfig()

	db, err := db.GetDB(ctx, conf.DatabaseDriver, conf.DatabaseUrl)
	if err != nil {
		log.Fatal(err)
	}

	router.Serve(conf.AppEnv, conf.ListenPort, db)
}
