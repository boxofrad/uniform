package uniform

import "testing"

func TestTextInput(t *testing.T) {
	examples := []struct {
		input  TextInput
		markup string
	}{
		{
			TextInput{},
			`<input type="text">`,
		},
		{
			TextInput{Value: "Hello"},
			`<input type="text" value="Hello">`,
		},
		{
			TextInput{Value: "<script>"},
			`<input type="text" value="&lt;script&gt;">`,
		},
		{
			TextInput{Name: "first_name"},
			`<input type="text" name="first_name">`,
		},
		{
			TextInput{Id: "first_name"},
			`<input type="text" id="first_name">`,
		},
		{
			TextInput{Disabled: true},
			`<input type="text" disabled="disabled">`,
		},
		{
			TextInput{ReadOnly: true},
			`<input type="text" readonly="readonly">`,
		},
		{
			TextInput{Class: []string{"form-input", "form-input-text"}},
			`<input type="text" class="form-input form-input-text">`,
		},
		{
			TextInput{MaxLength: 10},
			`<input type="text" maxlength="10">`,
		},
		{
			TextInput{Size: 10},
			`<input type="text" size="10">`,
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
