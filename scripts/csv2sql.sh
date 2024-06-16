#!/bin/bash
#
# CSVファイルからSQLを生成する
#
# 使用例 ./scripts/csv2sql.sh ./datasets/sample-school-labos.csv
#

CSVFILE=$1
cat $CSVFILE | go run cmd/csv2sql/main.go > ${CSVFILE%.*}.sql
