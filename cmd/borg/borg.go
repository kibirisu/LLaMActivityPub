package main

import (
	"io/fs"
	"log"
	"net/http"
	"strings"

	"borg/pkg/config"
	"borg/pkg/db"
	"borg/web"
)

var assets fs.FS

func main() {
	conf := config.GetConfig()

	db.GetDB(conf.DatabaseDriver, conf.DatabaseUrl)

	var err error
	assets, err = web.GetAssets()
	if err != nil {
		log.Fatal(err)
	}

	mux := http.NewServeMux()

	// Define routes
	if conf.AppEnv == "prod" {
		mux.HandleFunc("/", handleApp)
	}

	// Start the HTTP server
	addr := ":8080"
	log.Printf("🚀 Server running at http://localhost%s\n", addr)
	if err := http.ListenAndServe(addr, mux); err != nil {
		log.Fatalf("❌ Server failed: %v", err)
	}
}

// Can be done more effectively
func handleApp(w http.ResponseWriter, r *http.Request) {
	file := strings.TrimPrefix(r.URL.Path, "/")
	info, err := fs.Stat(assets, file)
	if err != nil {
		log.Println(err)
		file = "index.html"
	} else if !info.Mode().IsRegular() {
		file = "index.html"
	}
	http.ServeFileFS(w, r, assets, file)
}
