package csv2jsonparser

import (
	"testing"
)

func TestJsonParser_GenerateEncodedString(t *testing.T) {
	parser := &CsvToJsonParser{}
	mapping := map[string]string{
		"value1": "key1",
		"value2": "key2.key3",
		"value3": "key2.key4",
		"value4": "key5.key6.key7",
	}
	expected := "{key1:value1,key2:{key3:value2,key4:value3},key5:{key6:{key7:value4}}}"
	value := parser.GenerateEncodedString(mapping)
	if value != expected {
		t.Errorf("expected %s, got %s", expected, parser.GenerateEncodedString(mapping))
	}
}

func TestCsvToJsonParser_ConvertToJson(t *testing.T) {
	parser := &CsvToJsonParser{csvHeaderIndex: map[string]int{"value1": 0, "value2": 1, "value3": 2, "value4": 3}}
	csvRow := []string{"randomValue1", "randomValue2", "randomValue3", "randomValue4"}
	encodedStr := "{key1:value1,key2:{key3:value2,key4:value3},key5:{key6:{key7:value4}}}"
	expected := "{key1:randomValue1,key2:{key3:randomValue2,key4:randomValue3},key5:{key6:{key7:randomValue4}}}"
	result := parser.ConvertToJson(csvRow, encodedStr)
	if result != expected {
		t.Errorf("expected %s, got %s", expected, result)
	}
}
