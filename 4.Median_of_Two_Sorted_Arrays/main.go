package main

import "fmt"

func main() {
	nums1 := []int{1, 2}
	nums2 := []int{3, 5}
	fmt.Println(findMedianSortedArrays(nums1, nums2))
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	i, j := 0, 0
	nums3 := make([]int, len(nums1)+len(nums2))

	for k := 0; k < len(nums1)+len(nums2); k++ {
		if i == len(nums1) || (i < len(nums1) && j < len(nums2) && nums1[i] > nums2[j]) {
			nums3[k] = nums2[j]
			j++
			continue
		}
		if j == len(nums2) || (i < len(nums1) && j < len(nums2) && nums1[i] <= nums2[j]) {
			nums3[k] = nums1[i]
			i++
		}
	}

	l := len(nums3)
	if l%2 == 0 {
		return float64(nums3[l/2]+nums3[l/2-1]) / 2
	}
	return float64(nums3[l/2])
}
