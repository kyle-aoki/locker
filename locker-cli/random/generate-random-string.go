package random

import (
	"crypto/rand"
	"fmt"
	"log"
	"math/big"
	"strconv"
)

const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"

func GenerateRandomString(args []string) {
	length := args[1]

	n, err := strconv.ParseInt(length, 10, 64)
	if err != nil {
    log.Fatal("Must provide a length for random string")
	}

	randomBytes := make([]byte, n)

	for i := 0; i < int(n); i++ {
		num, err := rand.Int(rand.Reader, big.NewInt(int64(len(letters))))
		if err != nil {
			log.Fatal("Something went wrong generating a random string.")
		}
		randomBytes[i] = letters[num.Int64()]
	}

	fmt.Print(string(randomBytes))
}
