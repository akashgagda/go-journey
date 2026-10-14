package main

import "fmt"

func billTotal(price float64, qty int, gst bool) float64 {
	subtotal := price * float64(qty)
	gstVal := subtotal * 0.18

	if gst {
		subtotal += gstVal
	}
	return subtotal
}

func main() {
	fmt.Println(billTotal(12.5, 2, true))
}
