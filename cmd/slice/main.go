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

//go:embed slice.tmpl
var sliceTmpl string

func main() {

	var d data
	flag.StringVar(&d.Package, "pkg", "", "The name used for the package being generated")
	flag.StringVar(&d.Type, "type", "", "The subtype used for the slice being generated")
	flag.StringVar(&d.Name, "name", "", "The name used for the slice being generated.")
	flag.Parse()

	t := template.Must(template.New("slice").Parse(sliceTmpl))
	f, err := utils.Create(d.Package + "/" + strings.ToLower(d.Name) + ".go")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	t.Execute(f, d)

	// TODO: Create directory for newly generated code
	// TODO: Clean up cmd line flag parsing
}
