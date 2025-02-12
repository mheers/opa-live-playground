package pages

import (
	"github.com/maddalax/htmgo/framework/h"
	"github.com/mheers/opa-live-playground/partials"
)

func IndexPage(ctx *h.RequestContext) *h.Page {
	return RootPage(
		h.Div(
			h.Class("bg-white shadow-md rounded-lg p-6 w-full max-w-6xl"),
			h.H3(
				h.Text("OPA Playground"),
				h.Class("text-2xl font-bold mb-6 text-center"),
			),

			h.Div(
				h.Class("flex space-x-4"),
				h.Div(
					h.Class("w-1/2"),
					partials.UserForm(""),
				),
				h.Div(
					h.Class("w-1/2"),
					partials.PolicyForm(""),
				),
			),

			h.Br(),
			h.Br(),

			h.Div(
				h.Class("flex space-x-4 w-full max-w-2xl mx-auto"),
				h.Div(
					h.Class("w-1/2"),
					partials.PresetForm(""),
				),
				h.Div(
					h.Class("w-1/2"),
					partials.PolicyPathForm(""),
				),
			),

			h.Br(),

			h.Div(
				h.Class("space-y-4"),
				partials.EvaluateForm("", "", "", nil),
			),
		),
	)
}
