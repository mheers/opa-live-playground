package partials

import (
	"github.com/maddalax/htmgo/framework/h"
)

func PolicyPathForm(policyPath string) *h.Element {
	return h.Form(
		h.Id("policyPathForm"),
		h.PostPartial(EvaluatePartial),
		PolicyPathWidget(policyPath),
	)
}

func PolicyPathWidget(policyPath string) *h.Element {
	const submitId = "policy-path-submit-button"
	// Policy Dropdown
	return h.Div(
		h.Class("space-y-2"),
		h.Label(
			h.Class("block text-sm font-medium text-gray-700"),
			h.Text("(optional) Select a Policy - or type in the path"),
		),
		PolicySelector(policyPath, submitId),
		h.Button(
			h.Id(submitId),
			h.Type("submit"),
			h.Hidden(),
		),
	)
}
