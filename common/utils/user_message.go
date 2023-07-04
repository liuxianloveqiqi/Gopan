package utils

import (
	"fmt"
	"math/rand"
	"time"
)

func RandNickname() string {
	rand.Seed(time.Now().UnixNano())
	num := rand.Intn(100000000)
	return fmt.Sprintf("user%08d", num)
}
