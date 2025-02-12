package pages

import (
	"github.com/maddalax/htmgo/framework/h"
	"github.com/mheers/opa-live-playground/__htmgo/assets"
)

func RootPage(children ...h.Ren) *h.Page {
	title := "OPA Live Playground"
	description := "a playground for Open Policy Agent"
	author := "Marcel Heers"
	url := "https://marcelheers.de"

	return h.NewPage(
		h.Html(
			h.HxExtensions(
				h.BaseExtensions(),
			),
			h.Head(
				h.Title(
					h.Text(title),
				),
				h.Meta("viewport", "width=device-width, initial-scale=1"),
				h.Link(assets.FaviconIco, "icon"),
				h.Link(assets.AppleTouchIconPng, "apple-touch-icon"),
				h.Meta("title", title),
				h.Meta("charset", "utf-8"),
				h.Meta("author", author),
				h.Meta("description", description),
				h.Meta("og:title", title),
				h.Meta("og:url", url),
				h.Link("canonical", url),
				h.Meta("og:description", description),
				h.Link(assets.MainCss, "stylesheet"),
				h.Script(assets.HtmgoJs),

				// ace editor
				h.Script(assets.AceBuildsSrcNoconflictAceJs),
			),
			h.Body(
				h.Div(
					h.Class("bg-gray-100 min-h-screen flex items-center justify-center p-4"),
					h.Fragment(children...),
				),
			),
		),
	)
}
