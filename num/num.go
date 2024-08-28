package num

import (
	"math"
)

func Ceil(f float64) int64 {
	return int64(math.Ceil(f))
}

func Floor(f float64) int64 {
	return int64(math.Floor(f))
}
