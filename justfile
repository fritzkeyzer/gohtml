test: gen
    go test ./... # includes golden file comparison
    go build ./... # check that resulting code is valid

gen:
    (cd tests && go run ../cmd/gohtml/main.go generate) # generate tests
    (cd example && go run ../cmd/gohtml/main.go generate) # generate example

example: gen
    go run example/main.go