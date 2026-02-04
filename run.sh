#!/bin/bash

set -xe

cd cmd

go tool templ generate
go run .

cd ..
