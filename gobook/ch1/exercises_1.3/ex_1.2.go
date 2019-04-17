package main

import (
	"fmt"
	"os"
	"strconv"
)

func main(){
	s, sep := "", " "
	for idx, arg := range os.Args[1:] {
		position := strconv.Itoa(idx)
		s += sep + position + ")"  + arg
		sep = " "
	}

	fmt.Println(s)
}
