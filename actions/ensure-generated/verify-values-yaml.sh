#!/bin/bash

VALUES_SCHEMA=$(find ./helm -maxdepth 2 -name values.schema.json)
VALUES=$(find ./helm -maxdepth 2 -name values.yaml)

if diff --unified "${VALUES}" <(helm-values-gen "${VALUES_SCHEMA}") ; then
  echo "values.yaml is up to date"
  exit 0
else
  echo "values.yaml is out of date"
  exit 1
fi
