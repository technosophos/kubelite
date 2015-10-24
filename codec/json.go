package codec

import (
	"encoding/json"
)

type jsonDecoder struct {
	data []byte
}

func (d jsonDecoder) For() string {
	return "json"
}

// All returns all documents in the original.
// JSON does not really support multi-doc. But this provides an interface
// compatible with the YAML implementation.
func (d jsonDecoder) All() ([]*Manifest, error) {
	t, err := d.One()
	return []*Manifest{t}, err
}

func (d jsonDecoder) One() (*Manifest, error) {
	return &Manifest{
		data: d.data,
		dec: func(b []byte, v interface{}) error {
			return json.Unmarshal(b, v)
		},
	}, nil
}
