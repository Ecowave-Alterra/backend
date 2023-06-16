package customstatus

import (
	"strconv"
	"strings"
)

func CustomStatus(err string) (int, string) {
	codeValue := strings.Split(err, ", ")[0]
	codestr := strings.Split(codeValue, "=")[1]
	codeint, _ := strconv.Atoi(codestr)

	messageValue := strings.Split(err, ", ")[1]
	message := strings.Split(messageValue, "=")[1]

	return codeint, message
}
