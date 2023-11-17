package fp

func addOne(i int) int {
	return i + 1
}

func square(i int) int {
	return i * i
}

func isHello(s string) bool {
	return s == "hello"
}

func toInt(b bool) int {
	if b {
		return 1
	}
	return 0
}
