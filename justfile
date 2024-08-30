test:
    go run cmd/gohtml/main.go
    (cd example && go run ../cmd/gohtml/main.go)
    # TODO check git diff?

    go test ./...