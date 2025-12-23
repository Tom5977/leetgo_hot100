package lengthoflongestsubstring

func lengthOfLongestSubstring(s string) int {
	hashm := map[rune]bool{}
	res := 0
	for _, c := range s {
		if hashm[c] == true {
			if len(hashm) > res {
				res = len(hashm)
				hashm = map[rune]bool{}
			}
		}
		hashm[c] = true
	}
	return res
}
