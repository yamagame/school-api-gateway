#!/bin/bash
./scripts/csv2schema.sh ./datasets/sample-school-schema.csv
cat << EOS
package model

import (
	"time"

  "github.com/yamagame/school-api-gateway/pkg/conv"
)

EOS
cat ./datasets/sample-school-schema.tmpl | go run cmd/codegen/main.go ./datasets/sample-school-schema.json
