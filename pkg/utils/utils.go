package utils

import "strings"

func GetAlphabet(fullName string, id string) []rune {

	fullName = strings.ReplaceAll(fullName, " ", "")

	var result = []rune(fullName)

	result = append(result, []rune(id)...)
	result = append(result, '.')

	return result
}

func GetInitials(data string) []rune {

	var names = strings.Split(data, " ")

	var result []rune

	for _, s := range names {
		result = append(result, []rune(s)[0])
	}

	return result
}
