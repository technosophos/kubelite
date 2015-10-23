package types

import (
	"encoding/json"
	"fmt"
	"strconv"
)

// IntOrString is a type that can hold an int or a string.  When used in
// JSON or YAML marshalling and unmarshalling, it produces or consumes the
// inner type.  This allows you to have, for example, a JSON field that can
// accept a name or number.
type IntOrString struct {
	Kind   IntstrKind
	IntVal int
	StrVal string
}

// IntstrKind represents the stored type of IntOrString.
type IntstrKind int

const (
	IntstrInt    IntstrKind = iota // The IntOrString holds an int.
	IntstrString                   // The IntOrString holds a string.
)

// NewIntOrStringFromInt creates an IntOrString object with an int value.
func NewIntOrStringFromInt(val int) IntOrString {
	return IntOrString{Kind: IntstrInt, IntVal: val}
}

// NewIntOrStringFromString creates an IntOrString object with a string value.
func NewIntOrStringFromString(val string) IntOrString {
	return IntOrString{Kind: IntstrString, StrVal: val}
}

// UnmarshalJSON implements the json.Unmarshaller interface.
func (intstr *IntOrString) UnmarshalJSON(value []byte) error {
	if value[0] == '"' {
		intstr.Kind = IntstrString
		return json.Unmarshal(value, &intstr.StrVal)
	}
	intstr.Kind = IntstrInt
	return json.Unmarshal(value, &intstr.IntVal)
}

// String returns the string value, or Itoa's the int value.
func (intstr *IntOrString) String() string {
	if intstr.Kind == IntstrString {
		return intstr.StrVal
	}
	return strconv.Itoa(intstr.IntVal)
}

// MarshalJSON implements the json.Marshaller interface.
func (intstr IntOrString) MarshalJSON() ([]byte, error) {
	switch intstr.Kind {
	case IntstrInt:
		return json.Marshal(intstr.IntVal)
	case IntstrString:
		return json.Marshal(intstr.StrVal)
	default:
		return []byte{}, fmt.Errorf("impossible IntOrString.Kind")
	}
}
