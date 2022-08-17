package main

import (
	"flag"
	"io/ioutil"
	"os"
	"text/template"
)

func main() {
	templateFile := flag.String("t", "template.html", "template file")
	bodyFile := flag.String("b", "body.html", "body file")
	outputFile := flag.String("o", "output.html", "output file")
	flag.Parse()

	templateText, err := ioutil.ReadFile(*templateFile)
	if err != nil {
		panic(err)
	}
	bodyText, err := ioutil.ReadFile(*bodyFile)
	if err != nil {
		panic(err)
	}

	tmpl, err := template.New("test").Parse(string(templateText))
	if err != nil {
		panic(err)
	}
	output, err := os.Create(*outputFile)
	if err != nil {
		panic(err)
	}
	err = tmpl.Execute(output, string(bodyText))
	if err != nil {
		panic(err)
	}

}
