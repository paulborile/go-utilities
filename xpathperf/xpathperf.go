package main

import (
	"encoding/xml"
	"fmt"
	"os"
	"strings"

	"golang.org/x/net/html"
)

type Capability struct {
	XMLName xml.Name `xml:"capability"`
	Name    string   `xml:"name,attr"`
	Value   string   `xml:"value,attr"`
}

type Group struct {
	XMLName    xml.Name     `xml:"group"`
	ID         string       `xml:"id,attr"`
	Capability []Capability `xml:"capability"`
}

type Device struct {
	XMLName xml.Name `xml:"device"`
	ID      string   `xml:"id,attr"`
	Group   []Group  `xml:"group"`
}

type Wurfl struct {
	XMLName xml.Name `xml:"wurfl"`
	Version Version  `xml:"version"`
	Devices []Device `xml:"devices>device"`
}

type Version struct {
	Ver         string `xml:"ver"`
	LastUpdated string `xml:"last_updated"`
	OfficialURL string `xml:"official_url"`
	Maintainers struct {
		Maintainer []struct {
			Name     string `xml:"name,attr"`
			Email    string `xml:"email,attr"`
			HomePage string `xml:"home_page,attr"`
		} `xml:"maintainer"`
	} `xml:"maintainers"`
	Authors struct {
		Author struct {
			Name     string `xml:"name,attr"`
			Email    string `xml:"email,attr"`
			HomePage string `xml:"home_page,attr"`
		} `xml:"author"`
	}
	Statement string `xml:"statement"`
}

// DecodeXML : just parse into dom
func DecodeXML(filename string) *Wurfl {
	// Read the XML file
	xmlFile, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error opening XML file:", err)
		return nil
	}
	defer xmlFile.Close()

	// Decode the XML data into the struct
	var wurflData Wurfl
	decoder := xml.NewDecoder(xmlFile)
	err = decoder.Decode(&wurflData)
	if err != nil {
		fmt.Println("Error decoding XML:", err)
		return nil
	}
	return &wurflData
}

// EncodeXML : only the encoding to memory
func EncodeXML(wurflData *Wurfl) []byte {
	// Encode the modified struct back to XML
	outputXML, err := xml.MarshalIndent(wurflData, "", "  ")
	if err != nil {
		fmt.Println("Error encoding XML:", err)
		return []byte{}
	}
	return outputXML
}

// WriteXML : only the file write phase
func WriteXML(outputXML []byte, filename string) {
	// Write the modified XML to a file
	err := os.WriteFile(filename, outputXML, 0644)
	if err != nil {
		fmt.Println("Error writing output XML:", err)
		return
	}
}

func createXPathCondition(allowed []string) string {
	conditions := make([]string, len(allowed))
	for i, capName := range allowed {
		conditions[i] = "(@name='" + capName + "')"
	}
	return strings.Join(conditions, " or ")
}

func getNodeAttr(node *html.Node, attrName string) string {
	for _, attr := range node.Attr {
		if attr.Key == attrName {
			return attr.Val
		}
	}
	return ""
}

var allowedCapabilities = []string{"brand_name", "model_name", "device_os"}
