package main

import (
	"fmt"
	"strings"
	"unicode/utf8"

	"github.com/koltyakov/sp-metadata/config"
)

// ComparisonVector ...
type ComparisonVector struct {
	Name     string
	Link     string
	Presence map[string]bool
}

func genMDTable(parameter string, envCodes []string, compareMatrix []*ComparisonVector) string {
	eMap := config.GetEnvironmentsMap()

	// Constructing MD table
	table := parameter
	for _, envCode := range envCodes {
		if envCode != "spo.target" {
			table += " | " + eMap[envCode].Name
		}
	}
	table += "\n---------"
	for _, envCode := range envCodes {
		if envCode == "spo.target" {
			continue
		}
		table += "-|-" + strings.Repeat("-", utf8.RuneCountInString(eMap[envCode].Name))
	}
	table += "\n"

	for _, vector := range compareMatrix {
		if vector.Link == "" {
			table += vector.Name
		} else {
			table += fmt.Sprintf("[%s](%s)", vector.Name, vector.Link)
		}
		for i, envCode := range envCodes {
			status := "✖"
			if vector.Presence[envCode] {
				status = "✔"
			}
			if envCode == "spo.target" && i > 0 && envCodes[i-1] == "spo" {
				if vector.Presence["spo"] != vector.Presence["spo.target"] {
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
