#!/bin/bash

set -xe

mkdir -p bin
go tool templ generate
go build -o bin/main ./cmd
