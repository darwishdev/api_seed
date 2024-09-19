package random

import (
	"fmt"
	"math/rand"
	"strings"
	"time"

	"golang.org/x/exp/constraints"
)

const alphabet = "abcdefghijklmnopqrstuvwxyz"
const numbers = "0123456789"

func init() {
	rand.Seed(time.Now().UnixNano())
}

type Number interface {
	constraints.Integer | constraints.Float
}

func RandomNumber[T Number](min, max T) T {
	switch min := any(min).(type) {
	case int:
		currentMin := any(min).(int)
		currentMax := any(max).(int)
		random := rand.Intn(currentMax-currentMin) + currentMin
		return any(random).(T)
	case int32:
		currentMin := any(min).(int32)
		currentMax := any(max).(int32)
		random := rand.Int31n(currentMax-currentMin+1) + currentMin
		return any(random).(T)
	case int64:
		currentMin := any(min).(int64)
		currentMax := any(max).(int64)
		random := rand.Int63n(currentMax-currentMin+1) + currentMin
		return any(random).(T)
	case float32:
		currentMin := any(min).(float32)
		currentMax := any(max).(float32)
		random := rand.Float32()*(currentMax-currentMin) + currentMin
		return any(random).(T)
	case float64:
		currentMin := any(min).(float64)
		currentMax := any(max).(float64)
		random := rand.Float64()*(currentMax-currentMin) + currentMin
		return any(random).(T)
	default:
		panic("Unsupported numeric type")
	}

}

func RandomEmail() string {
	return fmt.Sprintf("%s@email.com", RandomString(6))
}

func RandomString(n int) string {
	var sb strings.Builder

	k := len(alphabet)
	for i := 0; i < n; i++ {
		c := alphabet[rand.Intn(k)]
		sb.WriteByte(c)
	}
	return sb.String()
}
