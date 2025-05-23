package main

import (
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

// функция сортировки
func puz_sort(nums []int64) {
	n := len(nums)
	//проходим по всем элементам
	for i := 0; i < n; i++ {
		swapped := false //если обменов не было массив отсортирован

		//цикл попарно сравнивает элементы
		for j := 0; j < n-1; j++ {
			if nums[j] > nums[j+1] {
				// Если число больше следующего, меняем их местами
				nums[j], nums[j+1] = nums[j+1], nums[j]
				swapped = true
			}
		}

		// Если не было обменов, выходим из цикла
		if !swapped {
			break
		}
	}
}

func main() {
	// Открытие файла
	input, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("No file") //Сообщение об ошибке, если файл не найден
		os.Exit(1)
	}

	//Чтение данных из файла
	input_data := make([]byte, 64)
	for {
		_, err := input.Read(input_data)
		if err == io.EOF {
			break
		}
	}

	var nums []int64 //массив для хранения чисел из файла

	// Обработка прочитанных данных
	for _, i := range strings.Split(string(input_data), " ") {
		for _, j := range strings.Split(string(i), "\r") {
			for _, num := range strings.Split(string(j), "\x00") {
				//Пытаемся преобразовать строку в число
				tmp, _ := strconv.ParseInt(num, 10, 64)
				if tmp != 0 { //Игнорируем нулевые значения
					nums = append(nums, tmp)
				}
			}
		}
	}
	input.Close() //Закрываем входной файл

	//Создаем файл output.txt
	output, err := os.Create("output.txt")
	if err != nil {
		fmt.Println("Error creating file:", err)
		os.Exit(1)
	}
	defer output.Close() //закрытие файла при выходе из функции

	// Сортируем числа с помощью функции
	puz_sort(nums)

	//строкa для записи в файл
	output_str := ""
	for _, i := range nums {
		output_str += fmt.Sprintf("%d ", i) //Добавляем каждое число с пробелом
	}

	//Записываем результат в файл
	_, err = output.WriteString(output_str)
	if err != nil {
		fmt.Println("Error", err)
	}

	//Выводим отсортированный массив в консоль для проверки
	fmt.Println(nums)
}
