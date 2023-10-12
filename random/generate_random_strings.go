package main

import (
	"flag"
	"fmt"
	"math/rand"
)

func main() {
	limit := flag.Int("c", 10, "number of string to generate")
	kv := flag.Bool("k", false, "generate sequential key, random value tsv data")
	min := flag.Int("m", 10, "min len of string")
	max := flag.Int("M", 35, "max len of string")
	seed := flag.Int("s", 0, "initial seed")

	flag.Parse()

	for i := 0; i < *limit; i++ {

		if *kv {
			fmt.Printf("%d\t%s\n", i, GenerateRandomString(*seed+i, *min, *max))
		} else {
			fmt.Printf("%s\n", GenerateRandomString(*seed+i, *min, *max))
		}
	}
}

// GenerateRandomStrings
func GenerateRandomString(seed int, minlen, maxlen int) string {
	source := rand.NewSource(int64(seed))
	rng := rand.New(source)

	return RandomString(rng, minlen, maxlen)
}

// RandomString
func RandomString(rng *rand.Rand, minlen, maxlen int) string {
	chars := make([]rune, rng.Intn(maxlen-minlen)+minlen)
	for i := range chars {
		chars[i] = rune(rng.Intn('}'-'0') + '0')
	}
	return string(chars)
}
