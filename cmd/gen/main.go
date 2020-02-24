package main

import (
	"fmt"
	"log"
	"path/filepath"

	"github.com/koltyakov/sp-metadata/pkg/edmx"
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

			// Complex types drill-down
			for _, complexType := range complexTypes {
				propsTable := complexTypePropsTable(models, namespace, complexType)
				if propsTable != "" {
					data := &struct {
						Table       string
						Namespace   string
						ComplexType string
					}{
						Table:       propsTable,
						Namespace:   namespace,
						ComplexType: complexType,
					}
					filePath := filepath.Join(folderPath, "ComplexTypes", complexType+".md")
					if err := genDoc("./pkg/templates/ComplexType.md", filePath, data); err != nil {
						log.Fatal(err)
					}
				}
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

			// Entity types drill-down
			for _, entityType := range entityTypes {
				propsTable := entityTypePropsTable(models, namespace, entityType)
				navPropsTable := entityTypeNavPropsTable(models, namespace, entityType)
				if propsTable != "" || navPropsTable != "" {
					data := &struct {
						PropsTable    string
						NavPropsTable string
						Namespace     string
						EntityType    string
					}{
						PropsTable:    propsTable,
						NavPropsTable: navPropsTable,
						Namespace:     namespace,
						EntityType:    entityType,
					}
					filePath := filepath.Join(folderPath, "EntityTypes", entityType+".md")
					if err := genDoc("./pkg/templates/EntityType.md", filePath, data); err != nil {
						log.Fatal(err)
					}
				}
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

	// Functions imports comparison
	fnImps := getFunctionsImports(models)
	for namespace, envFnMap := range fnImps {
		folderPath := filepath.Join("./docs", namespace)
		table = functionsImportsTable(envFnMap)
		if table != "" {
			// Generating Functions.md
			data := &struct {
				Table     string
				Namespace string
			}{table, namespace}
			filePath := filepath.Join(folderPath, "Functions.md")
			if err := genDoc("./pkg/templates/Functions.md", filePath, data); err != nil {
				log.Fatal(err)
			}

			// Functions imports drill-down
			for key, functionImport := range getAllFunctionsImports(envFnMap) {
				var functionDef edmx.FunctionImport
				paramsTable := functionParamsTable(envFnMap, key)
				if paramsTable == "" {
					continue
				}
				data := &struct {
					ParamsTable  string
					Namespace    string
					EntityType   string
					Function     string
					ReturnType   string
					IsComposable bool
					IsBindable   bool
					EntitySet    string
				}{
					ParamsTable:  paramsTable,
					Namespace:    namespace,
					EntityType:   getFuncImportThisType(functionImport),
					Function:     functionImport.Name,
					ReturnType:   functionDef.ReturnType,
					IsComposable: functionDef.IsComposable,
					IsBindable:   functionDef.IsBindable,
					EntitySet:    functionDef.EntitySet,
				}
				filePath := filepath.Join(folderPath, "Functions", key+".md")
				if err := genDoc("./pkg/templates/Function.md", filePath, data); err != nil {
					log.Fatal(err)
				}
			}
		}
	}

	fmt.Println("Done")
}
