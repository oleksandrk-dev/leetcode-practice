func lengthOfLongestSubstring(s string) int {
    var dict map[rune]int = map[rune]int{}
    var (
        counter int = 0
        maxSoFar int = 0
    )
    for pos, el := range s {
        if val, ok := dict[el]; !ok {
            dict[el] = pos
            counter += 1
            if counter > maxSoFar {
                maxSoFar = counter
            }
        } else {
            dict[el] = pos
            counter = pos - val 
            for key, value := range dict {
                if value < val {
                    delete(dict, key)
                }
            }
        }
    }
    return maxSoFar
}