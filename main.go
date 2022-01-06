package main

import (
	"fmt"
	cal "github.com/feresyan/go-module-calculator"
)

func main() {
	a := cal.SumCal(10.00,10.00)
	fmt.Println(a)

	b := cal.SubCal(10.00, 12.00)
	fmt.Println(b)
}
