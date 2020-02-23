package main

import (
	"fmt"
	"log"
)

func main() {
	// Reading models
	models, err := getModels()
	if err != nil {
		log.Fatal(err)
	}

	// Constructing MD table
	table := rootNamespacesTable(models)
	data := &struct{ Table string }{Table: table}
	if err := genDoc("./pkg/templates/namespaces.md", "./docs/namespaces.md", data); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Done")
}
