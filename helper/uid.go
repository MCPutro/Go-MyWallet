package helper

import "strings"

func GetUserIdAndAccountId(data string) (string, string) {
	split := strings.Split(data, "-")

	return strings.Join(split[:len(split)-1], "-"), split[len(split)-1]
}
