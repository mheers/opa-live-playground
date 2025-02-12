//go:build !prod
// +build !prod

package main

import (
	"io/fs"

	"github.com/mheers/opa-live-playground/internal/embedded"
)

func GetStaticAssets() fs.FS {
	return embedded.NewOsFs()
}
