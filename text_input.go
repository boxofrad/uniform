package uniform

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"html/template"
	"strings"
)

type TextInput struct {
	Value     string
	Id        string
	Name      string
	Class     []string
	Disabled  bool
	ReadOnly  bool
	MaxLength int
	Size      int
}

func (t TextInput) String() (string, error) {
	buffer := bytes.NewBuffer(nil)
	encoder := xml.NewEncoder(buffer)

	attrs := make(attributes, 0)
	attrs.add("type", "text")

	if t.Value != "" {
		attrs.add("value", t.Value)
	}

	if t.Id != "" {
		attrs.add("id", t.Id)
	}

	if t.Name != "" {
		attrs.add("name", t.Name)
	}

	if t.Disabled {
		attrs.add("disabled", "disabled")
	}

	if t.ReadOnly {
		attrs.add("readonly", "readonly")
	}

	if len(t.Class) != 0 {
		attrs.add("class", strings.Join(t.Class, " "))
	}

	if t.MaxLength != 0 {
		attrs.add("maxlength", fmt.Sprintf("%d", t.MaxLength))
	}

	if t.Size != 0 {
		attrs.add("size", fmt.Sprintf("%d", t.Size))
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

func buildAttr(values map[string]string) []xml.Attr {
	attr := make([]xml.Attr, 0)

	for name, value := range values {
		attr = append(attr, xml.Attr{
			Name:  xml.Name{Local: name},
			Value: value,
		})
	}

	return attr
}
