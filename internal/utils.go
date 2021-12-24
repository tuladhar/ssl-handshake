package internal

func Sum(arr []int64) int64 {
	var result int64
	for i := 0; i < len(arr); i++ {
		result += arr[i]
	}
	return result
}
