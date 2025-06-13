package main

import (
	"os"

	"github.com/xiahuaxiahua0616/ifonly/cmd/ifonly-apiserver/app"
	_ "go.uber.org/automaxprocs"
)

func main() {
	command := app.NewIfOnlyCommand()
	if err := command.Execute(); err != nil {
		os.Exit(1)
	}
}
