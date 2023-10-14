package validate

//go:generate protoc -I=.. --go_out=..  --go_opt=paths=source_relative validate/validate.proto validate/import.proto
