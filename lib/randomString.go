package lib

import (
	"math/rand"
)

// Define the valid characters that can be used in the random string
const validChars = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"

// GenerateRandomString generates a random string of a specified length
func GenerateRandomString(length int) string {

	// Create a byte slice of the specified length
	randomBytes := make([]byte, length)

	// Fill the byte slice with random characters
	for i := range randomBytes {
		randomBytes[i] = validChars[rand.Intn(len(validChars))]
	}

	// Convert the byte slice to a string and return it
	return string(randomBytes)
}
