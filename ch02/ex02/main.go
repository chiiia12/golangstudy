// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 43.
//!+

// Cf converts its numeric argument to Celsius and Fahrenheit.
package main

import (
	"os"
	"strconv"
	"fmt"
	"gopl.io/ch2/tempconv"
	"./weightconv"
)

func main() {
	for _, arg := range os.Args[1:] {
		t, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "cf: %v\n", err)
			os.Exit(1)
		}
		f := tempconv.Fahrenheit(t)
		c := tempconv.Celsius(t)
		k := weightconv.Kilo(t)
		p := weightconv.Pond(t)
		fmt.Printf("%s = %s, %s = %s,%s = %s,%s = %s\n",
			f, tempconv.FToC(f), c, tempconv.CToF(c), k, weightconv.KiloToPond(k), p, weightconv.PondToKilo(p))
	}
}

//!-
