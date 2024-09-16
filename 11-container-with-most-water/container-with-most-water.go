func maxArea(height []int) int {
	var i, j, max int
	i = 0
	j = len(height) - 1
	max = returnBigger(height[i], height[j]) * (j - i)
	for {
		fmt.Printf("%d - %d - %d\n", i, j, max)
		if i+1 >= j {
			return max
		}
		leftNewMax := returnBigger(height[i+1], height[j]) * (j - i - 1)
		rightNewMax := returnBigger(height[i], height[j-1]) * (j - i - 1)
		if leftNewMax > max {
			max = leftNewMax
		} else if rightNewMax > max {
			max = rightNewMax

		}

		if height[i] < height[j] {
			i = i + 1
		} else {
			j = j - 1
		}

	}
}

func returnBigger(i, j int) int {
	if i > j {
		return j
	}
	return i
}