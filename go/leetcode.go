package _go

func mergeAlternately(word1 string, word2 string) string {
	var retByte []byte
	for i := 0; i < 100; i++ {
		if i > len(word1) && i > len(word2) {
			break
		}
		if i < len(word1) {
			retByte = append(retByte, word1[i])
		}
		if i < len(word2) {
			retByte = append(retByte, word2[i])
		}
	}
	return string(retByte)
}

func gcdOfStrings(str1 string, str2 string) string {
	return ""
}
