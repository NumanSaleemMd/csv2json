package csv2jsonparser

// ICsvToJsonParser is an interface for CsvToJsonParser
// this set of interfaces will be used to generate the encoded string and convert the csv row to json
type ICsvToJsonParser interface {
	GenerateEncodedString(csvToJsonMapping map[string]string) string
	ConvertToJson(csvRow []string, encodedStr string) string
}
