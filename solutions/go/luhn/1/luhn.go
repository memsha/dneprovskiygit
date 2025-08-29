package luhn

import "strings"

func Valid(id string) bool {
	id = strings.ReplaceAll(id, " ", "")

	// Проверка длины
	if len(id) <= 1 {
		return false
	}

	sum := 0
	double := false

	// Проходим по цифрам справа налево
	for i := len(id) - 1; i >= 0; i-- {
		ch := id[i]

		// Проверяем, что это цифра
		if ch < '0' || ch > '9' {
			return false
		}

		digit := int(ch - '0')

		if double {
			digit *= 2
			if digit > 9 {
				digit -= 9
			}
		}

		sum += digit
		double = !double
	}

	return sum%10 == 0

}
