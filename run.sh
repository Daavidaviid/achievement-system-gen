#!/bin/sh

go generate
go test -cover ./...
if [ $? -eq 0 ]
then
  echo "The tests were successful! 🥰"
else
  echo "The tests failed! 😡" >&2
  exit 1
fi
go run main.go