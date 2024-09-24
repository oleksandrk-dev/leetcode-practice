func intToRoman(num int) string {
	var result string = ""

	if el := num / 1000; el > 0 {
		result += strings.Repeat("M", el)
		num = num - el*1000
	}
    fmt.Printf("%d\n", num)

	if num >= 900 && num % 900 <= 99 {
		result += "CM"
		num = num - 900
	}
     fmt.Printf("%d\n", num)

	if el := num / 500; el > 0 {
		result += strings.Repeat("D", el)
		num = num - el*500
	}
     fmt.Printf("%d\n", num)

	if num >= 400 && num % 400 <= 99 {
		result += "CD"
		num = num - 400
	}
     fmt.Printf("%d\n", num)

	if el := num / 100; el > 0 {
		result += strings.Repeat("C", el)
		num = num - el*100
	}
     fmt.Printf("%d\n", num)

	if num >= 90 && num % 90 <= 9{
		result += "XC"
		num = num - 90
	}
     fmt.Printf("%d\n", num)

	if el := num / 50; el > 0 {
		result += strings.Repeat("L", el)
		num = num - el*50

	}
     fmt.Printf("%d\n", num)

	if num >=40 && num % 40 <= 9 {
		result += "XL"
		num = num - 40
	}
     fmt.Printf("%d\n", num)

	if el := num / 10; el > 0 {
		result += strings.Repeat("X", el)
		num = num - el*10
	}
     fmt.Printf("%d\n", num)

	if num == 9 {
		result += "IX"
		num = num - 9
	}
     fmt.Printf("%d\n", num)
	if el := num / 5; el > 0 {
		result += strings.Repeat("V", el)
		num = num - el*5

	}
     fmt.Printf("%d\n", num)
	if num == 4 {
		result += "IV"
		num = num - 4
	}
     fmt.Printf("%d\n", num)

	if el := num / 1; el > 0 {
		result += strings.Repeat("I", el)
		num = num - el*1
	}

	return result
}