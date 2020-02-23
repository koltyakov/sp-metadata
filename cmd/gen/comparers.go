package main

import (
	"sort"

	"github.com/koltyakov/sp-metadata/config"
)

func rootNamespacesTable(models []*ModelMeta) string {
	envCodes := config.GetEnvironmentsCodes()
	var compareMatrix []*ComparisonVector

	// Namespaces in platform versions
	keyPresenceMap := map[string]map[string]bool{}
	var keys []string
	for _, m := range models {
		for _, s := range m.Model.Services.Schemas {
			n, ok := keyPresenceMap[s.Namespace]
			if !ok {
				n = map[string]bool{}
			}
			n[m.Environment.Code] = true
			keyPresenceMap[s.Namespace] = n
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

	// Map to comparison matrix
	for _, key := range keys {
		compareMatrix = append(compareMatrix, &ComparisonVector{
			Name:     key,
			Presence: keyPresenceMap[key],
		})
	}

	// Constructing MD table
	return genMDTable("Namespace", envCodes, compareMatrix)
}
