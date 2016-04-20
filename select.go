package uniform

import (
	"bytes"
	"encoding/xml"
)

type Select struct {
	Options []Option
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

	for _, option := range s.Options {
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

	encoder.encodeToken(xml.EndElement{Name: xml.Name{Local: "select"}})
	encoder.flush()

	if encoder.err != nil {
		return "", encoder.err
	}

	return buffer.String(), nil
}
