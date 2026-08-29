package main

import (
	"flag"
	"fmt"

	passgen "github.com/TiJon8/passgen-go/pkg/lib"
)

var (
	mode = flag.Uint("mode", 0, "mode password generation (0 - pseudo, 1 - crypto); default: 0")
	length = flag.Uint64("len", 0, "required: password length")
	seed = flag.Int64("seed", 0, "an optional seed for pseudo mode")
	source = flag.String("source", "", "an optional characters string for password combinations")
	count = flag.Int64("count", 1, "count of generation; default: 1")
)


func pseudoSource(seed int64, source string) passgen.PseudoGenerate {
	return passgen.NewSourcePseudo(seed, source)
}

func GenerateCryptoCLI(length uint64, source string) string {
	rsc := passgen.NewSourceCrypto(length, source)
	return rsc.Generate()
}

func main() {
	flag.Parse()
	if *length == 0 {
		flag.Usage()
		return
	}
	switch *mode {
		case 0:
			pseudoSource := pseudoSource(*seed, *source)
			if *count > 1 {
				for range *count {
					pass := pseudoSource.Generate(*length)
					fmt.Println(pass)
				}
			} else {
				pass := pseudoSource.Generate(*length)
				fmt.Println(pass)
			}
		case 1:
			if *count > 1 {
				for range *count {
					pass := GenerateCryptoCLI(*length, *source)
					fmt.Println(pass)
				}
			} else {
				pass := GenerateCryptoCLI(*length, *source)
				fmt.Println(pass)
			}
	}
}