package utils

// GroupPercentDifference answers: if n1 identical items are
// percentCheaper percent cheaper than one reference item, by what
// percent do n2 such items differ in price from the reference item?
// A positive result means the n2 items are more expensive than the
// reference item; a negative result means they are cheaper.
func GroupPercentDifference(n1 int, percentCheaper float64, n2 int) float64 {
	unitRatio := (1 - percentCheaper/100) / float64(n1)
	groupRatio := unitRatio * float64(n2)
	return (groupRatio - 1) * 100
}
