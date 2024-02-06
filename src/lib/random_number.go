package lib

import (
	"fmt"
	"math/rand"
)

func RandomNumber(length int) string {
	var result string
	for i := 0; i < length; i++ {
		result += fmt.Sprintf("%d", rand.Intn(10))
	}
	return result
}
