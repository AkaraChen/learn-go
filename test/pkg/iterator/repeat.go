package iterator

func Repeat(str string, time int) (repeated string) {
	repeated = ""
	for i := 0; i < time; i++ {
		repeated += str
	}
	return
}
