package main

import (
	"sync"

	"csv2json/csv2jsonparser"
)

const BatchSize = 100

func main() {
	csvFile := [][]string{
		{"id", "name", "dob", "address_street", "address_city"},
		{"1", "John Doe", "1990-01-01", "123 Main St", "Anytown"},
		{"2", "Jane Smith", "1991-02-02", "456 Park Ave", "Othercity"},
	}
	csvToJsonMapping := map[string]string{
		"id":             "userId",
		"name":           "userName",
		"dob":            "dateOfBirth",
		"address_street": "address.street",
		"address_city":   "address.city",
	}
	// first index represent header of file
	csvToJsonConverter := newCsvJsonConvertor(csvFile[0])
	encodedString := csvToJsonConverter.GenerateEncodedString(csvToJsonMapping)
	jsonConvertedVals := fetchParallelData(csvFile[1:], encodedString, csvToJsonConverter)
	for _, res := range jsonConvertedVals {
		println(res)
	}

}

// this can go inside interface
// function convert csv to json in parallel according to batchSize
func fetchParallelData(csvFile [][]string, encodedString string, parser csv2jsonparser.ICsvToJsonParser) []string {
	// concurrency control variable
	var wg sync.WaitGroup
	// specialised data structure for parallel processing and storing value
	var res = make(chan string, len(csvFile))
	// parallel fetching the data in batches
	for i := 0; i < len(csvFile); i += BatchSize {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			// converting the csv to json for current batch
			for j := i; j < i+BatchSize && j < len(csvFile); j++ {
				jsonRes := parser.ConvertToJson(csvFile[j], encodedString)
				res <- jsonRes
			}
		}(i)
	}
	wg.Wait()
	close(res)
	ans := make([]string, 0)
	for val := range res {
		ans = append(ans, val)
	}
	return ans
}

func newCsvJsonConvertor(csvHeader []string) csv2jsonparser.ICsvToJsonParser {
	return csv2jsonparser.NewCsvToJsonParser(csvHeader)
}
