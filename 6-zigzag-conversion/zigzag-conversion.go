func convert(s string, numRows int) string {
    if numRows <= 1 {
        return s
    }

	var result string = ""
	var group int = 2*numRows - 2
    var current int
	for i := 1; i < numRows+1; i++ {
		current = currentValue(numRows, i)
        var index int = i -1
        for index < len(s) {
            result += string(s[index])
            index += current
            if current != group {
                current = group - current
            }
        }
	}
    return result
}

func currentValue(numRows, i int) int {
	if i == numRows {
		return 2*numRows - 2
	} else {
		return 2*numRows - 2*i
	}
}