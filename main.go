package main

import (
	"github.com/m9gky/statistics/statistics"
)

func main() {
	nums := []float64{1, 2, 3, 4, 5}
	fact := 5
	println(statistics.Avg(nums))
	println(statistics.Fact(fact))
	println(statistics.Distance(statistics.Point{X: 1, Y: 1}, statistics.Point{X: 4, Y: 5}))

}
