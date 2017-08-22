#!/usr/bin/env bash

read -r -d '' OPSMAN <<EOF
{
	"address": "https://172.28.61.5",
	"user": "admin",
	"password": "welcome1"
}
EOF

export OPSMAN

go run main.go routes.go
