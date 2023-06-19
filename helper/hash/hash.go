package hash

import (
	"crypto/sha512"
	"encoding/hex"
	"fmt"
	"log"
	"os"
)

func Hash(orderid string, statusCode string, grosAmount string) string {

	log.Println(orderid)
	log.Println(statusCode)
	log.Println(grosAmount)
	log.Println(os.Getenv("MIDTRANS_SERVER_KEY"))

	input := orderid + statusCode + grosAmount + os.Getenv("MIDTRANS_SERVER_KEY")
	inputBytes := []byte(input)
	sha512Hasher := sha512.New()
	sha512Hasher.Write(inputBytes)
	hashedInputBytes := sha512Hasher.Sum(nil)
	fmt.Printf("%x\n", hashedInputBytes)
	hashedInputString := hex.EncodeToString(hashedInputBytes)

	return hashedInputString
}
