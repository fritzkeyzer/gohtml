test:
    go build ./...
    go test ./...

    # Run the tool on this directory and the example dir
    go run cmd/gohtml/main.go
    (cd example && go run ../cmd/gohtml/main.go)

    # TODO check git diff to check that the output is the same

gen-tests:
    go run cmd/gohtml/main.go

example:
    (cd example && go run ../cmd/gohtml/main.go) # run gohtml on the directory
    (cd example && go run main.go) # run the example main.go