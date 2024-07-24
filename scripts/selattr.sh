#!/bin/bash
#
# 例) ./scripts/selattr.sh labo.go
#
cat datasets/sample-school-schema.json | jq ".[] | select(.filename == \"${1}\") | .attributes"
