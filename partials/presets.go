package partials

import (
	"fmt"

	"github.com/maddalax/htmgo/framework/h"
	"github.com/maddalax/htmgo/framework/hx"
	"github.com/mheers/opa-live-playground/utils"
	"golang.org/x/exp/slices"
)

func PresetForm(presetName string) *h.Element {
	return h.Form(
		h.Id("presetForm"),
		h.PostPartial(EvaluatePartial),
		PresetWidget(presetName),
	)
}

func PresetWidget(presetName string) *h.Element {
	presetInput := ""
	if presetName != "" {
		preset, err := utils.GetPreset(presetName)
		if err != nil {
			presetInput = err.Error()
		} else {
			presetInput = preset.Input.Value
		}
		fmt.Println(presetInput)
	}

	// Presets Dropdown
	return h.Div(
		h.Class("space-y-2"),
		h.Label(
			h.Class("block text-sm font-medium text-gray-700"),
			h.Text("(optional) Select a Preset"),
		),
		PresetSelector(presetName, "preset-submit-button"),
		h.Button(
			h.Id("preset-submit-button"),
			h.Type("submit"),
			h.Hidden(),
		),
	)
}

func PresetSelector(presetName, submitId string) *h.Element {
	result, err := utils.GetPresets()
	if err != nil {
		return h.P(
			h.Text("Error getting presets"),
		)
	}

	selectElement := h.Select(
		h.Class("w-full rounded-md border-gray-300"),
		h.Name("presetName"),
		h.Id("preset-select"),
		h.Option(
			h.Value(""),
			h.Text("Select Presets"),
		),
		// submit form on change
		h.OnEvent(hx.ChangeEvent, h.EvalJs("document.getElementById('"+submitId+"').click()")),
	)

	paths := make([]string, 0)
	for _, preset := range result {
		paths = append(paths, preset.Name)
	}

	// Sort paths
	slices.Sort(paths)

	for _, path := range paths {
		selectElement.AppendChild(
			h.Option(
				h.Value(path),
				h.Text(path),

				func() *h.AttributeR {
					if presetName == path {
						return h.Selected()
					}
					return &h.AttributeR{}
				}(),
			),
		)
	}

	return selectElement
}
