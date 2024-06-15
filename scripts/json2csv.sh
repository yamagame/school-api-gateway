#!/bin/bash
#
# JSON形式をCSVファイルに変換する
#
# 使用例 ./scripts/json2csv.sh ./datasets/sample-school-labos.json
#

## jq のインストール
# brew install jq

JSONFILE=$1
echo "name, group, program, building" > ${JSONFILE%.*}.csv
cat $JSONFILE | jq -r '.labs[] | [.name, .group, .program, .building] | @csv' >> ${JSONFILE%.*}.csv
