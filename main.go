package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/0xc000022070/passgen/password"
)

func main() {
	passLength, hasSymbols := 32, true

	flag.IntVar(&passLength, "length", passLength, "Password length")
	flag.IntVar(&passLength, "l", passLength, "Shorthand for --length")
	flag.BoolVar(&hasSymbols, "symbols", hasSymbols, "Indicates if the password should contain symbols")
	flag.BoolVar(&hasSymbols, "s", hasSymbols, "Shorthand for --symbols")
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: %s [flags]...\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "Generates a random password from CLI.\n\n")

		fmt.Fprintf(os.Stderr, "Examples:\n")
		fmt.Fprintf(os.Stderr, "%s\n\n", strings.Join([]string{
			fmt.Sprintf("  %s -l=16 -s", os.Args[0]),
			fmt.Sprintf("  %s -l=128 --symbols=false", os.Args[0]),
		}, "\n"))

		fmt.Fprintf(os.Stderr, "Flags:\n")

		flag.VisitAll(func(f *flag.Flag) {
			if len(f.Name) == 1 {
				return
			}

			fmt.Fprintf(os.Stderr, "  -%s, --%s\t%s\n", string(f.Name[0]), f.Name, f.Usage)
		})
	}
	flag.Parse()

	result, err := password.Generate(passLength, hasSymbols)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Fprintf(os.Stdout, "%s\n", result)
}
