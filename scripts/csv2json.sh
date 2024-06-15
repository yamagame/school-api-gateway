#!/bin/bash
#
# CSVファイルをJSON形式に変換する
#
# 使用例 ./scripts/csv2json.sh ./datasets/sample-school-labos.csv
#

JSONFILE=$1
jq -n --argjson labs "`cat $JSONFILE | go run cmd/csv2json/main.go`" '{"labs":$labs}' > ${JSONFILE%.*}.json
