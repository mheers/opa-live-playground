package partials

import (
	"github.com/maddalax/htmgo/framework/h"
	"github.com/maddalax/htmgo/framework/hx"
	"github.com/mheers/opa-live-playground/utils"
	"golang.org/x/exp/slices"
)

func UserPartial(ctx *h.RequestContext) *h.Partial {
	userEmail := ctx.FormValue("userEmail")

	return h.SwapManyPartial(
		ctx,
		UserForm(userEmail),
	)
}

func UserForm(userEmail string) *h.Element {
	return h.Form(
		h.Id("userForm"),
		h.PostPartial(UserPartial),
		UserWidget(userEmail),
	)
}

func UserWidget(userEmail string) *h.Element {
	userJSON := ""
	if userEmail != "" {
		user, err := opaConnection.User(userEmail)
		if err != nil {
			userJSON = err.Error()
		} else {
			userJSON, _ = utils.MarshalInBeautified(user)
		}
	}

	// User Dropdown
	return h.Div(
		h.Class("space-y-2"),
		h.Label(
			h.Class("block text-sm font-medium text-gray-700"),
			h.Text("Inspect a User"),
		),
		UserSelector(userEmail),
		h.Button(
			h.Id("submit-button"),
			h.Type("submit"),
			h.Hidden(),
		),
		h.Div(
			h.Class("ace_editor ace_hidpi ace-tm"),
			h.Id("userJson"),
			h.UnsafeRaw(userJSON),
		),

		h.UnsafeRawScript(`
			var editorUser = ace.edit("userJson");
			editorUser.session.setMode("ace/mode/json");
			editorUser.setShowPrintMargin(false);
			editorUser.setReadOnly(true);
			editorUser.setOptions({
				// fontSize: "10pt",
			});
		`),
	)
}

func UserSelector(userEmail string) *h.Element {
	result, err := opaConnection.Users()
	if err != nil {
		return h.P(
			h.Text("Error getting users"),
		)
	}

	selectElement := h.Select(
		h.Class("w-full rounded-md border-gray-300"),
		h.Name("userEmail"),
		h.Id("user-select"),
		h.Option(
			h.Value(""),
			h.Text("Select User"),
		),
		// h.OnEvent(hx.ChangeEvent, h.EvalJs("document.getElementById('user-json').value = this")),

		// submit form on change
		// h.OnEvent(hx.ChangeEvent, h.EvalJs("document.getElementById('userForm').submit()")),

		// submit form on change
		h.OnEvent(hx.ChangeEvent, h.EvalJs("document.getElementById('submit-button').click()")),
	)

	emails := make([]string, 0)
	for _, user := range result {
		emails = append(emails, user["email"].(string))
	}

	// Sort emails
	slices.Sort(emails)

	for _, email := range emails {
		selectElement.AppendChild(
			h.Option(
				h.Value(email),
				h.Text(email),

				func() *h.AttributeR {
					if userEmail == email {
						return h.Selected()
					}
					return &h.AttributeR{}
				}(),
			),
		)
	}

	return selectElement
}
