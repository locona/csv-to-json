package main

import (
	"encoding/csv"
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"strings"
)

func writeJSON(infile, outfile *os.File) {
	reader := csv.NewReader(infile)
	reader.LazyQuotes = true
	keys, err := reader.Read()
	if err != nil {
		panic(err)
	}

	arr := make([]map[string]interface{}, 0)
	for {
		record, err := reader.Read()
		if err != nil {
			fmt.Println(err)
			break
		}
		if len(record) < len(keys) {
			continue
		}

		vmap := make(map[string]interface{})
		for idx := range keys {
			vmap[keys[idx]] = record[idx]
		}
		arr = append(arr, vmap)
	}

	wr := json.NewEncoder(outfile)
	if err := wr.Encode(arr); err != nil {
		panic(err)
	}
	fmt.Printf("Wrote %d JSON records\n", len(arr))
}

func main() {
	flag.Parse()
	csvFileName := flag.Arg(0)
	jsonFileName := strings.Replace(csvFileName, ".csv", ".json", 1)
	csvF, err := os.Open(csvFileName)
	defer csvF.Close()
	if err != nil {
		panic(err)
	}

	jsonF, err := os.Create(jsonFileName)
	defer jsonF.Close()
	if err != nil {
		panic(err)
	}

	writeJSON(csvF, jsonF)
}
