package utils

import (
	"bufio"
	"fmt"
	"os"
)

func Input(promt string) string {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print(promt)
	if !scanner.Scan() {
		if err := scanner.Err(); err != nil {
			panic(err)
		}
		panic("Input dihentikan")
	}
	return scanner.Text()
}
