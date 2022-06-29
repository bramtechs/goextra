package goextra

// Does not respect element order!
// https://stackoverflow.com/questions/37334119/how-to-delete-an-element-from-a-slice-in-golang
func RemoveIndexFreely[K any](src []K, i int) []K {
	src[i] = src[len(src)-1]
	return src[:len(src)-1]
}
