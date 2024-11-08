test:
    go test ./... # includes golden file comparison

    # Run the tool on this directory and the example dir
    go run cmd/gohtml/main.go
    (cd example && go run ../cmd/gohtml/main.go)

    go build ./... # check that resulting code is valid
    go test ./... # run tests again to ensure nothing has changed

gen-tests:
    go run cmd/gohtml/main.go

example:
    (cd example && go run ../cmd/gohtml/main.go) # run gohtml on the directory
    (cd example && go run main.go) # run the example main.go