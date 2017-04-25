package mysort

func MergeSort(inputSli []int) []int {
	var length = len(inputSli)
	var halfLen = length / 2
	if length == 1 {
		return inputSli
	}

	var left = inputSli[0:halfLen]
	var right = inputSli[halfLen:]

	return merge(MergeSort(left), MergeSort(right))
}

func merge(left []int, right []int) []int {
	var len1 = len(left)
	var len2 = len(right)
	var retSli = make([]int, 0)

	for i, j := 0, 0; ; {
		if left[i] < right[j] {
			retSli = append(retSli, left[i])
			i++
		} else {
			retSli = append(retSli, right[j])
			j++
		}

		if i > len1-1 {
			retSli = append(retSli, right[j:]...)
			return retSli
		}
		if j > len2-1 {
			retSli = append(retSli, left[i:]...)
			return retSli
		}
	}

	return retSli
}
