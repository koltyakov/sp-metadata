package edmx

import "regexp"

// Clean cleans EDMX from specifics
func Clean(edmx string) string {
	re := []*regexp.Regexp{
		regexp.MustCompile(`(<Schema Namespace="SP\.Data".*</Schema>)`),
		regexp.MustCompile(`(<EntityContainer Name="ListData".*</EntityContainer>)`),
	}
	for _, r := range re {
		edmx = r.ReplaceAllString(edmx, "")
	}
	return edmx
}