package arrayops

import "sort"

func FindItemAndDelete(slice []int, targetItem int) []int {
	newSlice := []int{}
	for _, item := range slice {
		if item != targetItem {
			newSlice = append(newSlice, item)

		}
	}
	return newSlice
}

func SortArryDesc(slice []int) {
	  sort.Slice(slice,func(i int, j int)bool{
		return slice[i]< slice[j]
	})
}