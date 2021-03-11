package main

import (
	_ "embed"
	"flag"
	"github.com/scottcagno/generators/internal/utils"
	"log"
	"strings"
	"text/template"
)

// https://play.golang.org/p/Jn4XQgcOEsk

type data struct {
	Package string
	Type    string
	Name    string
}

var d data

func init() {
	// setup flags
	flag.StringVar(&d.Package, "pkg", "", "The name used for the package being generated")
	flag.StringVar(&d.Type, "type", "", "The subtype used for the slice being generated")
	flag.StringVar(&d.Name, "name", "", "The name used for the slice being generated.")
}

//go:embed slice.tmpl
var sliceTmpl string

func main() {

	// parse flags
	flag.Parse()

	// compile template
	t := template.Must(template.New("slice").Parse(sliceTmpl))

	// create output directory and file
	f, err := utils.Create(d.Package + "/" + strings.ToLower(d.Name) + ".go")
	if err != nil {
		log.Fatalf("util.Create(...): %v\n", err)
	}

	// execute template and write to file
	if err := t.Execute(f, d); err != nil {
		log.Fatalf("t.Execute(...): %v\n", err)
	}

	// dont forget to close the file
	if err := f.Close(); err != nil {
		log.Fatalf("f.Close(...): %v\n", err)
	}
}
