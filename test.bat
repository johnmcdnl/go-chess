cls
go vet ./...
goimports -w ./chess
go test -v ./...
go test -v --race ./...
go run main.go