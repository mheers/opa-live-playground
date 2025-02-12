package partials

import (
	"os"
	"raygun/types"
	"strings"

	"github.com/maddalax/htmgo/framework/h"
	"github.com/mheers/opa-live-playground/utils"
)

var opaConnection = &utils.OPAConnection{
	URL: os.Getenv("OPA_URL"),
}

func EvaluatePartial(ctx *h.RequestContext) *h.Partial {
	input := ctx.Request.FormValue("inputContent")
	policyPath := ctx.Request.FormValue("policyPath")
	presetName := ctx.Request.FormValue("presetName")

	errors := make([]string, 0)

	var preset *types.TestRecord
	if presetName != "" {
		p, err := utils.GetPreset(presetName)
		if err != nil {
			errors = append(errors, "Preset not found")
		}

		preset = p
		policyPath = preset.DecisionPath
		input = preset.Input.Value
	}

	evaluationResult, err := opaConnection.Evaluate(policyPath, input)
	if err != nil {
		errors = append(errors, "Error evaluating input")
	}

	resultBeautified, _ := utils.BeautifyJson(evaluationResult)

	if len(errors) > 0 {
		resultBeautified = strings.Join(errors, "\n")
	}

	return h.SwapManyPartial(
		ctx,
		PresetForm(presetName),
		PolicyPathForm(policyPath),
		EvaluateForm(input, resultBeautified, policyPath, preset),
	)
}

func EvaluateForm(input, evaluationResult, policyPath string, preset *types.TestRecord) *h.Element {
	inputBeautified, err := utils.BeautifyJson(input)
	if err != nil {
		inputBeautified = input
	}

	return h.Form(
		h.Class("flex flex-col gap-4 w-full max-w-2xl mx-auto"),
		h.Id("evaluation-form"),
		h.PostPartial(EvaluatePartial),

		// Policy Path Input
		h.Div(
			h.Class("space-y-2"),
			h.Label(
				h.Class("block text-sm font-medium text-gray-700"),
				h.Text("Policy Path"),
			),
			h.TextInput(
				h.Class("w-full rounded-md border-black-300 mt-2"),
				h.Id("policy-path"),
				h.Name("policyPath"),
				h.Placeholder("Policy Path"),
				h.Value(policyPath),
			),
		),

		// Input Data Textarea
		h.Div(
			h.Class("space-y-2"),
			h.Label(
				h.Class("block text-sm font-medium text-gray-700"),
				h.Text("Input Data"),
			),
			h.Div(
				h.Class("ace_editor ace_hidpi ace-tm"),
				h.Id("inputData"),
				h.UnsafeRaw(inputBeautified),
			),
			h.TextInput(
				h.Name("inputContent"),
				h.Id("inputContent"),
				h.Value(inputBeautified),
				h.Hidden(),
			),
		),

		h.UnsafeRawScript(`
			var editorresult = ace.edit("inputData");
			var textinput = document.querySelector("input[name=input]");
			editorresult.session.setMode("ace/mode/json");
			editorresult.setOptions({
				// fontSize: "10pt",
				wrap: false
			});
			editorresult.getSession().on("change", function (e, f) {
				const value = f.getValue();
				console.log(value);
				document.getElementById('inputContent').value = value;
			});
		`),

		// Evaluate Button
		h.Div(
			h.Class("flex justify-center"),
			SubmitButton("Evaluate Policy"),
		),

		// Result Textarea
		h.Div(
			h.Class("space-y-2"),
			h.Label(
				h.Class("block text-sm font-medium text-gray-700"),
				h.Text("Evaluation Result"),
			),
			h.Div(
				h.Class("ace_editor ace_hidpi ace-tm"),
				h.Id("evaluationResult"),
				h.UnsafeRaw(evaluationResult),
			),
		),

		h.UnsafeRawScript(`
			var editorresult = ace.edit("evaluationResult");
			editorresult.session.setMode("ace/mode/json");
			editorresult.setReadOnly(true);
			editorresult.setOptions({
				// fontSize: "10pt"
			});
		`),
	)
}

func SubmitButton(text string) *h.Element {
	return h.Button(
		h.Class("bg-rose-400 hover:bg-rose-500 text-white font-bold py-2 px-4 rounded"),
		h.Id("swap-text"),
		h.Type("submit"),
		h.Text(text),
	)
}
