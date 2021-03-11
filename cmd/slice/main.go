package main

import (
	_ "embed"
	"flag"
	"os"
	"text/template"
)

// https://play.golang.org/p/Jn4XQgcOEsk

type data struct {
	Type string
	Name string
}

//go:embed slice.tmpl
var sliceTmpl string

func main() {

	var d data
	flag.StringVar(&d.Type, "type", "", "The subtype used for the slice being generated")
	flag.StringVar(&d.Name, "name", "", "The name used for the slice being generated.")

	flag.Parse()

	t := template.Must(template.New("slice").Parse(sliceTmpl))
	t.Execute(os.Stdout, d)

	// TODO: Create directory for newly generated code
	// TODO: Clean up cmd line flag parsing
}
