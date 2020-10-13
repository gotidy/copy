# ptr [![GoDoc](https://godoc.org/github.com/gotidy/copy?status.svg)](https://godoc.org/github.com/gotidy/copy) [![Go Report Card](https://goreportcard.com/badge/github.com/gotidy/copy)](https://goreportcard.com/report/github.com/gotidy/copy)

Fast structs copier.

## Installation

`go get github.com/gotidy/copy`

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
dest := Employee{}

c := New("") // New("json")
c.Copy(&src, &dest)
```

## Documentation

[GoDoc](http://godoc.org/github.com/gotidy/copy)

## License

[Apache 2.0](https://github.com/gotidy/copy/blob/master/LICENSE)
