package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

// Функция шифрования/дешифрования Цезаря
func caesar(text string, shift int) string {
	var result strings.Builder

	for _, s := range text {
		switch {
		case 'a' <= s && s <= 'z': // Английские строчные
			result.WriteRune('a' + (s-'a'+rune(shift)+26)%26)
		case 'A' <= s && s <= 'Z': // Английские заглавные
			result.WriteRune('A' + (s-'A'+rune(shift)+26)%26)
		case 'а' <= s && s <= 'я': // Русские строчные
			result.WriteRune('а' + (s-'а'+rune(shift)+32)%32)
		case 'А' <= s && s <= 'Я': // Русские заглавные
			result.WriteRune('А' + (s-'А'+rune(shift)+32)%32)
		default:
			result.WriteRune(s) // Остальные символы без изменений
		}
	}
	return result.String()
}

// Функция шифрования/дешифрования Атбаша
func atbash(text string) string {
	var result strings.Builder

	for _, s := range text {
		switch {
		case 'a' <= s && s <= 'z': // Английские строчные
			result.WriteRune('a' + 'z' - s)
		case 'A' <= s && s <= 'Z': // Английские заглавные
			result.WriteRune('A' + 'Z' - s)
		case 'а' <= s && s <= 'я': // Русские строчные
			result.WriteRune('а' + 'я' - s)
		case 'А' <= s && s <= 'Я': // Русские заглавные
			result.WriteRune('А' + 'Я' - s)
		default:
			result.WriteRune(s) // Остальные символы без изменений
		}
	}
	return result.String()
}

func main() {
	// Чтение исходного файла
	data, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatal("Ошибка чтения файла:", err)
	}
	text := string(data)

	key := 3 // Ключ для шифра Цезаря

	// Применяем шифрование
	caesarEncrypted := caesar(text, key)
	caesarDecrypted := caesar(caesarEncrypted, -key)
	atbashEncrypted := atbash(text)
	atbashDecrypted := atbash(atbashEncrypted)

	// Формируем результат для записи в файл
	output := fmt.Sprintf(`
Исходный текст:
%s

Зашифровано Цезарем :
%s

Расшифровано Цезарем:
%s

Зашифровано Атбашем:
%s

Расшифровано Атбашем:
%s
`, text, caesarEncrypted, caesarDecrypted, atbashEncrypted, atbashDecrypted)

	// Запись в файл output.txt
	err = ioutil.WriteFile("output.txt", []byte(output), 0644)
	if err != nil {
		log.Fatal("Ошибка записи в файл:", err)
	}

}
