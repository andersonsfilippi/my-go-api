#!/bin/bash
rm -f go.sum
go mod tidy
docker compose up --build