package main

import (
    "fmt"
    "math"
    "sort"
)

func main() {
     fmt.Println(findMedianSortedArrays([]int{1, 2 ,3, 4, 5}, []int{5, 6, 7, 8, 9}))
}


// TODO: CONSIDER DDGE CASE: 2 ARRAYS OF DIFFERENT LENGTH
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
    for (len(nums1) > 2 && len(nums2) > 2 ) {
        var nums1_mi float64 = float64(len(nums1) - 1) / 2
        var nums2_mi float64 = float64(len(nums2) - 1) / 2
        var nums1_mv float64 = float64((nums1[int(math.Ceil(nums1_mi))] +
            nums1[int(math.Floor(nums1_mi))])) / 2

        var nums2_mv float64 = float64((nums2[int(math.Ceil(nums2_mi))] +
            nums2[int(math.Floor(nums2_mi))])) / 2

        if nums1_mv > nums2_mv {
            nums1 = nums1[:int(math.Floor(nums1_mi)) + 1]
            if len(nums2) % 2 == 0 {
                nums2 = nums2[int(math.Floor(nums2_mi)) + 1:]
            } else {
                nums2 = nums2[int(math.Floor(nums2_mi)):]
            }
        } else {
            nums2 = nums2[:int(math.Floor(nums2_mi)) + 1]
            if len(nums1) % 2 != 0 {
                nums1 = nums1[int(math.Floor(nums1_mi)):]
            } else {
                nums1 = nums1[int(math.Floor(nums1_mi)) + 1:]
            }
        }
    }

    var pm []int = append(nums1, nums2...)
    sort.Ints(pm)
    var c_mi float64 = float64(len(pm) - 1) / 2

    var combined_median = float64((pm[int(math.Ceil(c_mi))] +
        pm[int(math.Floor(c_mi))])) / 2

    return float64(combined_median)
}