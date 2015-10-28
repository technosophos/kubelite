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

// Object decodes the manifest into the given object.
//
// You can use ObjectReference.Kind to figure out what kind of object to
// decode into.
//
// There are several shortcut methods that will allow you to decode directly
// to one of the common types, like Pod(), RC(), and Service().
func (m *Manifest) Object(v interface{}) error {
	return m.dec(m.data, v)
}

// Pod decodes a manifest into a Pod.
func (m *Manifest) Pod() (*v1.Pod, error) {
	o := new(v1.Pod)
	return o, m.Object(o)
}

// RC decodes a manifest into a ReplicationController.
func (m *Manifest) RC() (*v1.ReplicationController, error) {
	o := new(v1.ReplicationController)
	return o, m.Object(o)
}

// Service decodes a manifest into a Service
func (m *Manifest) Service() (*v1.Service, error) {
	o := new(v1.Service)
	return o, m.Object(o)
}

// PersistentVolume decodes a manifest into a PersistentVolume
func (m *Manifest) PersistentVolume() (*v1.PersistentVolume, error) {
	o := new(v1.PersistentVolume)
	return o, m.Object(o)
}

// Secret decodes a manifest into a Secret
func (m *Manifest) Secret() (*v1.Secret, error) {
	o := new(v1.Secret)
	return o, m.Object(o)
}

// Namespace decodes a manifest into a Namespace
func (m *Manifest) Namespace() (*v1.Namespace, error) {
	o := new(v1.Namespace)
	return o, m.Object(o)
}
