package gologspace

import "math"

func Generate(min float64, max float64, points uint) []float64 {
	var (
		base             = math.E
		logMin           = math.Log(min)
		logMax           = math.Log(max)
		delta            = (logMax - logMin) / float64(points)
		accDelta float64 = 0
		v        []float64
	)
	for i := 0; i < int(points); i++ {
		v = append(v, math.Pow(base, logMin+accDelta))
		accDelta += delta
	}
	return v
}
