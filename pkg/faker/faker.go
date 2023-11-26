package faker

import (
	"math"
	"math/rand"
	"time"
)

// RandFloat get a random float64
func RandFloat(min, max float64) float64 {
	randValue := min + rand.Float64()*(max-min)
	decimalMultiplier := math.Pow(10, float64(2))
	return math.Round(randValue*decimalMultiplier) / decimalMultiplier
}

func RandInt(min, max int) int {
	return rand.Intn(max-min+1) + min
}

// RandDate get a random date
func RandDate() time.Time {
	minDate := time.Date(2019, 1, 1, 0, 0, 0, 0, time.UTC).Unix()
	maxDate := time.Date(2023, 11, 1, 0, 0, 0, 0, time.UTC).Unix()
	delta := maxDate - minDate

	sec := rand.Int63n(delta) + minDate
	return time.Unix(sec, 0)
}

// RandPastTime get a random time
func RandPastTime(min, max int) time.Time {
	randomDuration := time.Duration(RandInt(min, max)) * time.Minute

	return time.Now().Add(-randomDuration)
}

// RandBool get a 50-50 chance
func RandBool() bool {
	return rand.Intn(2) == 0
}

// RandString get a random string
func RandString(n int) string {
	var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}
