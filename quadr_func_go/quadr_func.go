package main

import (
	"fmt"
	"math"
)

// Функция для решения квадратного уравнения ax² + bx + c = 0
func solveQuadratic(a, b, c float64) (x1, x2 complex128) {
	// Обрабатываем вырожденные случаи
	if a == 0 {
		if b == 0 {
			if c == 0 {
				// 0 = 0 - бесконечное множество решений
				return complex(math.Inf(1), 0), complex(math.Inf(1), 0)
			}
			// c ≠ 0, 0 = c - нет решений
			return complex(math.NaN(), 0), complex(math.NaN(), 0)
		}
		// Линейное уравнение bx + c = 0
		solution := complex(-c/b, 0)
		return solution, solution
	}

	// Вычисляем дискриминант
	discriminant := b*b - 4*a*c

	// Вычисляем корни уравнения
	if discriminant >= 0 {
		// Вещественные корни
		sqrtD := math.Sqrt(discriminant)
		x1 = complex((-b+sqrtD)/(2*a), 0)
		x2 = complex((-b-sqrtD)/(2*a), 0)
	} else {
		// Комплексные корни
		sqrtD := math.Sqrt(-discriminant)
		realPart := -b / (2 * a)
		imagPart := sqrtD / (2 * a)
		x1 = complex(realPart, imagPart)
		x2 = complex(realPart, -imagPart)
	}

	return x1, x2
}

// Функция для форматирования комплексного числа
func formatComplex(z complex128) string {
	re, im := real(z), imag(z)

	// Проверка на специальные значения
	if math.IsNaN(re) {
		return "нет решений"
	}
	if math.IsInf(re, 0) {
		return "бесконечное множество решений"
	}

	// Форматирование обычных чисел
	if im == 0 {
		return fmt.Sprintf("%.2f", re)
	}
	if re == 0 {
		return fmt.Sprintf("%.2fi", im)
	}

	sign := "+"
	if im < 0 {
		sign = "-"
		im = -im
	}
	return fmt.Sprintf("%.2f %s %.2fi", re, sign, im)
}

func main() {
	var a, b, c float64

	fmt.Println("Решение квадратного уравнения ax² + bx + c = 0")
	fmt.Println("Введите коэффициенты:")

	fmt.Print("a = ")
	fmt.Scan(&a)
	fmt.Print("b = ")
	fmt.Scan(&b)
	fmt.Print("c = ")
	fmt.Scan(&c)

	x1, x2 := solveQuadratic(a, b, c)

	fmt.Printf("\nУравнение: %.2fx² + %.2fx + %.2f = 0\n", a, b, c)

	// Проверка особых случаев
	if a == 0 {
		fmt.Println("Примечание: Это линейное уравнение (a = 0)")
	}

	// Вывод результатов
	switch {
	case math.IsInf(real(x1), 0):
		fmt.Println("Результат: Бесконечное множество решений")
	case math.IsNaN(real(x1)):
		fmt.Println("Результат: Нет действительных решений")
	case x1 == x2:
		fmt.Printf("Результат: Один корень x = %s\n", formatComplex(x1))
	default:
		fmt.Printf("Результат: Два корня\nx₁ = %s\nx₂ = %s\n", formatComplex(x1), formatComplex(x2))
	}
}
