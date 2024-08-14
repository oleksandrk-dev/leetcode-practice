func longestPalindrome(s string) string {
	if len(s) == 1 {
        return s
    }
    
    table := make([][]int, len(s))

	var startOfSubstring, lengthOfSubstring int = 0, 1

	for i := range table {
		table[i] = make([]int, len(s))
		table[i][i] = 1
        if i< len(s)-1 && s[i] == s[i+1] {
            table[i][i+1] = 2
            startOfSubstring = i
			lengthOfSubstring = 2
        }
	}



	for i :=3; i <= len(s); i++ {
		for left := 0; left < len(s)-i+1; left++ {
			var right int = left + i - 1
			if table[left+1][right-1] > 0 && s[left] == s[right] {
				table[left][right] = table[left+1][right-1] + 1
				
				if i > lengthOfSubstring {
                    // fmt.Printf("start will be - %d and end will be %d, this %s\n", left,i, s[startOfSubstring : startOfSubstring+lengthOfSubstring])
					startOfSubstring = left
					lengthOfSubstring = i
				}
			}
		}
	}
	// for _, row := range table {
	// 	fmt.Printf("%v\n", row)
	// }
	return s[startOfSubstring : startOfSubstring+lengthOfSubstring]
}