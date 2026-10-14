package repeat

func Repeat(word string, num int) string {
	result := ""
	for i := 0; i < num; i++ {

		result += word
	}
	return result
}
