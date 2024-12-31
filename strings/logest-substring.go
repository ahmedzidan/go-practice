package main

func lengthOfLongestSubstring(s string) int {
	var result int
	// aabc //
	// a => a X
	// a => ab Y
	// a => abc 3
	// longest > len(bc) yes break
	lenString := len(s)
	longest := 0
	for i := 0; i < len(s); i++ {
		if longest > result {
			result = longest
		}

		if lenString < result {
			break
		}
		longest = 0
		uniqueValues := map[byte]struct{}{}
		uniqueValues[s[i]] = struct{}{}
		for j := i + 1; j < len(s); j++ {
			if _, ok := uniqueValues[s[j]]; ok {
				break
			}
			longest += 1
			uniqueValues[s[j]] = struct{}{}
		}
		lenString--
	}
	return result
}

func lengthOfLongestSubstring2(s string) int {
	maxLength := 0
	start := 0
	uniqueChar := map[byte]int{}
	for end, v := range []byte(s) { // aba
		if index, ok := uniqueChar[v]; ok {
			start = index + 1
		}
		uniqueChar[v] = end
		maxLength = max(maxLength, end-start+1)
	}
	return maxLength
}

// q b q c
/*
map[q] = 0
for 0, q //1 b 2, q
if q in map { // b in map X
  start = max (start,index+1) // 0 + 1 = 1
}
map[q] = 0  // b = 1 //q = 2
max := max(0 , end-start+1) // 0-0+1 = 1 // 1-0+1 = 2 // (2, 2-1+1) 2
*/
