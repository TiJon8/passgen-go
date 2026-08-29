package main

import (
	"flag"
	"fmt"

	passgen "github.com/TiJon8/passgen-go/internal"
)

var (
	mode = flag.Uint("mode", 0, "mode password generation")
	length = flag.Uint64("len", 0, "required: password length")
	seed = flag.Int64("seed", 0, "an optional seed for pseudo mode")
	source = flag.String("source", "", "an optional characters string for password combinations")
	count = flag.Int64("count", 1, "count of generation")
)

func GeneratePseudoCLI(length uint64, seed int64, source string) string {
	rsp := passgen.NewSourcePseudo(seed, source)
	return rsp.Generate(length)
}

func GenerateCryptoCLI(length uint64, source string) string {
	rsc := passgen.NewSourceCrypto(length, source)
	return rsc.Generate()
}

func main() {
	fmt.Println("OK")
	flag.Parse()
	fmt.Println(*length)
	if *length == 0 {
		flag.Usage()
		return
	}
	switch *mode {
		case 0:
			if *count > 1 {
				for range *count {
					pass := GeneratePseudoCLI(*length, *seed, *source)
					fmt.Println(pass)
				}
			} else {
				pass := GeneratePseudoCLI(*length, *seed, *source)
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
	

	// for {
	// 	password := make([]byte, length)
	// 	for i := range password {
	// 		password[i] = source[rand.Intn(len(source))]
	// 	}
	// 	fmt.Println(string(password))

	// 	fmt.Printf("Сделать новый? Y/n\n",)
	// 	var input string
	// 	fmt.Scan(&input)
	// 	if input != "Y" && input != "y" {
	// 		return
	// 	}
	// 	fmt.Println("Введи длину:")
	// 	fmt.Scan(&length)
	// }
}