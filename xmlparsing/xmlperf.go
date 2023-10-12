package main

import (
	"encoding/xml"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
)

// Device : struct that matches the XML structure
//
//	<device id="generic_ms_nokia_phone_os7_5" user_agent="DO_NOT_MATCH_MS_NOKIA_PHONE_OS7_5" fall_back="generic_ms_phone_os7_5">
type Device struct {
	XMLName   xml.Name `xml:"device"`
	ID        string   `xml:"id,attr"`
	UserAgent string   `xml:"user_agent,attr"`
	FallBack  string   `xml:"fall_back,attr"`
	InnerXML  string   `xml:",innerxml"`
}

// Capability : struct that matches inner device tags
//
//	<capability name="css_border_image" value="webkit"/>
type Capability struct {
	XMLName xml.Name `xml:"capability"`
	Name    string   `xml:"name"`
	Value   string   `xml:"value"`
}

func main() {
	// Open the XML file

	xmlFile := flag.String("x", "", "Input XML File")
	flag.Parse()

	data, err := os.ReadFile(*xmlFile)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	// Create a new XML decoder
	decoder := xml.NewDecoder(strings.NewReader(string(data)))

	for {
		// Read the next XML token
		tok, err := decoder.Token()
		if err == io.EOF {
			break // End of file
		} else if err != nil {
			fmt.Println("Error reading XML token:", err)
			break
		}

		switch elem := tok.(type) {
		case xml.StartElement:
			// Check if the element is the start of the XML fragment
			if elem.Name.Local == "device" {
				var device Device

				// Decode the XML fragment into the 'device' struct
				if err := decoder.DecodeElement(&device, &elem); err != nil {
					fmt.Println("Error decoding XML:", err)
					break
				}
				// fmt.Printf("device id %s, User-Agent %s, fallback %s\n", device.ID, device.UserAgent, device.FallBack)
			}
		}
	}
}
