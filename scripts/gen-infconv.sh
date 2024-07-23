#!/bin/bash
cat ./datasets/sample-school-infconv.tmpl | go run cmd/codegen/main.go ./datasets/sample-school-schema.json Labos
