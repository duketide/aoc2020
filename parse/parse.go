package parse

import "strings"

func Lines(s string) []string {
	return strings.Split(strings.TrimRight(s, "\n"), "\n")
}

func Words(s string) []string {
	return strings.Split(s, " ")
}
