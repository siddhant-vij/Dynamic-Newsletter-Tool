#!/bin/bash

go mod tidy
go build -o newsletter-tool ../cmd/main.go
./newsletter-tool

rm newsletter-tool