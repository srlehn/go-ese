package main

import (
	"fmt"
	"os"

	kingpin "gopkg.in/alecthomas/kingpin.v2"
	"www.velocidex.com/golang/go-ese/parser"
)

var (
	catalog_command = app.Command(
		"catalog", "Dump the catalog")
	catalog_command_file_arg = catalog_command.Arg(
		"file", "The image file to inspect",
	).Required().OpenFile(os.O_RDONLY, os.FileMode(0666))
)

func doCatalog() {
	ese_ctx, err := parser.NewESEContext(*catalog_command_file_arg)
	kingpin.FatalIfError(err, "Unable to open ese file")

	catalog, err := parser.ReadCatalog(ese_ctx)
	kingpin.FatalIfError(err, "Unable to open ese file")
	fmt.Printf(catalog.Dump())
}

func init() {
	command_handlers = append(command_handlers, func(command string) bool {
		switch command {
		case catalog_command.FullCommand():
			doCatalog()

		default:
			return false
		}
		return true
	})
}
