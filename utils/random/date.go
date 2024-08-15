package random

import (
	"math/rand"
	"time"
)

func RandDate(from time.Time, to time.Time) time.Time {
	delta := to.Unix() - from.Unix()

	sec := rand.Int63n(delta+1) + from.Unix()
	return time.Unix(sec, 0)
}

func RandDateLast30Days() time.Time {
	to := time.Now()
	from := to.Add(-1 * 30 * 24 * time.Hour)

	return RandDate(from, to)
}
