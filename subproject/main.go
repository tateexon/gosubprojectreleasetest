package main

import (
	"fmt"

	"github.com/tateexon/gosubprojectreleasetest/somesrc"
)

func main() {
	fmt.Println("Hello, World from subpackage!")
	somesrc.SomeFunc()

}
