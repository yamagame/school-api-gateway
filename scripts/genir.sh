#!/bin/bash
cat ./datasets/sample-school-schema.tmpl | go run cmd/codegen/main.go ./datasets/sample-school-schema.json
