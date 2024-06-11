#!/bin/bash
#
# JSON形式をCSVファイルに変換する
# 使用例 ./scripts/json2csv.sh ./private/labos.json
#

## jq のインストール
# brew install jq

JSONFILE=$1
echo "name, group, program" > ${JSONFILE%.*}.csv
cat $JSONFILE | jq -r '.labs[] | [.name, .group, .program] | @csv' >> ${JSONFILE%.*}.csv
