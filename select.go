package uniform

import (
	"bytes"
	"encoding/xml"
	"html/template"
	"strings"
)

type SelectChild interface{}

type Select struct {
	Id         string
	Name       string
	Disabled   bool
	Required   bool
	AutoFocus  bool
	Class      []string
	Size       int
	Attributes map[string]string
	Dir        Dir

	Options []SelectChild
}

type OptGroup struct {
	Label    string
	Disabled bool
	Options  []Option
}

type Option struct {
	Value    string
	Text     string
	Selected bool
	Disabled bool
}

func (s Select) String() (string, error) {
	buffer := bytes.NewBuffer(nil)
	encoder := errEncoder{encoder: xml.NewEncoder(buffer)}

	attrs := make(attributes, 0)
	attrs.putString("id", s.Id)
	attrs.putString("name", s.Name)
	attrs.putString("dir", string(s.Dir))
	attrs.putString("class", strings.Join(s.Class, " "))

	attrs.putBool("disabled", s.Disabled)
	attrs.putBool("required", s.Required)
	attrs.putBool("autofocus", s.AutoFocus)

	if s.Size != 0 {
		attrs.putInt("size", s.Size)
	}

	attrs.putAttributes(s.Attributes)

	encoder.encodeToken(xml.StartElement{
		Name: xml.Name{Local: "select"},
		Attr: attrs.xmlAttr(),
	})

	for _, child := range s.Options {
		switch child.(type) {
		case OptGroup:
			encodeOptGroup(encoder, child.(OptGroup))
		case Option:
			encodeOption(encoder, child.(Option))
		}
	}

	encoder.encodeToken(xml.EndElement{Name: xml.Name{Local: "select"}})
	encoder.flush()

	if encoder.err != nil {
		return "", encoder.err
	}

	return buffer.String(), nil
}

func (t Select) HTML() (template.HTML, error) {
	markup, err := t.String()
	return template.HTML(markup), err
}

func encodeOptGroup(encoder errEncoder, optGroup OptGroup) {
	attrs := make(attributes, 0)
	attrs.putString("label", optGroup.Label)
	attrs.putBool("disabled", optGroup.Disabled)

	encoder.encodeToken(xml.StartElement{
		Name: xml.Name{Local: "optgroup"},
		Attr: attrs.xmlAttr(),
	})

	for _, option := range optGroup.Options {
		encodeOption(encoder, option)
	}

	encoder.encodeToken(xml.EndElement{Name: xml.Name{Local: "optgroup"}})
}

func encodeOption(encoder errEncoder, option Option) {
	attrs := make(attributes, 0)
	attrs.putString("value", option.Value)
	attrs.putBool("selected", option.Selected)
	attrs.putBool("disabled", option.Disabled)

	encoder.encodeToken(xml.StartElement{
		Name: xml.Name{Local: "option"},
		Attr: attrs.xmlAttr(),
	})

	if option.Text != "" {
		encoder.encodeToken(xml.CharData(option.Text))
	}

	encoder.encodeToken(xml.EndElement{Name: xml.Name{Local: "option"}})
}
