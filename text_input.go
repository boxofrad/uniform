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
	Data         map[string]string
	Dir          Dir
	AutoComplete OnOff
}

func (t TextInput) String() (string, error) {
	buffer := bytes.NewBuffer(nil)
	encoder := xml.NewEncoder(buffer)

	attrs := make(attributes, 0)
	attrs.add("type", "text")

	attrs.addString("value", t.Value)
	attrs.addString("placeholder", t.Placeholder)
	attrs.addString("id", t.Id)
	attrs.addString("name", t.Name)
	attrs.addString("dir", string(t.Dir))
	attrs.addString("autocomplete", string(t.AutoComplete))

	attrs.addBoolean("disabled", t.Disabled)
	attrs.addBoolean("readonly", t.ReadOnly)
	attrs.addBoolean("required", t.Required)
	attrs.addBoolean("autofocus", t.AutoFocus)

	if len(t.Class) != 0 {
		attrs.add("class", strings.Join(t.Class, " "))
	}

	if t.MaxLength != 0 {
		attrs.addInt("maxlength", t.MaxLength)
	}

	if t.Size != 0 {
		attrs.addInt("size", t.Size)
	}

	for key, value := range t.Data {
		attrs.addString("data-"+key, value)
	}

	start := xml.StartElement{
		Name: xml.Name{Local: "input"},
		Attr: attrs.xmlAttr(),
	}

	if err := encoder.EncodeToken(start); err != nil {
		return "", err
	}

	if err := encoder.Flush(); err != nil {
		return "", err
	}

	return buffer.String(), nil
}

func (t TextInput) HTML() (template.HTML, error) {
	markup, err := t.String()
	return template.HTML(markup), err
}
