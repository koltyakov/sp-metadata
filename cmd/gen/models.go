package main

import (
	"fmt"
	"log"
	"os"

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
