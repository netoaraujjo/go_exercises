// Exibe seus argumentos de linha de comando
package main

import (
	"fmt"
	"os"
	// "strings"
)

func main() {
	// fmt.Println(strings.Join(os.Args[1:], " "))
	fmt.Println(os.Args[1:])
}