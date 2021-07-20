package main

import (
	"bytes"
	"fmt"

	"github.com/sunmi-OS/gocore/v2/tools/gocore/template"
)

// hero -source=./test/hero/template -extensions=.got,.html,.docker,.md
// hero -source=./tools/gocore/template -extensions=.got,.html,.docker,.md
func main() {

	buffer := new(bytes.Buffer)

	template.FromMain("test", []string{"Api"}, buffer)

	fmt.Println(buffer.String())

	fmt.Println("22")

	fmt.Println(buffer.String())
}