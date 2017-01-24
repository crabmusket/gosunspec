package main

import (
	"flag"
	"github.com/crabmusket/gosunspec/layout"
	"github.com/fatih/camelcase"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"strings"
	"text/template"
)

const layout_template = `
package layout

import (
	"github.com/crabmusket/gosunspec/layout"
	"github.com/crabmusket/gosunspec/generators/layout/util"
)

var {{.Layout.Name}} = &layout.RawLayout{
	Devices: []layout.RawDeviceLayout{
	{{range .Layout.Devices}}
		layout.RawDeviceLayout{
			Models: []layout.RawModelLayout{
			{{range .Models}}layout.RawModelLayout{ModelId: {{.ModelId}},Address: util.MakeUint16({{.Address}}),{{if .Repeats}}Repeats: {{.Repeats}},{{end}}},
			{{end}}
			},
		},
	{{end}}
	},
}
`

func go_filename(typename string) string {
	tmp := camelcase.Split(typename)
	for i, s := range tmp {
		tmp[i] = strings.ToLower(s)
	}
	return strings.Join(tmp, "_") + ".go"
}

func main() {

	var layoutDir string

	flag.StringVar(&layoutDir, "layout-dir", "./", "The location of the layout XML directory")
	flag.Parse()

	files, err := ioutil.ReadDir(layoutDir)
	if err != nil {
		log.Fatal(err)
	}

	layouts := []*layout.RawLayout{}
	for _, file := range files {
		if strings.HasPrefix(file.Name(), "layout") && strings.HasSuffix(file.Name(), ".xml") {
			layoutFile, err := os.OpenFile(layoutDir+file.Name(), os.O_RDONLY, 0644)
			if err != nil {
				log.Fatal(err)
			}

			raw, err := layout.FromLayoutXML(layoutFile)
			if err != nil {
				layoutFile.Close()
				log.Fatal(err)
			}
			layoutFile.Close()
			layouts = append(layouts, raw)
		}
	}

	t := template.New("file")
	t.Funcs(map[string]interface{}{})
	modelTemplate := template.Must(t.Parse(layout_template))

	// write out all the model files

	for _, l := range layouts {
		gofile := go_filename(l.Name)
		outputFilename := "./" + gofile + ".tmp"
		outputFile, err := os.OpenFile(outputFilename, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
		if err != nil {
			log.Fatal(err)
		}
		if err := modelTemplate.Execute(outputFile, map[string]interface{}{
			"Layout": l,
		}); err != nil {
			log.Fatalf("template execution failed: %v", err)
		}
		outputFile.Close()
		cmd := exec.Command("/bin/sh", "-c", "gofmt -w "+outputFilename)
		if err := cmd.Run(); err != nil {
			log.Fatalf("gofmt failed: %v", err)
		}
		if err := os.Rename(outputFilename, "./"+gofile); err != nil {
			log.Fatalf("replacing %s failed: %v", gofile, err)
		}
	}

}
