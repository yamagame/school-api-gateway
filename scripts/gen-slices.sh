#!/bin/bash
cat << EOS > ./infra/model/slices.go
package model

import "github.com/yamagame/school-api-gateway/pkg/irconv"

EOS
cat ./datasets/sample-school-schema.tmpl | go run cmd/codegen/main.go ./datasets/sample-school-schema.json >> ./infra/model/slices.go
