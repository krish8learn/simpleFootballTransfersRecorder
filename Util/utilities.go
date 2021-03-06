package Util

import (
	"math/rand"
	"strings"
	"time"
)

//Need--> This function will be called automatically when the package is first used.
/*Details-->Normally the seed value is often set to the current time.
rand.Seed() expect an int64 as input, we should convert the time to unix nano before passing it to the function.
every time we run the code, the generated values will be different. If we don’t call rand.Seed(),
the random generator will behave like it is seeded by 1, so the generated values will be the same for every run.*/
func init() {
	rand.Seed(time.Now().UnixNano())
}

/*Need--> Generate Random Integer
Details--> min to this expression, the final result will be a random integer between min and max.*/
func RandomIntGenerator(min, max int64) int64 {
	return min + rand.Int63n(max-min+1)
}

/*Need --> Generate Random string*/
const alphabet = "abcdefghijklmopqrstuvwxyzabcdefghijklmopqrstuvwxyzabcdefghijklmopqrstuvwxyz"

func RandomStringGenerator(n int) string {
	length := len(alphabet)
	var sb strings.Builder

	for i := 0; i < n; i++ {
		c := alphabet[rand.Intn(length)]
		sb.WriteByte(c)
	}

	return sb.String()
}
