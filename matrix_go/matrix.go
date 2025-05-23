package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	// Чтение матрицы из файла
	matrix, err := readMatrix("input.txt")
	if err != nil {
		fmt.Printf("Ошибка чтения матрицы: %v\n", err)
		return
	}

	// Проверка что матрица квадратная
	if len(matrix) != len(matrix[0]) {
		fmt.Println("Ошибка: матрица не квадратная, нельзя вычислить определитель и след")
		return
	}

	// Вычисление определителя
	det := determinant(matrix)

	// Вычисление следа
	trace := matrixTrace(matrix)

	// Транспонирование матрицы
	transposed := transposeMatrix(matrix)

	// Запись результатов в файл
	err = writeResults("output.txt", det, trace, transposed)
	if err != nil {
		fmt.Printf("Ошибка записи результатов: %v\n", err)
		return
	}

	fmt.Println("Вычисления завершены успешно. Результаты сохранены в output.txt")
}

// Чтение матрицы из файла
func readMatrix(filename string) ([][]float64, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var matrix [][]float64
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		if strings.TrimSpace(line) == "" {
			continue // Пропускаем пустые строки
		}

		fields := strings.Fields(line)
		row := make([]float64, len(fields))

		for i, field := range fields {
			val, err := strconv.ParseFloat(field, 64)
			if err != nil {
				return nil, fmt.Errorf("ошибка преобразования числа: %v", err)
			}
			row[i] = val
		}

		matrix = append(matrix, row)
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	if len(matrix) == 0 {
		return nil, fmt.Errorf("файл пустой")
	}

	return matrix, nil
}

// Запись результатов в файл
func writeResults(filename string, det, trace float64, transposed [][]float64) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := bufio.NewWriter(file)

	_, err = fmt.Fprintf(writer, "Определитель матрицы: %.2f\n", det)
	if err != nil {
		return err
	}

	_, err = fmt.Fprintf(writer, "След матрицы: %.2f\n", trace)
	if err != nil {
		return err
	}

	_, err = fmt.Fprintln(writer, "\nТранспонированная матрица:")
	if err != nil {
		return err
	}

	for _, row := range transposed {
		for _, val := range row {
			_, err = fmt.Fprintf(writer, "%8.2f ", val)
			if err != nil {
				return err
			}
		}
		_, err = fmt.Fprintln(writer)
		if err != nil {
			return err
		}
	}

	return writer.Flush()
}

// Вычисление определителя матрицы
func determinant(matrix [][]float64) float64 {
	n := len(matrix)

	// Базовые случаи
	if n == 1 {
		return matrix[0][0]
	}

	if n == 2 {
		return matrix[0][0]*matrix[1][1] - matrix[0][1]*matrix[1][0]
	}

	// Разложение по первой строке
	var det float64
	for col := 0; col < n; col++ {
		// Создаем минор для текущего элемента
		minor := make([][]float64, n-1)
		for i := range minor {
			minor[i] = make([]float64, n-1)
		}

		// Заполняем минор
		for i := 1; i < n; i++ {
			minorCol := 0
			for j := 0; j < n; j++ {
				if j == col {
					continue
				}
				minor[i-1][minorCol] = matrix[i][j]
				minorCol++
			}
		}

		// Рекурсивно вычисляем определитель минора
		sign := math.Pow(-1, float64(col))
		det += sign * matrix[0][col] * determinant(minor)
	}

	return det
}

// Вычисление следа матрицы
func matrixTrace(matrix [][]float64) float64 {
	var trace float64
	for i := 0; i < len(matrix); i++ {
		trace += matrix[i][i]
	}
	return trace
}

// Транспонирование матрицы
func transposeMatrix(matrix [][]float64) [][]float64 {
	rows := len(matrix)
	if rows == 0 {
		return nil
	}
	cols := len(matrix[0])

	transposed := make([][]float64, cols)
	for i := range transposed {
		transposed[i] = make([]float64, rows)
	}

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			transposed[j][i] = matrix[i][j]
		}
	}

	return transposed
}
