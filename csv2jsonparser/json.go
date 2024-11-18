package csv2jsonparser

import (
	"sort"
	"strings"
)

const Delimiter = "."

// This struct represents the json
// it contains map of fields which can be single string or nested json
type Json struct {
	fields map[string]*JsonValue
}

func NewJson() *Json {
	return &Json{
		fields: make(map[string]*JsonValue),
	}
}

// this function converts the json to string
// it iterates over the fields and recursively calls the tostring method of the field
// if the field is single string value then it returns the string
// if the field is nested json then it recursively calls the nested field tostring method
func (j *Json) ToString() string {
	var res strings.Builder
	keys := make([]string, 0)
	for key, _ := range j.fields {
		keys = append(keys, key)
	}
	res.WriteString("{")
	sort.Strings(keys)
	for index, key := range keys {
		res.WriteString(key)
		res.WriteString(":")
		res.WriteString(j.fields[key].ToString())
		if index != len(keys)-1 {
			res.WriteString(",")
		}
	}
	res.WriteString("}")
	return res.String()
}

// this function recursively build structure of json according to the csv to json mapping
// if the field is single string value then it adds the field to the json
// if the field is nested json then it recursively calls the nested field and add value in hierarchy
func (j *Json) AddField(key string, value string) {
	// if there is no dot in the value then it is a single string value
	if !strings.Contains(value, Delimiter) {
		j.fields[value] = NewJsonValueWithValue(key)
	} else {
		valSplit := strings.SplitN(value, Delimiter, 2)
		if j.fields[valSplit[0]] == nil {
			json := NewJson()
			j.fields[valSplit[0]] = NewJsonValueWithNestedField(json)
		}
		j.fields[valSplit[0]].nestedField.AddField(key, valSplit[1])
	}
}
