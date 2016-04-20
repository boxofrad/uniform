package uniform

import (
	"encoding/xml"
	"fmt"
)

type attribute struct {
	name, value string
}

type attributes []attribute

func (a *attributes) add(name, value string) {
	*a = append(*a, attribute{name, value})
}

func (a *attributes) addBoolean(name string, value bool) {
	if value {
		a.add(name, name)
	}
}

func (a *attributes) addString(name, value string) {
	if value != "" {
		a.add(name, value)
	}
}

func (a *attributes) addInt(name string, value int) {
	a.add(name, fmt.Sprintf("%d", value))
}

func (a attributes) xmlAttr() []xml.Attr {
	attr := make([]xml.Attr, len(a))

	for idx, attribute := range a {
		attr[idx] = xml.Attr{
			Name:  xml.Name{Local: attribute.name},
			Value: attribute.value,
		}
	}

	return attr
}
