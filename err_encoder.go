package uniform

import "encoding/xml"

type errEncoder struct {
	encoder *xml.Encoder
	err     error
}

func (e errEncoder) encodeToken(token xml.Token) {
	if e.err != nil {
		return
	}

	e.err = e.encoder.EncodeToken(token)
}

func (e errEncoder) flush() {
	if e.err != nil {
		return
	}

	e.err = e.encoder.Flush()
}
