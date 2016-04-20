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
			TextInput{Placeholder: "Alan Partridge"},
			`<input type="text" placeholder="Alan Partridge">`,
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
			TextInput{Required: true},
			`<input type="text" required="required">`,
		},
		{
			TextInput{AutoFocus: true},
			`<input type="text" autofocus="autofocus">`,
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
		{
			TextInput{Attributes: map[string]string{"data-user-id": "123"}},
			`<input type="text" data-user-id="123">`,
		},
		{
			TextInput{Attributes: map[string]string{"empty": ""}},
			`<input type="text">`,
		},
		{
			TextInput{Dir: LTR},
			`<input type="text" dir="ltr">`,
		},
		{
			TextInput{Dir: RTL},
			`<input type="text" dir="rtl">`,
		},
		{
			TextInput{AutoComplete: On},
			`<input type="text" autocomplete="on">`,
		},
		{
			TextInput{AutoComplete: Off},
			`<input type="text" autocomplete="off">`,
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
