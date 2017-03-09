package main

import (
	"os"
	"encoding/csv"
	"encoding/json"
	"flag"
	"strings"
	"fmt"
)

// Returns a handle to an existing file
func getInfile(filename string) *os.File {
	f, err := os.Open(filename)
	check(err)
	return f
}

// Returns a handle to a new file
func getOutfile(filename string) *os.File {
	f, err := os.Create(filename)
	check(err)
	return f
}

func writeJSON(infile, outfile *os.File) {
	rd := csv.NewReader(infile)
	rd.LazyQuotes = true
	wr := json.NewEncoder(outfile)
	arr := make([]map[string]string, 0)
	var vals []string;
	keys, rderr := rd.Read()
	check(rderr)
	for rderr == nil {
		vals, rderr = rd.Read()
		if len(vals) < len(keys) {
			continue;
		}
		vmap := make(map[string]string)
		for idx, key := range keys {
			vmap[key] = vals[idx]
		}
		arr = append(arr, vmap)
	}
	wr.Encode(arr)
	fmt.Printf("Wrote %d JSON records\n", len(arr))
}

// Yikes
func check(e error) {
    if e != nil {
        panic(e)
    }
}

func main() {
	flag.Parse()
	infname := flag.Arg(0)
	// Output file name
	var outfname = strings.Replace(infname, ".csv", ".json", 1)
	infile := getInfile(infname)
	defer infile.Close()
	outfile := getOutfile(outfname)
	defer outfile.Close()
	writeJSON(infile, outfile)
}


