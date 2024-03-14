package main

import (
	"fmt"
	"xagero/go-helper/helper"
)

func main() {
	fmt.Println(helper.IsEmpty(""))
	fmt.Println(helper.IsNotEmpty(" "))

	fmt.Println(helper.IsBlank(""))
	fmt.Println(helper.IsNotBlank(""))
}
