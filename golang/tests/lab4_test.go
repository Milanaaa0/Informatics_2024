package labs_test

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
	"isuct.ru/informatics2022/labs"
)

//	var tests = []Test_lab4{
//		{in: 1.2, out: 95.65},
//		{in: 1.28, out: 105.92},
//		{in: 1.36, out:116.80},
func TestEquation(t *testing.T) {
	var a = 0.05
	var b = 0.06
	tests := []struct {
		name string
		x    float64
		want float64
	}{
		{"in:", 1.2, 95.65},
		{"in:", 1.28, 105.92},
		{"in:", 1.36, 116.80},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := labs.Calculate_y(tt.x, a, b)
			if math.IsNaN(tt.x) {
				assert.True(t, math.IsNaN(result), "ожидалось NaN %f", result)
			} else {
				assert.InDelta(t, tt.want, result, 0.01, "Ожидалось %f, но получено %f", tt.want, result)
			}
		})
	}
}
