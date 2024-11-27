package main

import "fmt"

const ebK float64 = 373

func main() {
	tempK := ebK
	tempC := tempK - 273

	fmt.Printf("A temperatura de ebulição da água em K é %g e em °C é %g.", tempK, tempC)
}
