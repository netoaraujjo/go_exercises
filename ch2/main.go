package main

import (
	"bufio"
	"ch2/spaceconv"
	"ch2/tempconv"
	"ch2/weightconv"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var args = os.Args[1:]
	if len(args) > 0 {
		printResult((args))
	} else {
		reader := bufio.NewReader(os.Stdin)
		print("Digite os valores: ")
		texto, _ := reader.ReadString('\n')
		texto = strings.Trim(texto, "\n")
		values := strings.Split(texto, " ")
		fmt.Print(values)
		printResult(values)
	}
}

func printResult(values []string) {
	for _, a := range values {
		v, err := strconv.ParseFloat(a, 64)
		if err != nil {
			continue
		}
		fmt.Printf("%s ", tempconv.CToF(tempconv.Celsius(v)))
		fmt.Printf("%s ", tempconv.CToK(tempconv.Celsius(v)))
		fmt.Printf("%s ", tempconv.FToC(tempconv.Fahrenheit(v)))
		fmt.Printf("%s ", tempconv.FToK(tempconv.Fahrenheit(v)))
		fmt.Printf("%s ", tempconv.KToC(tempconv.Kelvin(v)))
		fmt.Printf("%s ", tempconv.KToF(tempconv.Kelvin(v)))
		fmt.Printf("%s ", spaceconv.MToF(spaceconv.Meters(v)))
		fmt.Printf("%s ", spaceconv.FToM(spaceconv.Foot(v)))
		fmt.Printf("%s ", weightconv.KToL(weightconv.Kilo(v)))
		fmt.Printf("%s\n", weightconv.LToK(weightconv.Libra(v)))
	}
}
