package gologspace

import "math"

func Generate(min float64, max float64, points uint) []float64 {
	const base = math.E

	var (
		logMin           = math.Log(min)
		logMax           = math.Log(max)
		delta            = (logMax - logMin) / float64(points)
		accDelta float64 = 0
		pts              = make([]float64, 0, points)
	)

	for range points {
		pts = append(pts, math.Pow(base, logMin+accDelta))
		accDelta += delta
	}

	return pts
}
