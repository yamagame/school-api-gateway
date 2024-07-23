#!/bin/bash
#
# ./scripts/csv2schema.sh ./datasets/sample-school-schema.csv
#
CSVFILE=$1
cat $CSVFILE | go run cmd/csv2json/main.go fill | go run cmd/jsonconv/main.go > ${CSVFILE%.*}.json
