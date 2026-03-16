package main

import (
	"fmt"
	"math"
)

func main() {
	var l1, l2, l3, r, area float64
	fmt.Scan(&l1, &l2, &l3)
	r = (l1 + l2 + l3) / 2
	area = math.Sqrt(r * (r - l1) * (r - l2) * (r - l3))
	fmt.Printf("%.2f\n", area)
}
