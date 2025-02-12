package partials

import (
	"github.com/maddalax/htmgo/framework/h"
	"github.com/maddalax/htmgo/framework/hx"
	"golang.org/x/exp/slices"
)

func PolicyPartial(ctx *h.RequestContext) *h.Partial {
	policyPath := ctx.FormValue("policyPath")

	return h.SwapManyPartial(
		ctx,
		PolicyForm(policyPath),
	)
}

func PolicyForm(policyPath string) *h.Element {
	return h.Form(
		h.Id("policyForm"),
		h.PostPartial(PolicyPartial),
		PolicyWidget(policyPath),
	)
}

func PolicyWidget(policyPath string) *h.Element {
	policyRego := ""
	if policyPath != "" {
		policies, err := opaConnection.Policies()
		if err != nil {
			policyRego = err.Error()
		}

		for _, p := range policies {
			pathToCheck := p.Path()
			if pathToCheck == policyPath {
				policyRego = p.Raw
				break
			}
		}
	}

	// Policies Dropdown
	return h.Div(
		h.Class("space-y-2"),
		h.Label(
			h.Class("block text-sm font-medium text-gray-700"),
			h.Text("Inspect a Policy"),
		),
		PolicySelector(policyPath, "policy-submit-button"),
		h.Button(
			h.Id("policy-submit-button"),
			h.Type("submit"),
			h.Hidden(),
		),

		h.Div(
			h.Class("ace_editor ace_hidpi ace-tm"),
			h.Id("policyRego"),
			h.UnsafeRaw(policyRego),
		),

		h.UnsafeRawScript(`
			var editor = ace.edit("policyRego");
			editor.session.setMode("ace/mode/golang");
			editor.setShowPrintMargin(false);
			editor.setReadOnly(true);
			editor.setOptions({
				// fontSize: "10pt",
			});
		`),
	)
}

func PolicySelector(policyPath, submitId string) *h.Element {
	result, err := opaConnection.Policies()
	if err != nil {
		return h.P(
			h.Text("Error getting policies"),
		)
	}

	selectElement := h.Select(
		h.Class("w-full rounded-md border-gray-300"),
		h.Name("policyPath"),
		h.Id("policy-select-"+submitId+""),
		h.Option(
			h.Value(""),
			h.Text("Select a Policy"),
		),
		// submit form on change
		h.OnEvent(hx.ChangeEvent, h.EvalJs("document.getElementById('"+submitId+"').click()")),
	)

	paths := make([]string, 0)
	for _, policy := range result {
		paths = append(paths, policy.Path())
	}

	// Sort paths
	slices.Sort(paths)

	for _, path := range paths {
		selectElement.AppendChild(
			h.Option(
				h.Value(path),
				h.Text(path),

				func() *h.AttributeR {
					if policyPath == path {
						return h.Selected()
					}
					return &h.AttributeR{}
				}(),
			),
		)
	}

	return selectElement
}
