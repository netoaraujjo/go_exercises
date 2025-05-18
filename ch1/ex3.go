// Compara versões
package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	t1 := time.Now()
	echo_with_for()
	t2 := time.Now()
	echo_with_join()
	t3 := time.Now()
	time_with_for := t2.Sub(t1)
	time_with_join := t3.Sub(t2)
	fmt.Printf("echo_with_for: %v\necho_with_join: %v\n", time_with_for, time_with_join)
}

func echo_with_for() {
	var s, sep string
	for i := 0; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)
}

func echo_with_join() {
	fmt.Println(strings.Join(os.Args[1:], " "))
}
