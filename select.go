package uniform

import (
	"bytes"
	"encoding/xml"
	"html/template"
)

type SelectChild interface{}

type Select struct {
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

	encoder.encodeToken(xml.StartElement{
		Name: xml.Name{Local: "select"},
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
	optGroupAttrs := make(attributes, 0)
	optGroupAttrs.addString("label", optGroup.Label)
	optGroupAttrs.addBoolean("disabled", optGroup.Disabled)

	encoder.encodeToken(xml.StartElement{
		Name: xml.Name{Local: "optgroup"},
		Attr: optGroupAttrs.xmlAttr(),
	})

	for _, option := range optGroup.Options {
		encodeOption(encoder, option)
	}

	encoder.encodeToken(xml.EndElement{Name: xml.Name{Local: "optgroup"}})
}

func encodeOption(encoder errEncoder, option Option) {
	optionAttrs := make(attributes, 0)
	optionAttrs.addString("value", option.Value)
	optionAttrs.addBoolean("selected", option.Selected)
	optionAttrs.addBoolean("disabled", option.Disabled)

	encoder.encodeToken(xml.StartElement{
		Name: xml.Name{Local: "option"},
		Attr: optionAttrs.xmlAttr(),
	})

	if option.Text != "" {
		encoder.encodeToken(xml.CharData(option.Text))
	}

	encoder.encodeToken(xml.EndElement{Name: xml.Name{Local: "option"}})
}
