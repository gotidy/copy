# Package for fast copying structs of different types

[![GoDev](https://img.shields.io/static/v1?label=godev&message=reference&color=00add8)][godev] [![Go Report Card][goreport]][goreport]

[godev]: https://pkg.go.dev/github.com/google/go-cmp/cmp
[goreport]: https://goreportcard.com/badge/github.com/gotidy/copy

This package is meant to make copying of structs to/from others structs a bit easier.

## Installation

```sh
go get -u github.com/gotidy/copy
```

## Example

```go
type User struct {
    Name string
    MiddleName *string
    Surname string
    Email  string
    Age int
    Married  bool
}

type Employee struct {
    Name string
    MiddleName string
    Surname string
    Email  string
    Age int
}

src := User{
    Name:  "John",
    MiddleName: nil,
    Surname: "Smith",
    Email:"john.smith@joy.me",
    Age: 33,
    Married: false,
}
dst := Employee{}

copiers := New() // New("json")
copiers.Copy(&dst, &src)

// Or more fast use case is to create the type specific copier.

copier := copiers.Get(&dst, &src)
copier.Copy(&dst, &src)

```

See the [documentation][godev] for more information.

## License

[Apache 2.0](https://github.com/gotidy/copy/blob/master/LICENSE)
