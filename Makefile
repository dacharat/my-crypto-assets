dev:
	ENV=dev go run cmd/api/main.go

test:
	go test -v ./...

vet:
	go vet ./...

cleantest:
	go clean -testcache
	go test -v ./...
