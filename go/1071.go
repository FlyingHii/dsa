package main

// based on the fact that for two strings str1 and str2, the largest common divisor string can only exist if both str1 + str2 == str2 + str1
func gcd(a, b int) int {
	for b > 0 {
		a, b = b, a%b
	}
	return a
}

func gcdOfStrings(str1 string, str2 string) string {
	if str1+str2 != str2+str1 {
		return ""
	}
	gcdIndex := gcd(len(str1), len(str2))
	return str1[:gcdIndex]
}
