func longestCommonPrefix(strs []string) string {
    var minLength int = 201
    for _, str := range strs {
        if curLen := len(str); curLen < minLength {
            minLength = curLen
        }
    }
    var result string
    for i := range minLength {
        firstChar := strs[0][i]
        for _, val := range strs {
            if val[i] != firstChar {
                return result
            }
        }
        result += string(firstChar) 
    } 
    return result
}