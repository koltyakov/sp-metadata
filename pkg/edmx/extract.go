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
			Namespace string `xml:"Namespace,attr"`
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
