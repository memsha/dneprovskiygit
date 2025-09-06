package reverse

func Reverse(input string) string {
	runes := []rune(input)

	// Разворачиваем срез рун
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}

	// Возвращаем строку
	return string(runes)
}
