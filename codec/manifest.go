package codec

import (
	"github.com/technosophos/kubelite/v1"
)

type Manifest struct {
	data []byte
	dec  DecodeFunc
	//enc EncodeFunc
}

type DecodeFunc func([]byte, interface{}) error

// Ref returns an ObjectReference with basic information about the object.
//
// This can be used to perform simple operations, as well as to instrospect
// a record enough to know how to unmarshal it.
func (m *Manifest) Ref() (*v1.ObjectReference, error) {
	or := &v1.ObjectReference{}
	return or, m.dec(m.data, or)
}
