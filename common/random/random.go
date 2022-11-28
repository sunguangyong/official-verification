package random

import (
	"time"

	"math/rand"
)

func RandomStringArry(arry []string) (val string) {
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	if len(arry) > 0 {
		i := r1.Intn(len(arry))
		val = arry[i]
	}
	return val
}
