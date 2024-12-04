package commonUtils

import "strings"

func FindAllIndexesOfString(input, subStr string) []int {
	result := make([]int, 0, 100)
	subStrLen := len(subStr)

	index := strings.Index(input, subStr)
	for index != -1 {
		result = append(result, index)

		index = strings.Index(input[index+subStrLen:], subStr)
		if index != -1 {
			index += result[len(result)-1] + subStrLen
		}
	}
	return result
}
