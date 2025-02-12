package main

import (
	"fmt"
	"io/fs"
	"net/http"
	"os"

	"github.com/maddalax/htmgo/framework/config"
	"github.com/maddalax/htmgo/framework/h"
	"github.com/maddalax/htmgo/framework/service"
	"github.com/mheers/opa-live-playground/__htmgo"
)

func main() {
	locator := service.NewLocator()
	cfg := config.Get()

	// preflight check if OPA_URL is set
	if os.Getenv("OPA_URL") == "" {
		panic("OPA_URL is not set")
	}

	h.Start(h.AppOpts{
		ServiceLocator: locator,
		LiveReload:     true,
		Register: func(app *h.App) {
			sub, err := fs.Sub(GetStaticAssets(), "assets/dist")

			if err != nil {
				panic(err)
			}

			http.FileServerFS(sub)

			// change this in htmgo.yml (public_asset_path)
			app.Router.Handle(fmt.Sprintf("%s/*", cfg.PublicAssetPath),
				http.StripPrefix(cfg.PublicAssetPath, http.FileServerFS(sub)))

			__htmgo.Register(app.Router)
		},
	})
}
