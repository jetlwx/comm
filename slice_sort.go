package comm

import "sort"

//slice(string类型)元素去重
func RemoveReplicaSliceString(slc []string) []string {
	/*
	   slice(string类型)元素去重
	*/
	result := make([]string, 0)
	tempMap := make(map[string]bool, len(slc))
	for _, e := range slc {
		if tempMap[e] == false {
			tempMap[e] = true
			result = append(result, e)
		}
	}
	return result
}

//slice(int类型)元素去重
func RemoveReplicaSliceInt(slc []int) []int {
	/*
	   slice(int类型)元素去重
	*/
	result := make([]int, 0)
	tempMap := make(map[int]bool, len(slc))
	for _, e := range slc {
		if tempMap[e] == false {
			tempMap[e] = true
			result = append(result, e)
		}
	}
	return result
}

// SliceDeleteOne 删除字符串slice 中的值，若没找到，则返回原slice
func SliceDeleteOneString(arr []string, delstr string) (newArr []string) {
	for k, v := range arr {
		if v == delstr {
			arr = append(arr[:k], arr[k+1:]...)
			return arr
		}
	}
	return arr
}

//SortStringArry 字符串slice排序，directory 为bigToSmall-倒序,smallToBig-顺序
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
