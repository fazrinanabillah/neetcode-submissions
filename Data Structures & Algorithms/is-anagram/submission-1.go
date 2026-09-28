func isAnagram(s string, t string) bool {
	for len(s) != len(t) {
		return false
	}
	var alphabet [26]int
	for i := 0; i < len(s); i++ {
		alphabet[s[i] - 'a']++
		alphabet[t[i] - 'a']--
	}

	for _, a := range alphabet {
		if a != 0 {
			return false
		}
	}
	return true
}
