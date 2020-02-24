package main

import (
	"fmt"
	"sort"
	"strings"

	"github.com/koltyakov/sp-metadata/config"
	"github.com/koltyakov/sp-metadata/pkg/edmx"
)

func getFunctionsImports(models []*ModelMeta) map[string]map[string][]edmx.FunctionImport {
	namespaces := getNamespaces(models)
	for i, j := 0, len(namespaces)-1; i < j; i, j = i+1, j-1 {
		namespaces[i], namespaces[j] = namespaces[j], namespaces[i]
	} // reverse slice of namespaces

	// map[namespace]map[envCode][]*edmx.FunctionImport
	funcs := map[string]map[string][]edmx.FunctionImport{}
	for _, ns := range namespaces {
		funcs[ns] = map[string][]edmx.FunctionImport{}
	}

	for _, m := range models {
		for _, s := range m.Model.Services.Schemas {
			if s.Namespace == "SP" { // all FunctionsImports are in SP namespace
				for _, fnImp := range s.EntityContainer.FunctionImports {
					contains := false
					for _, ns := range namespaces {
						if contains {
							continue
						}
						thisType := getFuncImportThisType(fnImp)
						if thisType != "" {
							nss := strings.Replace(strings.Replace(thisType, "Collection(", "", 1), ")", "", 1)
							if strings.HasPrefix(nss, ns) {
								contains = true
							}
						}
						if thisType == "" {
							if strings.HasPrefix(fnImp.ReturnType, ns) {
								contains = true
							}
							if !contains && strings.Contains(fnImp.Name, "_") {
								nss := strings.Replace(fnImp.Name, "_", ".", -1)
								if strings.HasPrefix(nss, ns) {
									contains = true
								}
							}
						}
						if contains {
							if _, ok := funcs[ns][m.Environment.Code]; !ok {
								funcs[ns][m.Environment.Code] = []edmx.FunctionImport{}
							}
							funcs[ns][m.Environment.Code] = append(funcs[ns][m.Environment.Code], fnImp)
						}
					}
					if !contains {
						if _, ok := funcs["SP"][m.Environment.Code]; !ok {
							funcs["SP"][m.Environment.Code] = []edmx.FunctionImport{}
						}
						funcs["SP"][m.Environment.Code] = append(funcs["SP"][m.Environment.Code], fnImp)
					}
				}
			}
		}
	}

	return funcs
}

func getFuncImportThisType(fnImp edmx.FunctionImport) string {
	for _, par := range fnImp.Parameters {
		if par.Name == "this" {
			return par.Type
		}
	}
	return ""
}

func getAllFunctionsImports(fnImps map[string][]edmx.FunctionImport) map[string]edmx.FunctionImport {
	funcs := map[string]edmx.FunctionImport{}
	for _, fnImpArr := range fnImps {
		for _, fnImp := range fnImpArr {
			key := fnImp.Name
			thisType := getFuncImportThisType(fnImp)
			if thisType != "" {
				nss := strings.Replace(strings.Replace(thisType, "Collection(", "", 1), ")", "", 1)
				key = nss + "__" + key
			}
			if _, ok := funcs[key]; !ok {
				funcs[key] = fnImp
			}
		}
	}
	return funcs
}

func functionsImportsTable(fnImps map[string][]edmx.FunctionImport) string {
	envCodes := config.GetEnvironmentsCodes()
	var compareMatrix []*ComparisonVector

	keyPresenceMap := map[string]map[string]bool{}
	for envCode, fnImpArr := range fnImps {
		for _, fnImp := range fnImpArr {
			key := fnImp.Name
			thisType := getFuncImportThisType(fnImp)
			if thisType != "" {
				key = key + " (" + thisType + ")"
			}
			n, ok := keyPresenceMap[key] // check for uniqueness
			if !ok {
				n = map[string]bool{}
			}
			n[envCode] = true
			keyPresenceMap[key] = n
		}
	}

	if len(keyPresenceMap) == 0 {
		return ""
	}

	var keys []string
	for key := range keyPresenceMap {
		keys = append(keys, key)
	}
	sort.Strings(keys)

	// Map to comparison matrix
	for _, key := range keys {
		name := key
		if len(strings.Split(name, " ")[0]) > 50 {
			runes := []rune(name)
			name = fmt.Sprintf("<span title=\"%s\">%s</span>", name, strings.Trim(string(runes[0:50]), "_")+"...")
			if strings.Contains(key, "_") {
				name += " (" + strings.Replace(key, "_", " ", -1) + ")"
			}
		}

		nArr := strings.Split(key, " ")
		link := "./Functions/" + nArr[0] + ".md"
		if len(nArr) > 1 {
			ns := strings.Replace(strings.Replace(strings.Replace(nArr[1], "Collection(", "", -1), "(", "", -1), ")", "", -1)
			link = "./Functions/" + ns + "__" + nArr[0] + ".md"
		}

		if functionParamsTable(fnImps, key) == "" {
			link = ""
		}

		// name = fmt.Sprintf("[%s](./Functions/%s.md)", name, key)
		compareMatrix = append(compareMatrix, &ComparisonVector{
			Name:     name,
			Link:     link,
			Presence: keyPresenceMap[key],
		})
	}

	// Constructing MD table
	return genMDTable("Functions Imports", envCodes, compareMatrix)
}

func functionParamsTable(envFnMap map[string][]edmx.FunctionImport, key string) string {
	envCodes := config.GetEnvironmentsCodes()
	var compareMatrix []*ComparisonVector

	keyArr := strings.Split(key, "__")
	fnName := keyArr[0]
	thisType := ""
	if len(keyArr) > 1 {
		fnName = keyArr[1]
		thisType = keyArr[0]
	}

	keyPresenceMap := map[string]map[string]bool{}
	var keys []string
	for envCode, fnArr := range envFnMap {
		for _, functionImp := range fnArr {
			if functionImp.Name == fnName {
				if thisType == getFuncImportThisType(functionImp) {
					for _, par := range functionImp.Parameters {
						if par.Name != "this" {
							key := par.Name + " (" + par.Type + ")"
							n, ok := keyPresenceMap[key]
							if !ok {
								n = map[string]bool{}
							}
							n[envCode] = true
							keyPresenceMap[key] = n
							newKey := true
							for _, k := range keys {
								if k == key {
									newKey = false
								}
							}
							if newKey {
								keys = append(keys, key)
							}
						}
					}
				}
			}
		}
	}

	if len(keyPresenceMap) == 0 {
		return ""
	}

	// Map to comparison matrix
	for _, key := range keys {
		name := key
		compareMatrix = append(compareMatrix, &ComparisonVector{
			Name:     name,
			Presence: keyPresenceMap[key],
		})
	}

	// Constructing MD table
	return genMDTable("Parameter", envCodes, compareMatrix)
}
