package web

import (
	"embed"
	"io/fs"
)

var assets embed.FS

func GetAssets() (res fs.FS, err error) {
	res, err = fs.Sub(assets, "dist")
	return
}
