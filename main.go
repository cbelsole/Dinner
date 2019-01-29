package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println(dinner(os.Args[1:]))
}
