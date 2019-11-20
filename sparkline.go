// Sparklines
package spark

import (
	"math"
)

var levels = []rune("▁▂▃▄▅▆▇█")

// lineF64 generates a sparkline string from a slice of float64
func lineF64(ys []float64) string {
	if len(ys) == 0 {
		return ""
	}
	min := math.Inf(1)
	max := math.Inf(-1)
	for _, y := range ys {
		if y < min {
			min = y
		}
		if y > max {
			max = y
		}
	}
	if max == min {
		max = min + 1
	}
	line := make([]rune, len(ys))
	for i := range ys {
		j := int(math.Floor(
			(float64(len(levels)) - 1.) *
				(ys[i] - min) / (max - min)))
		line[i] = levels[j]
	}
	return string(line)
}

// Line generates a sparkline string from a slice of supported numeric type
func Line(ys interface{}) string {
	var yf []float64
	switch ys := ys.(type) {
	case []uint8:
		yf = make([]float64, len(ys))
		var i int
		var v uint8
		for i, v = range ys {
			yf[i] = float64(v)
		}
	case []int8:
		yf = make([]float64, len(ys))
		var i int
		var v int8
		for i, v = range ys {
			yf[i] = float64(v)
		}
	case []uint16:
		yf = make([]float64, len(ys))
		var i int
		var v uint16
		for i, v = range ys {
			yf[i] = float64(v)
		}
	case []int16:
		yf = make([]float64, len(ys))
		var i int
		var v int16
		for i, v = range ys {
			yf[i] = float64(v)
		}
	case []uint32:
		yf = make([]float64, len(ys))
		var i int
		var v uint32
		for i, v = range ys {
			yf[i] = float64(v)
		}
	case []int32:
		yf = make([]float64, len(ys))
		var i int
		var v int32
		for i, v = range ys {
			yf[i] = float64(v)
		}
	case []uint64:
		yf = make([]float64, len(ys))
		var i int
		var v uint64
		for i, v = range ys {
			yf[i] = float64(v)
		}
	case []int64:
		yf = make([]float64, len(ys))
		var i int
		var v int64
		for i, v = range ys {
			yf[i] = float64(v)
		}
	case []uint:
		yf = make([]float64, len(ys))
		var i int
		var v uint
		for i, v = range ys {
			yf[i] = float64(v)
		}
	case []int:
		yf = make([]float64, len(ys))
		var i int
		var v int
		for i, v = range ys {
			yf[i] = float64(v)
		}
	case []float32:
		yf = make([]float64, len(ys))
		var i int
		var v float32
		for i, v = range ys {
			yf[i] = float64(v)
		}
	case []float64:
		yf = ys
	default:
		panic("bad type")
	}
	return lineF64(yf)
}
