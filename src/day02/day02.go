package day02

func LetterCounter(input string) map[string]int {
	output := map[string]int{}
	for _, char := range input{
		output[string(char)] += 1
	}
	return output
}

func ChecksumCalculator(input []string) int {
	seenTwo := 0
	seenThree := 0
	for _, element := range input {
		countMap := LetterCounter(element)
		twoChecked := false
		threeChecked := false
		for _, count := range countMap {
			if count == 2 && !twoChecked {
				seenTwo += 1
				twoChecked = true
			}
			if count == 3 && !threeChecked {
				seenThree += 1
				threeChecked = true
			}
		}
	}
	return seenTwo * seenThree
}

func DifferentCharCount(str1 string, str2 string) int {
	result := 0
	for i := 0; i < len(str1); i += 1 {
		if str1[i] != str2[i] {
			result += 1
		}
	}
	return result
}

func CommonLetters(str1 string, str2 string) string {
	result := ""
	for i := 0; i < len(str1); i += 1 {
		if str1[i] == str2[i] {
			result += string(str1[i])
		}
	}
	return result
}

func BoxIdComparer(input []string) string {
	for _, boxId := range input {
		for _, otherBoxId := range input {
			if boxId == otherBoxId {
				continue
			}
			if DifferentCharCount(boxId, otherBoxId) == 1 {
				return CommonLetters(boxId, otherBoxId)
			}
		}
	}
	return ""
}
