REM @echo off
del /q go.sum
go mod tidy
docker compose up --build