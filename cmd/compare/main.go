package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"sort"
	"strings"
	"text/template"
	"unicode/utf8"

	"github.com/koltyakov/sp-metadata/config"
	"github.com/koltyakov/sp-metadata/pkg/edmx"
)

// Metadata ...
type Metadata struct {
	Environment *config.Environment
	Model       *edmx.Edmx
}

func main() {
	var models []*Metadata

	// Reading models
	for _, e := range config.Environments {
		edmxFilePath := fmt.Sprintf("./meta/%s.xml", e.Code)

		reader, err := os.Open(edmxFilePath)
		if err != nil {
			log.Println(err)
			continue
		}

		m, err := edmx.Extract(reader)
		if err != nil {
			log.Fatal(err)
		}

		models = append(models, &Metadata{
			Environment: e,
			Model:       m,
		})
	}

	// Constructing MD table
	table := generateTable(models)

	ns, err := ioutil.ReadFile("./pkg/templates/namespaces.md")
	if err != nil {
		log.Fatal(err)
	}

	t, err := template.New("ns").Parse(fmt.Sprintf("%s", ns))
	if err != nil {
		log.Fatal(err)
	}

	nsData := &struct {
		Table string
	}{
		Table: table,
	}

	var b bytes.Buffer
	if err := t.Execute(&b, nsData); err != nil {
		log.Fatal(err)
	}

	res, err := ioutil.ReadAll(&b)
	if err != nil {
		log.Fatal(err)
	}

	file, _ := os.OpenFile("./docs/namespaces.md", os.O_RDONLY|os.O_CREATE, 0644)
	_ = file.Close()

	if err := ioutil.WriteFile("./docs/namespaces.md", res, 0644); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Done")
}

func generateTable(models []*Metadata) string {
	// Namespaces in platform versions
	namespaces := map[string]map[string]bool{}
	var namespacesKeys []string
	for _, m := range models {
		for _, s := range m.Model.Services.Schemas {
			n, ok := namespaces[s.Namespace]
			if !ok {
				n = map[string]bool{}
			}
			n[m.Environment.Code] = true
			namespaces[s.Namespace] = n
			found := false
			for _, nn := range namespacesKeys {
				if nn == s.Namespace {
					found = true
				}
			}
			if !found {
				namespacesKeys = append(namespacesKeys, s.Namespace)
			}
		}
	}
	sort.Strings(namespacesKeys)

	// Constructing MD table
	table := "Namespace"
	for _, m := range models {
		if m.Environment.Code != "spo.target" {
			table += " | " + m.Environment.Name
		}
	}
	table += "\n---------"
	for _, m := range models {
		if m.Environment.Code == "spo.target" {
			continue
		}
		table += "-|-" + strings.Repeat("-", utf8.RuneCountInString(m.Environment.Name))
	}
	table += "\n"

	for _, namespace := range namespacesKeys {
		table += namespace
		presence := namespaces[namespace]
		for i, m := range models {
			status := "✖"
			if presence[m.Environment.Code] {
				status = "✔"
			}
			if m.Environment.Code == "spo.target" && i > 0 && models[i-1].Environment.Code == "spo" {
				if presence["spo"] != presence["spo.target"] {
					table += " (" + status + ")"
				}
				continue
			}
			table += " | " + status
		}
		table += "\n"
	}
	return table
}
