package sorting

func MergeSort(arr []int) []int {
	res := make([]int, len(arr))
	i, j, k := 0, 0, 0
	lo := 0
	hi := len(arr)
	if len(arr) <= 1 {
		return arr
	}
	mid := (lo + hi) / 2
	leftArr := arr[lo:mid]
	rightArr := arr[mid:hi]
	sortedLeftArr := MergeSort(leftArr)
	sortedRightArr := MergeSort(rightArr)

	for i < len(sortedLeftArr) && j < len(sortedRightArr) {
		if sortedLeftArr[i] <= sortedRightArr[j] {
			res[k] = sortedLeftArr[i]
			i += 1
			k += 1
		} else {
			res[k] = sortedRightArr[j]
			j += 1
			k += 1
		}
	}

	for i < len(sortedLeftArr) {
		res[k] = sortedLeftArr[i]
		i += 1
		k += 1
	}

	for j < len(sortedRightArr) {
		res[k] = sortedRightArr[j]
		j += 1
		k += 1
	}
	return res
}

func QuickSort(arr []int) {
	QuickSortHelper(arr, 0, len(arr)-1)
}

func QuickSortHelper(arr []int, lo, hi int) {
	if lo < hi {
		split := Partition(arr, lo, hi)
		QuickSortHelper(arr, lo, split)
		QuickSortHelper(arr, split+1, hi)
	}
}

func Partition(arr []int, lo, hi int) int {
	pivotVal := arr[lo]
	left := lo + 1
	right := hi
	done := false
	for !done {
		for left <= right && arr[left] < pivotVal {
			left = left + 1
		}
		for arr[right] > pivotVal && right >= left {
			right = right - 1
		}
		if left > right {
			done = true
		} else {
			arr[left], arr[right] = arr[right], arr[left]
		}
	}
	arr[right], arr[lo] = arr[lo], arr[right]
	return right
}
