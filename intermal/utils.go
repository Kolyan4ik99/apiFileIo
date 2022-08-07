package intermal

import (
	"fmt"
	"math/rand"
	"strings"
)

func RandomString(n int) string {
	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

	s := make([]rune, n)
	for i := range s {
		s[i] = letters[rand.Intn(len(letters))]
	}
	return string(s)
}

func printMap(mapa map[string][]string) string {
	str := strings.Builder{}
	for s, strings := range mapa {
		str.WriteString(fmt.Sprintf("\t'%s:", s))
		for _, s2 := range strings {
			str.WriteString(fmt.Sprintf("%s'", s2))
		}
	}
	return str.String()
}
