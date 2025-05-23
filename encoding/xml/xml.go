package xml

import (
	"encoding/xml"

	"github.com/fzf-labs/gopkg/encoding"
)

// Name is the name registered for the xml codec.
const Name = "xml"

//nolint:gochecknoinits
func init() {
	encoding.RegisterCodec(codec{})
}

// codec is a Codec implementation with xml.
type codec struct{}

func (codec) Marshal(v any) ([]byte, error) {
	return xml.Marshal(v)
}

func (codec) Unmarshal(data []byte, v any) error {
	return xml.Unmarshal(data, v)
}

func (codec) Name() string {
	return Name
}
