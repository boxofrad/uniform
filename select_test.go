package uniform

import "testing"

func TestSelect(t *testing.T) {
	examples := []struct {
		input  Select
		markup string
	}{
		{
			Select{},
			`<select></select>`,
		},
		{
			Select{Name: "fruit"},
			`<select name="fruit"></select>`,
		},
		{
			Select{Id: "fruit"},
			`<select id="fruit"></select>`,
		},
		{
			Select{Disabled: true},
			`<select disabled="disabled"></select>`,
		},
		{
			Select{Required: true},
			`<select required="required"></select>`,
		},
		{
			Select{AutoFocus: true},
			`<select autofocus="autofocus"></select>`,
		},
		{
			Select{Class: []string{"form-input", "form-input-select"}},
			`<select class="form-input form-input-select"></select>`,
		},
		{
			Select{Size: 10},
			`<select size="10"></select>`,
		},
		{
			Select{Attributes: map[string]string{"data-user-id": "123"}},
			`<select data-user-id="123"></select>`,
		},
		{
			Select{Attributes: map[string]string{"empty": ""}},
			`<select></select>`,
		},
		{
			Select{Dir: LTR},
			`<select dir="ltr"></select>`,
		},
		{
			Select{Dir: RTL},
			`<select dir="rtl"></select>`,
		},
		{
			Select{
				Options: []SelectChild{
					Option{
						Value:    "banana",
						Text:     "Banana",
						Selected: true,
					},

					Option{
						Value:    "orange",
						Text:     "Orange",
						Disabled: true,
					},
				},
			},
			`<select><option value="banana" selected="selected">Banana</option><option value="orange" disabled="disabled">Orange</option></select>`,
		},
		{
			Select{
				Options: []SelectChild{
					OptGroup{
						Label:    "Yellow Fruits",
						Disabled: true,
						Options: []Option{
							Option{
								Value:    "banana",
								Text:     "Banana",
								Selected: true,
								Disabled: true,
							},
						},
					},
				},
			},
			`<select><optgroup label="Yellow Fruits" disabled="disabled"><option value="banana" selected="selected" disabled="disabled">Banana</option></optgroup></select>`,
		},
	}

	for _, example := range examples {
		got, err := example.input.String()

		if err != nil {
			t.Fatal(err)
		}

		if got != example.markup {
			t.Errorf("expected: %s, got: %s", example.markup, got)
		}
	}
}
