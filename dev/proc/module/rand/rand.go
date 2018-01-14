package rand

import (
	"math/rand"
	"time"
)

func Randomize() {
	rand.Seed(int64(time.Now().Nanosecond()))
}

func Int() int {
	return int(rand.Int63())
}

func IntR() int {
	Randomize()
	return Int()
}
