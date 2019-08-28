package comm

import "sort"

//SortStringArry 字符串slice排序，directory 为bigToSmall-从小到大,smallToBig-从大到小
//正常返回新的排序，若directory传错，则返回nil
func SortStringArry(arr []string, directory string) (newArray []string) {
	switch directory {
	case "bigToSmall":
		sort.Slice(arr, func(i, j int) bool {
			return arr[i] > arr[j]
		})
		return arr
	case "smallToBig":
		sort.Slice(arr, func(i, j int) bool {
			return arr[i] < arr[j]
		})
		return arr
	default:
		return nil
	}

}

//SortIntArry slice排序，directory 为bigToSmall-从小到大,smallToBig-从大到小
//正常返回新的排序，若directory传错，则返回nil
func SortIntArry(arr []int, directory string) (newArray []int) {
	switch directory {
	case "bigToSmall":
		sort.Slice(arr, func(i, j int) bool {
			return arr[i] > arr[j]
		})
		return arr
	case "smallToBig":
		sort.Slice(arr, func(i, j int) bool {
			return arr[i] < arr[j]
		})
		return arr
	default:
		return nil
	}

}
