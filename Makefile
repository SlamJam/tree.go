test:
	go test -timeout 30s github.com/SlamJam/tree.go

generate:
	go generate ./...
