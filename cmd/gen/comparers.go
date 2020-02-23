package main

import (
	"fmt"
	"sort"
	"strings"

	"github.com/koltyakov/sp-metadata/config"
)

func rootNamespacesTable(models []*ModelMeta) string {
	envCodes := config.GetEnvironmentsCodes()
	var compareMatrix []*ComparisonVector

	// Namespaces in platform versions
	keyPresenceMap := map[string]map[string]bool{}
	for _, m := range models {
		for _, s := range m.Model.Services.Schemas {
			n, ok := keyPresenceMap[s.Namespace]
			if !ok {
				n = map[string]bool{}
			}
			n[m.Environment.Code] = true
			keyPresenceMap[s.Namespace] = n
		}
	}

	if len(keyPresenceMap) == 0 {
		return ""
	}

	// Map to comparison matrix
	for _, key := range getNamespaces(models) {
		compareMatrix = append(compareMatrix, &ComparisonVector{
			Name:     key,
			Link:     "./" + key,
			Presence: keyPresenceMap[key],
		})
	}

	// Constructing MD table
	return genMDTable("Namespace", envCodes, compareMatrix)
}

func complexTypesTable(models []*ModelMeta, namespace string) string {
	envCodes := config.GetEnvironmentsCodes()
	var compareMatrix []*ComparisonVector

	// Namespaces in platform versions
	keyPresenceMap := map[string]map[string]bool{}
	for _, m := range models {
		for _, s := range m.Model.Services.Schemas {
			if s.Namespace == namespace {
				for _, ct := range s.ComplexTypes {
					n, ok := keyPresenceMap[ct.Name]
					if !ok {
						n = map[string]bool{}
					}
					n[m.Environment.Code] = true
					keyPresenceMap[ct.Name] = n
				}
			}
		}
	}

	if len(keyPresenceMap) == 0 {
		return ""
	}

	// Map to comparison matrix
	for _, key := range getComplexTypes(models, namespace) {
		compareMatrix = append(compareMatrix, &ComparisonVector{
			Name:     key,
			Presence: keyPresenceMap[key],
		})
	}

	// Constructing MD table
	return genMDTable("Complex Type", envCodes, compareMatrix)
}

func entityTypesTable(models []*ModelMeta, namespace string) string {
	envCodes := config.GetEnvironmentsCodes()
	var compareMatrix []*ComparisonVector

	// Namespaces in platform versions
	keyPresenceMap := map[string]map[string]bool{}
	for _, m := range models {
		for _, s := range m.Model.Services.Schemas {
			if s.Namespace == namespace {
				for _, ent := range s.EntityTypes {
					n, ok := keyPresenceMap[ent.Name]
					if !ok {
						n = map[string]bool{}
					}
					n[m.Environment.Code] = true
					keyPresenceMap[ent.Name] = n
				}
			}
		}
	}

	if len(keyPresenceMap) == 0 {
		return ""
	}

	// Map to comparison matrix
	for _, key := range getEntityTypes(models, namespace) {
		compareMatrix = append(compareMatrix, &ComparisonVector{
			Name:     key,
			Link:     "./EntityTypes/" + key + ".md",
			Presence: keyPresenceMap[key],
		})
	}

	// Constructing MD table
	return genMDTable("Entity Type", envCodes, compareMatrix)
}

func entitySetsTable(models []*ModelMeta, namespace string) string {
	envCodes := config.GetEnvironmentsCodes()
	var compareMatrix []*ComparisonVector

	// Namespaces in platform versions
	keyPresenceMap := map[string]map[string]bool{}
	for _, m := range models {
		for _, s := range m.Model.Services.Schemas {
			if s.Namespace == namespace {
				for _, ent := range s.EntityContainer.EntitySets {
					n, ok := keyPresenceMap[ent.Name]
					if !ok {
						n = map[string]bool{}
					}
					n[m.Environment.Code] = true
					keyPresenceMap[ent.Name] = n
				}
			}
		}
	}

	if len(keyPresenceMap) == 0 {
		return ""
	}

	// Map to comparison matrix
	for _, key := range getEntitySets(models, namespace) {
		compareMatrix = append(compareMatrix, &ComparisonVector{
			Name:     key,
			Presence: keyPresenceMap[key],
		})
	}

	// Constructing MD table
	return genMDTable("Entity Set", envCodes, compareMatrix)
}

func functionsImportsTable(models []*ModelMeta, namespace string) string {
	envCodes := config.GetEnvironmentsCodes()
	var compareMatrix []*ComparisonVector

	// Namespaces in platform versions
	keyPresenceMap := map[string]map[string]bool{}
	for _, m := range models {
		for _, s := range m.Model.Services.Schemas {
			if s.Namespace == namespace {
				for _, ent := range s.EntityContainer.FunctionImports {
					n, ok := keyPresenceMap[ent.Name]
					if !ok {
						n = map[string]bool{}
					}
					n[m.Environment.Code] = true
					keyPresenceMap[ent.Name] = n
				}
			}
		}
	}

	if len(keyPresenceMap) == 0 {
		return ""
	}

	// Map to comparison matrix
	for _, key := range getFunctionsImports(models, namespace) {
		name := key
		if len(name) > 50 {
			runes := []rune(name)
			name = fmt.Sprintf("<span title=\"%s\">%s</span>", name, strings.Trim(string(runes[0:50]), "_")+"...")
			if strings.Contains(key, "_") {
				name += " (" + strings.Replace(key, "_", " ", -1) + ")"
			}
		}
		compareMatrix = append(compareMatrix, &ComparisonVector{
			Name:     name,
			Presence: keyPresenceMap[key],
		})
	}

	// Constructing MD table
	return genMDTable("Functions Imports", envCodes, compareMatrix)
}

func entityTypePropsTable(models []*ModelMeta, namespace string, entityType string) string {
	envCodes := config.GetEnvironmentsCodes()
	var compareMatrix []*ComparisonVector

	// Namespaces in platform versions
	keyPresenceMap := map[string]map[string]bool{}
	for _, m := range models {
		for _, s := range m.Model.Services.Schemas {
			if s.Namespace == namespace {
				for _, ent := range s.EntityTypes {
					if ent.Name == entityType {
						for _, prop := range ent.Properties {
							key := prop.Name + " (" + prop.Type + ")"
							n, ok := keyPresenceMap[key]
							if !ok {
								n = map[string]bool{}
							}
							n[m.Environment.Code] = true
							keyPresenceMap[key] = n
						}
					}
				}
			}
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
		compareMatrix = append(compareMatrix, &ComparisonVector{
			Name:     name,
			Presence: keyPresenceMap[key],
		})
	}

	// Constructing MD table
	return genMDTable("Property", envCodes, compareMatrix)
}

func entityTypeNavPropsTable(models []*ModelMeta, namespace string, entityType string) string {
	envCodes := config.GetEnvironmentsCodes()
	var compareMatrix []*ComparisonVector

	// Namespaces in platform versions
	keyPresenceMap := map[string]map[string]bool{}
	for _, m := range models {
		for _, s := range m.Model.Services.Schemas {
			if s.Namespace == namespace {
				for _, ent := range s.EntityTypes {
					if ent.Name == entityType {
						for _, prop := range ent.NavigationProperties {
							n, ok := keyPresenceMap[prop.Name]
							if !ok {
								n = map[string]bool{}
							}
							n[m.Environment.Code] = true
							keyPresenceMap[prop.Name] = n
						}
					}
				}
			}
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
		compareMatrix = append(compareMatrix, &ComparisonVector{
			Name:     name,
			Presence: keyPresenceMap[key],
		})
	}

	// Constructing MD table
	return genMDTable("Navigation Property", envCodes, compareMatrix)
}

func functionPropsTable(models []*ModelMeta, namespace string, functionImport string) string {
	envCodes := config.GetEnvironmentsCodes()
	var compareMatrix []*ComparisonVector

	// Namespaces in platform versions
	keyPresenceMap := map[string]map[string]bool{}
	for _, m := range models {
		for _, s := range m.Model.Services.Schemas {
			if s.Namespace == namespace {
				for _, functionImp := range s.EntityContainer.FunctionImports {
					if functionImp.Name == functionImport {
						for _, par := range functionImp.Parameters {
							n, ok := keyPresenceMap[par.Name]
							if !ok {
								n = map[string]bool{}
							}
							n[m.Environment.Code] = true
							keyPresenceMap[par.Name] = n
						}
					}
				}
			}
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
		compareMatrix = append(compareMatrix, &ComparisonVector{
			Name:     name,
			Presence: keyPresenceMap[key],
		})
	}

	// Constructing MD table
	return genMDTable("Navigation Property", envCodes, compareMatrix)
}
