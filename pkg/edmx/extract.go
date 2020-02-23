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
		Schemas []Schema `xml:"Schema"`
	} `xml:"DataServices"`
}

// Schema ...
type Schema struct {
	Namespace       string          `xml:"Namespace,attr"`
	ComplexTypes    []ComplexType   `xml:"ComplexType"`
	EntityTypes     []EntityType    `xml:"EntityType"`
	Associations    []Association   `xml:"Association"`
	EntityContainer EntityContainer `xml:"EntityContainer"`
	Annotations     []Annotation    `xml:"Annotations"`
}

// ComplexType ...
type ComplexType struct {
	Name       string     `xml:"Name,attr"`
	Properties []Property `xml:"Property"`
}

// Property ...
type Property struct {
	Name     string `xml:"Name,attr"`
	Type     string `xml:"Type,attr"`
	Nullable bool   `xml:"Nullable,attr"`
}

// EntityType ...
type EntityType struct {
	Name                 string               `xml:"Name,attr"`
	BaseType             string               `xml:"BaseType,attr"`
	Abstract             bool                 `xml:"Abstract,attr"`
	OpenType             bool                 `xml:"OpenType,attr"`
	Properties           []Property           `xml:"Property"`
	NavigationProperties []NavigationProperty `xml:"NavigationProperty"`
}

// NavigationProperty ...
type NavigationProperty struct {
	Name         string `xml:"Name,attr"`
	Relationship string `xml:"Relationship,attr"`
	Nullable     bool   `xml:"Nullable,attr"`
	ToRole       string `xml:"ToRole,attr"`
	FromRole     string `xml:"FromRole,attr"`
}

// Association ...
type Association struct {
	Name string `xml:"Name,attr"`
	Ends []struct {
		Type         string `xml:"Type,attr"`
		Role         string `xml:"Role,attr"`
		Multiplicity string `xml:"Multiplicity,attr"`
	} `xml:"End"`
}

// EntityContainer ...
type EntityContainer struct {
	EntitySets      []EntitySet      `xml:"EntitySet"`
	FunctionImports []FunctionImport `xml:"FunctionImport"`
	AssociationSets []AssociationSet `xml:"AssociationSet"`
}

// EntitySet ...
type EntitySet struct {
	Name       string `xml:"Name,attr"`
	EntityType string `xml:"EntityType,attr"`
}

// FunctionImport ...
type FunctionImport struct {
	Name         string     `xml:"Name,attr"`
	ReturnType   string     `xml:"ReturnType,attr"`
	IsComposable bool       `xml:"IsComposable,attr"`
	EntitySet    string     `xml:"EntitySet,attr"`
	Parameters   []Property `xml:"Parameter"`
}

// AssociationSet ...
type AssociationSet struct {
	Name        string `xml:"Name,attr"`
	Association string `xml:"Association,attr"`
	Ends        []struct {
		Role      string `xml:"Role,attr"`
		EntitySet string `xml:"EntitySet,attr"`
	} `xml:"End"`
}

// Annotation ...
type Annotation struct {
	Target          string `xml:"Target,attr"`
	ValueAnnotation struct {
		Term string `xml:"Term,attr"`
		Bool bool   `xml:"Bool,attr"`
	} `xml:"ValueAnnotation"`
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
