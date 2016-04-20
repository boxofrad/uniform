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
			Select{
				Options: []Option{
					{
						Value:    "banana",
						Text:     "Banana",
						Selected: true,
					},

					{
						Value:    "orange",
						Text:     "Orange",
						Disabled: true,
					},
				},
			},
			`<select><option value="banana" selected="selected">Banana</option><option value="orange" disabled="disabled">Orange</option></select>`,
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
