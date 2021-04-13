package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/hashicorp/hcl/hcl/printer"
	jsonParser "github.com/hashicorp/hcl/json/parser"
)


func main() {

        in_files := os.Args[1:]
        if len(in_files) < 1{
            fmt.Println("No files detected!!!")
            os.Exit(1)
        }        
        input, err := ioutil.ReadFile(in_files[0])
	if err != nil {
		fmt.Errorf("unable to read from stdin: %s", err)
	}

	ast, err := jsonParser.Parse([]byte(input))
	if err != nil {
		fmt.Errorf("unable to parse JSON: %s", err)
	}

	err = printer.Fprint(os.Stdout, ast)
	if err != nil {
		fmt.Errorf("unable to print HCL: %s", err)
	} 
        

}
