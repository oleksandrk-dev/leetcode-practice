func romanToInt(s string) int {
	allLen := len(s)
	// fmt.Print(string(s[1]))
	skipNext := false
	result := 0
	for i, val := range s {
		if skipNext {
			skipNext = false
			continue
		}
		currVal := string(val)

		switch currVal {
		case "M":
			result += 1000
			continue
		case "D":
			result += 500
			continue
		case "C":
			if allLen-i > 1 {
				nexto := string(s[i+1])
				if nexto == "M" {
					result += 900
					skipNext = true
					continue
				} else if nexto == "D" {
					result += 400
					skipNext = true
					continue
				}
			}
			result += 100
			continue
		case "L":
			result += 50
			continue
		case "X":
			if allLen-i > 1 {
				nexto := string(s[i+1])
				if nexto == "C" {
					result += 90
					skipNext = true
					continue
				} else if nexto == "L" {
					result += 40
					skipNext = true
					continue
				}

			}
			result += 10
			continue
		case "V":
			result += 5
			continue
		case "I":
			if allLen-i > 1 {
				nexto := string(s[i+1])
				if nexto == "X" {
					result += 9
					skipNext = true
					continue
				} else if nexto == "V" {
					result += 4
					skipNext = true
					continue
				}
			}
			result += 1
			continue
		}
	}
	return result
}
