package lpath

import "strings"

func Split(s string) []string {
	return strings.Split(s, "/")
}

func Split2(s string) (string, string) {
	components := strings.Split(s, "/")
	if len(components) != 2 {
		panic("Something went wrong splitting path into two parts")
	}
	return components[0], components[1]
}

func Split3(s string) (string, string, string) {
	components := strings.Split(s, "/")
	if len(components) != 3 {
		panic("Something went wrong splitting path into three parts")
	}
	return components[0], components[1], components[2]
}
