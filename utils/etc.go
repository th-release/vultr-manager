package utils

import (
	"fmt"
	"math"
	"strconv"
)

func ToFloat64(num int) float64 {
	return float64(num)
}

func ToInt(v interface{}) (int, error) {
	// interface{}를 int로 변환 시도
	i, ok := v.(int)
	if !ok {
		return 0, fmt.Errorf("value is not an int: %v", v)
	}
	return i, nil
}

func CalculateQuantity(
	price float64,
	usdt float64,
	leverage int,
	percent float64,
) float64 {
	investmentRatio := percent / 100
	investmentAmount := usdt * investmentRatio
	leveragedInvestment := investmentAmount * float64(leverage)
	return leveragedInvestment / price
}

func MaxFloat64(numbers ...float64) float64 {
	if len(numbers) == 0 {
		return 0
	}

	maxValue := numbers[0]
	for _, value := range numbers {
		if value > maxValue {
			maxValue = value
		}
	}

	return maxValue
}

func MaxInt(numbers ...int) int {
	if len(numbers) == 0 {
		return 0
	}

	maxValue := numbers[0]
	for _, value := range numbers {
		if value > maxValue {
			maxValue = value
		}
	}

	return maxValue
}

func MinFloat64(a, b float64) float64 {
	if math.IsNaN(a) && math.IsNaN(b) {
		return math.NaN()
	}
	if math.IsNaN(a) {
		return b
	}
	if math.IsNaN(b) {
		return a
	}
	return math.Min(a, b)
}

func MinInt(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// floatToInt는 float64 값을 int로 변환합니다 (소수점 이하 버림).
func FloatToInt(f float64) int {
	return int(f) // 소수점 이하를 버리고 int로 변환
}

// floatToIntWithRound는 float64 값을 반올림하여 int로 변환합니다.
func FloatToIntWithRound(f float64) int {
	return int(math.Round(f)) // 반올림 후 int로 변환
}

func ToFixed(value float64, decimalPlaces int) float64 {
	// decimalPlaces가 음수일 경우 0으로 처리
	if decimalPlaces < 0 {
		decimalPlaces = 0
	}

	// 포맷 문자열 생성 (예: "%.2f" for 2 decimal places)
	format := fmt.Sprintf("%%.%df", decimalPlaces)

	// value를 지정된 소수점 자릿수로 포맷
	formatted := fmt.Sprintf(format, value)

	// 문자열을 float64로 변환
	result, err := strconv.ParseFloat(formatted, 64)
	if err != nil {
		// 에러 발생 시 기본값 반환 (또는 에러 처리 로직 추가)
		return 0
	}

	return result
}

func CalculatePnlByPercent(
	price float64,
	percent float64,
	leverage int,
	positionType string,
) float64 {
	if positionType != "LONG" && positionType != "SHORT" {
		return 0.0
	} else {
		if positionType == "LONG" {
			return price + (price / 100 * (percent / float64(leverage)))
		} else {
			return price - (price / 100 * (percent / float64(leverage)))
		}
	}
}

func PercentageDifference(a float64, b float64) float64 {
	diff := a - b
	abs := math.Abs(diff)
	return abs / a * 100
}
