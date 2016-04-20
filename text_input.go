package uniform

import (
	"bytes"
	"encoding/xml"
	"html/template"
	"strings"
)

type TextInput struct {
	Value        string
	Placeholder  string
	Id           string
	Name         string
	Class        []string
	Disabled     bool
	ReadOnly     bool
	Required     bool
	AutoFocus    bool
	MaxLength    int
	Size         int
	Attributes   map[string]string
	Dir          Dir
	AutoComplete OnOff
}

func (t TextInput) String() (string, error) {
	buffer := bytes.NewBuffer(nil)
	encoder := errEncoder{encoder: xml.NewEncoder(buffer)}

	attrs := make(attributes, 0)
	attrs.add("type", "text")

	attrs.putString("value", t.Value)
	attrs.putString("placeholder", t.Placeholder)
	attrs.putString("id", t.Id)
	attrs.putString("name", t.Name)
	attrs.putString("dir", string(t.Dir))
	attrs.putString("autocomplete", string(t.AutoComplete))
	attrs.putString("class", strings.Join(t.Class, " "))

	attrs.putBool("disabled", t.Disabled)
	attrs.putBool("readonly", t.ReadOnly)
	attrs.putBool("required", t.Required)
	attrs.putBool("autofocus", t.AutoFocus)

	if t.MaxLength != 0 {
		attrs.putInt("maxlength", t.MaxLength)
	}

	if t.Size != 0 {
		attrs.putInt("size", t.Size)
	}

	attrs.putAttributes(t.Attributes)

	encoder.encodeToken(xml.StartElement{
		Name: xml.Name{Local: "input"},
		Attr: attrs.xmlAttr(),
	})
	encoder.flush()

	if encoder.err != nil {
		return "", encoder.err
	}

	return buffer.String(), nil
}

func (t TextInput) HTML() (template.HTML, error) {
	markup, err := t.String()
	return template.HTML(markup), err
}
