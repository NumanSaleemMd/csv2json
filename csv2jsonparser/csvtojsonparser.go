package csv2jsonparser

import "strings"

type CsvToJsonParser struct {
	csvHeaderIndex map[string]int
}

func NewCsvToJsonParser(csvHeader []string) *CsvToJsonParser {
	return &CsvToJsonParser{csvHeaderIndex: storeCsvHeader(csvHeader)}
}

func (j *CsvToJsonParser) GenerateEncodedString(csvToJsonMapping map[string]string) string {
	// parse the json string and populate the fields
	res := NewJson()
	for key, value := range csvToJsonMapping {
		res.AddField(key, value)
	}
	return res.ToString()
}

// store the csv header in map with index
// eg:
// csvHeader = ["name", "age", "city"]
// csvHeaderIndex = {"name": 0, "age": 1, "city": 2}
func storeCsvHeader(csvHeader []string) map[string]int {
	csvHeaderIndex := make(map[string]int)
	for index, header := range csvHeader {
		csvHeaderIndex[header] = index
	}
	return csvHeaderIndex
}

// takes the encode json structure and replace it with actual value of row
// we will use csvHeaderIndex to get the index of the field in csv row and replace it in json structure
func (j *CsvToJsonParser) ConvertToJson(csvRow []string, encodedStr string) string {
	res := encodedStr
	for key, value := range j.csvHeaderIndex {
		res = strings.ReplaceAll(res, key, csvRow[value])
	}
	return res
}
