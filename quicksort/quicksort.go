package main

import (
	"bufio"   // для чтения файла
	"fmt"     // для вывода
	"log"     // для ошибок
	"os"      // работа с файлами
	"sort"    // сортировка
	"strconv" // преобразование строк в числа
	"strings" // обработка строк
)

// Функция для превращения текста в числа
func convertTextToNumbers(textParts []string) []int64 {
	var numbersList []int64 // здесь будут числа

	for _, part := range textParts {
		cleanPart := strings.TrimSpace(part) // убираем пробелы
		if cleanPart == "" {
			continue // пропускаем пустые строки
		}

		number, err := strconv.ParseInt(cleanPart, 10, 64) // пробуем сделать число
		if err != nil {
			log.Printf("Пропускаем '%s' - это не число", part)
			continue
		}

		numbersList = append(numbersList, number) // добавляем число в список
	}

	return numbersList
}

func main() {
	// 1. Открываем файл
	fileWithNumbers, err := os.Open("input.txt")
	if err != nil {
		log.Fatal("Не могу открыть файл:", err)
	}
	defer fileWithNumbers.Close() // закроем файл в конце

	// Настраиваем сканер
	numberScanner := bufio.NewScanner(fileWithNumbers)
	numberScanner.Split(bufio.ScanWords) // читаем по словам

	var allTextParts []string // здесь будут все слова из файла

	// Читаем файл
	for numberScanner.Scan() {
		allTextParts = append(allTextParts, numberScanner.Text())
	}

	// Проверяем ошибки сканера
	if err := numberScanner.Err(); err != nil {
		log.Fatal("Ошибка при чтении файла:", err)
	}

	// 2. Превращаем текст в числа
	numbers := convertTextToNumbers(allTextParts)

	// 3. Сортируем числа
	sort.Slice(numbers, func(i, j int) bool {
		return numbers[i] < numbers[j] // сортируем по возрастанию
	})

	// 4. Создаем файл для результата
	resultFile, err := os.Create("output.txt")
	if err != nil {
		log.Fatal("Не могу создать файл для результата:", err)
	}
	defer resultFile.Close()

	// Записываем результат
	for _, num := range numbers {
		fmt.Fprintf(resultFile, "%d ", num) // пишем числа через пробел
	}

	fmt.Println("Готово! Результат в output.txt")
}
