package main

import (
	"fmt"
)

func main() {
	var a int32 = 42
	var b *int32 = &a
	var c int32 = a
	var d *int32 = &c
	fmt.Println(b, d)
}
