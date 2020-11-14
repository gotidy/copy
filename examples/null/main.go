package main

import (
	"database/sql"
	"encoding/json"
	"fmt"

	"github.com/gotidy/copy"
	"github.com/gotidy/ptr"
)

func main() {
	src := struct {
		Name       sql.NullString
		MiddleName sql.NullString
		Surname    sql.NullString
	}{
		Name:       sql.NullString{String: "John", Valid: true},
		MiddleName: sql.NullString{},
		Surname:    sql.NullString{String: "Kennedy", Valid: true},
	}

	dst := struct {
		Name       *string
		MiddleName *string
		Surname    *string
	}{}

	copiers := copy.New()
	copier := copiers.Get(&dst, &src)
	copier.Copy(&dst, &src)

	if data, err := json.MarshalIndent(dst, "", "    "); err == nil {
		fmt.Println(string(data))
	}

	dst.MiddleName = ptr.String("Fitzgerald")
	copier = copiers.Get(&src, &dst)
	copier.Copy(&src, &dst)

	if data, err := json.MarshalIndent(dst, "", "    "); err == nil {
		fmt.Println(string(data))
	}
}
