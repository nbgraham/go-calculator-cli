Running `go test -coverprofile=cover.out` does not produce any coverage on `internal/multiply` nor `internal/add`

```
PASS
coverage: 77.8% of statements
ok      calculator      0.212s
```

Running `sh ./int-coverage.sh` does

```
usage: calculator OPERATION A B
unknown operation: invalid
strconv.Atoi: parsing "invalid": invalid syntax
strconv.Atoi: parsing "invalid": invalid syntax
sum: 8
product: 40
        calculator              coverage: 100.0% of statements
        calculator/cmd/calculator               coverage: 100.0% of statements
        calculator/internal/add         coverage: 100.0% of statements
        calculator/internal/multiply            coverage: 100.0% of statements
```