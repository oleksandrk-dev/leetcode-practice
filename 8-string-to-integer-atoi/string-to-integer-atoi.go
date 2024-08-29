func myAtoi(s string) int {
    var unwhitespaced string = strings.Trim(s, " ")
    if len(unwhitespaced) == 0 {
        return 0
    }
    isPositive := true 
    var curPos int = 0
    fmt.Printf("%s", unwhitespaced)
    if unwhitespaced[curPos] == '-' {
        curPos += 1
        isPositive = false
        fmt.Printf("negative")
    } else if unwhitespaced[curPos] == '+' {
        curPos += 1
        isPositive = true
        fmt.Printf("positive")
    }
    var result int = 0
    for pos :=curPos;pos<len(unwhitespaced);pos++ {
        digit := int(unwhitespaced[pos] - '0')
        if digit >= 0 && digit <= 9 {
            if checkIfInInt32Max(result, digit) {
                if isPositive {
                    return math.MaxInt32
                }
                return math.MinInt32
            }
            
            result = result*10+digit
        } else {
             if isPositive {
                return result
            }
            return result*-1
        }
    }

    if isPositive {
        return result
    }
    return result*-1
}

func checkIfInInt32Max(currentNumber int, digit int) bool {
    if currentNumber > 0 {
        if currentNumber > (math.MaxInt32-digit)/10 {
            return true
        }
        if currentNumber <= (math.MinInt32-digit+1)/10 {
            return true
        }
    }
    return false
} 