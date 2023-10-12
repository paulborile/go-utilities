package main

import "testing"

func BenchmarkDecodeXMLFull(b *testing.B) {
	for i := 0; i < b.N; i++ {
		DecodeXML("wurfl-full.xml")
	}
}

func BenchmarkEncodeXML(b *testing.B) {
	w := DecodeXML("wurfl-full.xml")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		EncodeXML(w)
	}
}

func BenchmarkWriteXML(b *testing.B) {
	w := DecodeXML("wurfl-full.xml")
	bytes := EncodeXML(w)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		WriteXML(bytes, "bench_output.xml")
	}
}
