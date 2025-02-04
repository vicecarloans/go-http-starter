package main

import (
	"fmt"
	"io"
	"os"

	"ariga.io/atlas-provider-gorm/gormschema"

	"go-http-server/internal/authors"
	"go-http-server/internal/books"
)


func main() {
	stmts, err := gormschema.New("postgres").Load(
		&books.Books{},
		&authors.Authors{},
	)

	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to load gorm schema: %v\n", err)
        os.Exit(1)
	}
	_, err = io.WriteString(os.Stdout, stmts)

	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to write to stdout: %v\n", err)
		os.Exit(1)
	}
}
