package web

import (
	"embed"
	"io/fs"
)

//go:embed dist
var assets embed.FS

func GetAssets() (res fs.FS, err error) {
	res, err = fs.Sub(assets, "dist")
	return
}
