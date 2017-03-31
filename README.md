# CSV-to-JSON
A very simple tool for pivoting rows in a CSV file into an 
array of JSON objects

# Compiling
If you have a golang compiler installed, then you're good to go. Just run
`go build main.go`.

If you don't have go development tools set up on your computer, check out the golang
installation documentation at https://golang.org/doc/install

# Running
The program currently takes one argument, which should be the path to a CSV file. For example,
to convert a file named `test.csv` to a JSON file named `test.json`, run:
`./csv-to-json test.csv`

# Contributing
Pull requests for features and fixes are welcome!
