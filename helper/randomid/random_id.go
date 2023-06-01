package randomid

import (
	"strconv"

	"github.com/google/uuid"
)

func GenerateRandomNumber() uint {
	uuidString := uuid.New().String()
	digits := ""
	for _, char := range uuidString {
		if char >= '0' && char <= '9' {
			digits += string(char)
		}
	}

	numDigits := 11
	if len(digits) < numDigits {
		numDigits = len(digits)
	}
	randomNumberAsString := digits[:numDigits]
	randomNumber, _ := strconv.Atoi(randomNumberAsString)
	return uint(randomNumber)
}
