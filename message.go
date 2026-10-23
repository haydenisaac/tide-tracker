package main

import "fmt"

func predictedHeightMessage(h float64) string {
	return fmt.Sprintf("The tide height is approximately %g. The tide might be overtopping", h)
}
