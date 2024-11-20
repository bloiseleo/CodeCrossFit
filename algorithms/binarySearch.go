package algorithms

func BinarySearch(data []int, term int) bool {
	low, high := 0, len(data)-1
	return recursionSearch(data, low, high, term)
}

func recursionSearch(data []int, low, high, term int) bool {
	for high >= low {
		middle := low + (high-low)/2
		if data[middle] == term {
			return true
		}
		if data[middle] < term {
			low = middle + 1
		} else {
			high = middle - 1
		}
	}
	return false
}
