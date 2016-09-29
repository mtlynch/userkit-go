package utils

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"os"
)

// RandStr returns a crypto random string of given length
func RandStr(length int) string {
	pool := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	poolSizeInt := len(pool)
	poolSizeBig := big.NewInt(int64(poolSizeInt))
	var randBytes = make([]byte, poolSizeInt)
	for i := 0; i < length; i++ {
		randBigInt, err := rand.Int(rand.Reader, poolSizeBig)
		if err != nil {
			panic("userkit: error gettting random number")
		}
		randInt := randBigInt.Int64()
		randBytes[i] = pool[int(randInt)]
	}
	return string(randBytes)
}

// RandEmail returns a random email address for testing purposes
func RandEmail() string {
	return fmt.Sprintf("test.user.%s@userkit.co", RandStr(8))
}

// GetTestKey returns the UserKit api key to use for a unit test
func GetTestKey() string {
	key := os.Getenv("USERKIT_KEY")
	if len(key) == 0 {
		panic("No UserKit API key specified, set USERKIT_KEY env variable.")
	}
	return key
}
