package edmx

import (
	"encoding/xml"
	"io"
	"io/ioutil"
)

// Edmx ...
type Edmx struct {
	XMLName  xml.Name `xml:"Edmx"`
	Services struct {
		Schemas []struct {
			Namespace    string `xml:"Namespace,attr"`
			ComplexTypes []struct {
				Name       string `xml:"Name,attr"`
				Properties []struct {
					Name     string `xml:"Name,attr"`
					Type     string `xml:"Type,attr"`
					Nullable bool   `xml:"Nullable,attr"`
				} `xml:"Property"`
			} `xml:"ComplexType"`
			EntityTypes []struct {
				// Key / PropertyRef / Name
				Name       string `xml:"Name,attr"`
				BaseType   string `xml:"BaseType,attr"`
				Abstract   bool   `xml:"Abstract,attr"`
				OpenType   bool   `xml:"OpenType,attr"`
				Properties []struct {
					Name     string `xml:"Name,attr"`
					Type     string `xml:"Type,attr"`
					Nullable bool   `xml:"Nullable,attr"`
				} `xml:"Property"`
				NavigationProperties []struct {
					Name         string `xml:"Name,attr"`
					Relationship string `xml:"Relationship,attr"`
					Nullable     bool   `xml:"Nullable,attr"`
					ToRole       string `xml:"ToRole,attr"`
					FromRole     string `xml:"FromRole,attr"`
				} `xml:"NavigationProperty"`
			} `xml:"EntityType"`
			Associations []struct {
				Name string `xml:"Name,attr"`
				Ends []struct {
					Type         string `xml:"Type,attr"`
					Role         string `xml:"Role,attr"`
					Multiplicity string `xml:"Multiplicity,attr"`
				} `xml:"End"`
			} `xml:"Association"`
			EntityContainer struct {
				EntitySets []struct {
					Name       string `xml:"Name,attr"`
					EntityType string `xml:"EntityType,attr"`
				} `xml:"EntitySet"`
				FunctionImports []struct {
					Name         string `xml:"Name,attr"`
					ReturnType   string `xml:"ReturnType,attr"`
					IsComposable bool   `xml:"IsComposable,attr"`
					EntitySet    string `xml:"EntitySet,attr"`
					Parameters   []struct {
						Name     string `xml:"Name,attr"`
						Type     string `xml:"Type,attr"`
						Nullable bool   `xml:"Nullable,attr"`
					} `xml:"Parameter"`
				} `xml:"FunctionImport"`
				AssociationSets []struct {
					Name        string `xml:"Name,attr"`
					Association string `xml:"Association,attr"`
					Ends        []struct {
						Role      string `xml:"Role,attr"`
						EntitySet string `xml:"EntitySet,attr"`
					} `xml:"End"`
				} `xml:"AssociationSet"`
			} `xml:"EntityContainer"`
			Annotations []struct {
				Target          string `xml:"Target,attr"`
				ValueAnnotation struct {
					Term string `xml:"Term,attr"`
					Bool bool   `xml:"Bool,attr"`
				} `xml:"ValueAnnotation"`
			} `xml:"Annotations"`
		} `xml:"Schema"`
	} `xml:"DataServices"`
}

// Extract extracts metadata from EDMX XML definition
func Extract(xmlReader io.Reader) (*Edmx, error) {
	data, err := ioutil.ReadAll(xmlReader)
	if err != nil {
		return nil, err
	}

	var edmx *Edmx
	if err := xml.Unmarshal(data, &edmx); err != nil {
		return nil, err
	}

	return edmx, nil
}
