package main

import (
	"fmt"
	"github.com/xagero/go-helper/helper"
)

func main() {
	fmt.Println(helper.IsEmpty(""))
	fmt.Println(helper.IsNotEmpty(" "))

	fmt.Println(helper.IsBlank(""))
	fmt.Println(helper.IsNotBlank(""))
}
