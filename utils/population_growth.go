package utils

import "math"

// PopulationAfterGrowth applies two consecutive years of percentage
// growth to an initial population: first yearOnePercent, then
// yearTwoPercent on the result of the first year. The result is
// rounded to the nearest whole person.
func PopulationAfterGrowth(initial int, yearOnePercent, yearTwoPercent float64) int {
	afterYearOne := float64(initial) * (1 + yearOnePercent/100)
	afterYearTwo := afterYearOne * (1 + yearTwoPercent/100)
	return int(math.Round(afterYearTwo))
}
