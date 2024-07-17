#!/bin/bash
CSVFILE=$1
cat $CSVFILE | go run cmd/csv2json/main.go > ./work/test.json
cat ./work/test.json | go run cmd/jsonconv/main.go > ${CSVFILE%.*}.json
