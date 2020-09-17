package main

import (
	"fmt"
	"github.com/pulumi/pulumi-azure-nextgen/provider/pkg/arm2pulumi"
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	var readFrom io.Reader
	var langs string

	if len(os.Args) < 2 {
		fmt.Printf("Usage: %s [arm-template path] lang\n", os.Args[0])
		fmt.Print("arm-template   Path to arm-template or assumed to be stdin if omitted\n")
		fmt.Print("lang           Comma separated list of languages - e.g. nodejs,python,dotnet,go\n")
		fmt.Println()
		return
	}
	if len(os.Args) == 2 {
		readFrom = os.Stdin
		langs = os.Args[1]
	} else {
		var err error
		filePath := os.Args[1]
		readFrom, err = os.Open(filePath)
		if err != nil {
			log.Fatalf("Failed to read from file: %s: %+v", filePath, err)
		}
		langs = os.Args[2]
	}

	body, diagnostics, err := arm2pulumi.RenderIR(readFrom, nil, nil)
	if err != nil {
		log.Fatalf("Failure rendering IR from template: %+v", err)
	}
	languages := strings.Split(langs, ",")
	programsMap, err := arm2pulumi.RenderPrograms(body, nil, languages)
	if err != nil {
		log.Fatalf("Failure rendering programs: %+v", err)
	}
	for k, v := range programsMap {
		fmt.Printf("Language: %s\n", k)
		fmt.Println()
		fmt.Printf("%s\n", v)
		fmt.Println()
		fmt.Println()
	}
	for k, diags := range diagnostics {
		fmt.Printf("Diagnostics for %s\n", k)
		for _, diag := range diags {
			fmt.Printf("WARN: [%s] at '%s' - '%s'\n", diag.Severity, diag.SourceToken, diag.Description)
		}
		fmt.Println()
	}
	return

}

