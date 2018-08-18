package typo

// 获取元素在数组中的索引，如果未找到则返回-1
func IndexOf(arr *[]string, elmt Any) int {
	if arr == nil {
		return -1
	}

	for i, v := range *arr {
		if v == elmt {
			return i
		}
	}

	return -1
}

// 判断元素是否在数组中，并且返回其索引
func InArray(elmt Any, arr Array) (bool, int) {
	for i, v := range arr {
		if elmt == v {
			return true, i
		}
	}
	return false, -1
}

// 数组逆序
func ReverseStringArray(arr []string) []string {
	for i := 0; i < len(arr)/2; i++ {
		j := len(arr) - i - 1
		arr[i], arr[j] = arr[j], arr[i]
	}
	return arr
}

func IsStringArraySorted(arr []string) bool {
	if l := len(arr); l < 2 {
		return true
	}

	var prev string
	for _, v := range arr {
		if v < prev {
			return false
		}
		prev = v
	}

	return true
}
