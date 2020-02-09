test:
	go test `go list ./... | grep -v examples` -coverprofile=coverage.txt -covermode=atomic
	go tool cover -func=coverage.out

bench:
	go test -v -benchmem -bench ./...
