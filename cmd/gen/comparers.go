package main

import (
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

	// Map to comparison matrix
	for _, key := range getEntityTypes(models, namespace) {
		compareMatrix = append(compareMatrix, &ComparisonVector{
			Name:     key,
			Presence: keyPresenceMap[key],
		})
	}

	// Constructing MD table
	return genMDTable("Entity Type", envCodes, compareMatrix)
}
