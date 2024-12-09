package utils

func RemoveElementFromSlice[T any](arr []T, index int) []T {
	ret := make([]T, 0)
	ret = append(ret, arr[:index]...)
	return append(ret, arr[index+1:]...)
}
