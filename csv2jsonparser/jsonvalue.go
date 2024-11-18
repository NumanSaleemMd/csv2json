package csv2jsonparser

// This struct represents the json value
// It can be either a string value or a nested field
type JsonValue struct {
	val         string
	nestedField *Json
}

// Constructor for nested json field
func NewJsonValueWithNestedField(nestedField *Json) *JsonValue {
	return &JsonValue{
		nestedField: nestedField,
	}
}

// Constructor for string value
func NewJsonValueWithValue(val string) *JsonValue {
	return &JsonValue{
		val: val,
	}
}

// this function prints the json value
// if it is string value then it returns the string
// if it is nested field then it recursively calls the nested field tostring method
func (j *JsonValue) ToString() string {
	if j.val != "" {
		return j.val
	} else {
		return j.nestedField.ToString()
	}

}
