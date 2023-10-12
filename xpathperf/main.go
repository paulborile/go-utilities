package main

import (
	"flag"
)

func main() {

	xmlFileName := flag.String("x", "", "Input XML File")
	flag.Parse()

	w := DecodeXML(*xmlFileName)
	bytes := EncodeXML(w)
	WriteXML(bytes, "output.xml")
}
