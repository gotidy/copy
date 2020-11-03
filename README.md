# Package for fast copying structs of different types

[![GoDev](https://img.shields.io/static/v1?label=godev&message=reference&color=00add8)][godev] [![Go Report Card](https://goreportcard.com/badge/github.com/gotidy/copy)][goreport]

[godev]: https://pkg.go.dev/github.com/gotidy/copy
[goreport]: https://goreportcard.com/report/github.com/gotidy/copy

This package is meant to make copying of structs to/from others structs a bit easier.

## Installation

```sh
go get -u github.com/gotidy/copy
```

## Example

```go
type Person struct {
    Name       string
    MiddleName *string
    Surname    string
}

type User struct {
    Person
    Email   string
    Age     int8
    Married bool
}

type Employee struct {
    Name       string
    MiddleName string
    Surname    string
    Email      string
    Age        int
}

src := User{
    Person: Person{
        Name:       "John",
        MiddleName: nil,
        Surname:    "Smith",
    },
    Email:   "john.smith@joy.me",
    Age:     33,
    Married: false,
}
dst := Employee{}

copiers := copy.New() // New("json")
copiers.Copy(&dst, &src)

// Or more fast use case is to create the type specific copier.

copier := copiers.Get(&Employee{}, &User{}) // Created once for a pair of types.
copier.Copy(&dst, &src)

```

### [Benchmark](https://github.com/gotidy/copy-bench)

Benchmarks source code can be found [here](https://github.com/gotidy/copy-bench)

```sh
go test -bench=. -benchmem ./...
goos: darwin
goarch: amd64
pkg: github.com/gotidy/copy-bench
BenchmarkManualCopy-12         177310519         6.92 ns/op          0 B/op        0 allocs/op
BenchmarkCopiers-12             13476417         84.1 ns/op          0 B/op        0 allocs/op
BenchmarkCopier-12              40226689         27.5 ns/op          0 B/op        0 allocs/op
BenchmarkJinzhuCopier-12          407480         2711 ns/op       2480 B/op       34 allocs/op
BenchmarkDeepcopier-12            262836         4346 ns/op       4032 B/op       73 allocs/op
PASS
ok      github.com/gotidy/copy-bench    6.922s
```

See the [documentation][godev] for more information.

## License

[Apache 2.0](https://github.com/gotidy/copy/blob/master/LICENSE)
