package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Printf("Usage: %s <sec>\n", os.Args[0])
		return
	}

	var sec int
	_, err := fmt.Sscanf(os.Args[1], "%d", &sec)
	if err != nil {
		panic(err)
	}

	ch := time.After(time.Duration(sec) * time.Second)
	t := <-ch
	fmt.Printf("%v\n", t)
}
