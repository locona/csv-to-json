.DEFAULT_GOAL := run

run:
	@go install
	@csv-to-json ./test.csv
