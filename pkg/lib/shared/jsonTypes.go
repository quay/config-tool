package shared

import (
	"strconv"

	"gopkg.in/yaml.v3"
)

// IntOrString is a type that can hold an int or a string.  When used in
// JSON or YAML marshalling and unmarshalling, it produces or consumes the
// inner type.  This allows you to have, for example, a JSON field that can
// accept a name or number.
type IntOrString struct {
	Type   Type
	IntVal int
	StrVal string
}

// Type represents the stored type of IntOrString.
type Type int

const (
	Int    = 1 // The IntOrString holds an int.
	String = 2 // The IntOrString holds a string.
)

// UnmarshalJSON implements the yaml.Unmarshaller interface.
func (intOrString *IntOrString) UnmarshalYAML(value *yaml.Node) error {
	if value == nil {
		return nil
	}
	if value.Tag == "!!str" {
		intOrString.Type = String
		return value.Decode(&intOrString.StrVal)
	}
	intOrString.Type = Int
	return value.Decode(&intOrString.IntVal)

}

// String returns the string value, or the Itoa of the int value.
func (intOrString *IntOrString) String() string {
	if intOrString == nil {
		return "<nil>"
	}
	if intOrString.Type == String {
		return intOrString.StrVal
	}
	return strconv.Itoa(intOrString.IntValue())
}

// IntValue returns the IntVal if type Int, or if
// it is a String, will attempt a conversion to int,
// returning 0 if a parsing error occurs.
func (intOrString *IntOrString) IntValue() int {
	if intOrString.Type == String {
		i, _ := strconv.Atoi(intOrString.StrVal)
		return i
	}
	return int(intOrString.IntVal)
}

// MarshalJSON implements the json.Marshaller interface.
func (intOrString *IntOrString) MarshalYAML() (interface{}, error) {

	if intOrString == nil {
		return []byte{}, nil
	}
	switch intOrString.Type {
	case Int:
		return intOrString.IntVal, nil
	case String:
		return intOrString.StrVal, nil
	default:
		return []byte{}, nil
	}

}
