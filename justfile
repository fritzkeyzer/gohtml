test:
    go test ./...

    # Run the tool on this directory and the example dir
    go run cmd/gohtml/main.go
    (cd example && go run ../cmd/gohtml/main.go)

    # TODO check git diff to check that the output is the same

example:
    (cd example && go run ../cmd/gohtml/main.go) # run gohtml on the directory
    (cd example && go run main.go) # run the example main.go