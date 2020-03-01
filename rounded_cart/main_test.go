package rounded_cart

import (
	"sort"
	"testing"

	"github.com/magiconair/properties/assert"
)

func downSum(arr []float64) float64 {
	s := float64(0)
	for i := 0; i < len(arr); i++ {
		down := arr[i] - float64(int(arr[i]))
		s += down
	}
	return s
}

func upSum(arr []float64) float64 {
	s := float64(0)
	for i := 0; i < len(arr); i++ {
		up := float64(int(arr[i])+1) - arr[i]
		s += up
	}
	return s
}

func Solution(prices []float64) []int {
	decimalPart := make([]float64, 0)
	total := float64(0)
	for i := 0; i < len(prices); i++ {
		total += prices[i]
		decimal := prices[i] - float64(int(prices[i]))
		if decimal != 0 {
			decimalPart = append(decimalPart, decimal)
		}
	}
	sort.Float64s(decimalPart)
	up := total - float64(int(total))
	variant := float64(0)
	if up >= 0.5 {
		variant = float64(int(total)+1) - total
	} else {
		variant = total - float64(int(total))
	}
	result := make([]int, len(prices))

	median := len(decimalPart) / 2
	down := downSum(decimalPart[0:median])
	up = upSum(decimalPart[median:])
	change := up - down - variant
	for change == 0 {
		if change > 0 {
			median += 1
		} else {
			median -= 1
		}

		if median == len(decimalPart) {
			median--
			break
		}
		if median == -1 {
			median++
			break
		}

		down = downSum(decimalPart[0:median])
		up = upSum(decimalPart[median:])
		change = up - down - variant
	}

	for i := 0; i < len(result); i++ {
		down := prices[i] - float64(int(prices[i]))
		if down == 0 {
			result[i] = int(prices[i])
			continue
		}
		if down >= decimalPart[median] {
			result[i] = int(prices[i]) + 1

		} else {
			result[i] = int(prices[i])
		}
	}

	return result
}

func TestCase1(t *testing.T) {
	assert.Equal(t, Solution([]float64{5.4, 3.3, 5}), []int{6, 3, 5})
}
