func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	var num3 []int = []int{}

	var medianPlace int = (len(nums1) + len(nums2)) / 2
	fmt.Printf("median - %d\n", medianPlace)
	var isEven bool = (len(nums1)+len(nums2))%2 == 0
	i, j := 0, 0
	for (i < len(nums1) || j < len(nums2)) || i+j < medianPlace {
		if i >= len(nums1) {
			fmt.Printf("i = %d, j = %d left ended  - %d\n", i, j, nums2[j])
			num3 = append(num3, nums2[j])
			j++
			continue
		} 
        
        if j >= len(nums2) {
			fmt.Printf("i = %d, j = %d right ended - %d\n", i, j, nums1[i])
			num3 = append(num3, nums1[i])
			i++
			continue
		}

		if nums1[i] < nums2[j] {
			num3 = append(num3, nums1[i])
			fmt.Printf("adding left - %d\n", nums1[i])
			i++
		} else {
			num3 = append(num3, nums2[j])
			fmt.Printf("adding right - %d\n", nums2[j])
			j++
		}
	}

	if isEven {
		return float64((num3[medianPlace-1] + num3[medianPlace])) / 2
	}
	return float64(num3[medianPlace])
}