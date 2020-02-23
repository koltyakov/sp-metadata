package main

import (
	"fmt"
	"log"
	"os"
	"sort"

	"github.com/koltyakov/sp-metadata/config"
	"github.com/koltyakov/sp-metadata/pkg/edmx"
)

// ModelMeta ...
type ModelMeta struct {
	Environment *config.Environment
	Model       *edmx.Edmx
}

// Read EDMX models
func getModels() ([]*ModelMeta, error) {
	var models []*ModelMeta

	for _, e := range config.Environments {
		edmxFilePath := fmt.Sprintf("./meta/%s.xml", e.Code)

		reader, err := os.Open(edmxFilePath)
		if err != nil {
			log.Println(err)
			continue
		}

		m, err := edmx.Extract(reader)
		if err != nil {
			return nil, err
		}

		models = append(models, &ModelMeta{
			Environment: e,
			Model:       m,
		})
	}

	return models, nil
}

func getNamespaces(models []*ModelMeta) []string {
	var keys []string
	for _, m := range models {
		for _, s := range m.Model.Services.Schemas {
			found := false
			for _, nn := range keys {
				if nn == s.Namespace {
					found = true
				}
			}
			if !found {
				keys = append(keys, s.Namespace)
			}
		}
	}
	sort.Strings(keys)
	return keys
}

func getComplexTypes(models []*ModelMeta, namespace string) []string {
	var keys []string
	for _, m := range models {
		for _, s := range m.Model.Services.Schemas {
			if s.Namespace == namespace {
				for _, ct := range s.ComplexTypes {
					found := false
					for _, nn := range keys {
						if nn == ct.Name {
							found = true
						}
					}
					if !found {
						keys = append(keys, ct.Name)
					}
				}
			}
		}
	}
	sort.Strings(keys)
	return keys
}

func getEntityTypes(models []*ModelMeta, namespace string) []string {
	var keys []string
	for _, m := range models {
		for _, s := range m.Model.Services.Schemas {
			if s.Namespace == namespace {
				for _, ent := range s.EntityTypes {
					found := false
					for _, nn := range keys {
						if nn == ent.Name {
							found = true
						}
					}
					if !found {
						keys = append(keys, ent.Name)
					}
				}
			}
		}
	}
	sort.Strings(keys)
	return keys
}
