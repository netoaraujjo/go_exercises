// Exibe os argumentos de linha de comando e seus respectivos índices
package main

import (
	"fmt"
	"os"
)

func main() {
	for i, arg := range os.Args {
		fmt.Printf("%d %s\n", i, arg)
	}
}
