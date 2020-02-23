package main

import (
	"fmt"
	"log"
	"path/filepath"
)

func main() {
	// Reading models
	models, err := getModels()
	if err != nil {
		log.Fatal(err)
	}

	// Namespaces root comparison
	table := rootNamespacesTable(models)
	data := &struct{ Table string }{table}
	if err := genDoc("./pkg/templates/Namespaces.md", "./docs/Namespaces.md", data); err != nil {
		log.Fatal(err)
	}

	// Namespaces iteration
	for _, namespace := range getNamespaces(models) {
		folderPath := filepath.Join("./docs", namespace)

		// Complex types comparison
		complexTypes := getComplexTypes(models, namespace)
		if len(complexTypes) > 0 {
			table = complexTypesTable(models, namespace)
			data := &struct {
				Table     string
				Namespace string
			}{table, namespace}
			filePath := filepath.Join(folderPath, "ComplexTypes.md")
			if err := genDoc("./pkg/templates/ComplexTypes.md", filePath, data); err != nil {
				log.Fatal(err)
			}
		}

		// Entity types comparison
		entityTypes := getEntityTypes(models, namespace)
		if len(entityTypes) > 0 {
			table = entityTypesTable(models, namespace)
			data := &struct {
				Table     string
				Namespace string
			}{table, namespace}
			filePath := filepath.Join(folderPath, "EntityTypes.md")
			if err := genDoc("./pkg/templates/EntityTypes.md", filePath, data); err != nil {
				log.Fatal(err)
			}
		}

		// Entity sets comparison
		entitySets := getEntitySets(models, namespace)
		if len(entitySets) > 0 {
			table = entitySetsTable(models, namespace)
			data := &struct {
				Table     string
				Namespace string
			}{table, namespace}
			filePath := filepath.Join(folderPath, "EntitySets.md")
			if err := genDoc("./pkg/templates/EntitySets.md", filePath, data); err != nil {
				log.Fatal(err)
			}
		}

	}

	fmt.Println("Done")
}
