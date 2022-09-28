package arrayops

func FindItemAndDelete(slice []int, targetItem int) []int {
	newSlice := []int{}
	for _, item := range slice {
		if item != targetItem {
			newSlice = append(newSlice, item)

		}
	}
	return newSlice
}