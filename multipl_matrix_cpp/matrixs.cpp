#include <iostream>   
#include <fstream>   
#include <vector>     
#include <sstream>    
#include <stdexcept>  

using namespace std; 


vector<vector<double>> readMatrix(const string& filename) {
    // Открываем файл для чтения
    ifstream file(filename);

    // Проверяем успешность открытия файла
    if (!file.is_open()) {
        throw runtime_error("Не удалось открыть файл: " + filename);
    }

    vector<vector<double>> matrix; // Создаем пустую матрицу
    string line; // Буфер для хранения строки из файла

    // Читаем файл построчно
    while (getline(file, line)) {
        vector<double> row; // Вектор для текущей строки матрицы
        stringstream ss(line); // Поток для разбора строки
        double value; // Переменная для хранения чисел

        // Разбираем строку на отдельные числа
        while (ss >> value) {
            row.push_back(value); // Добавляем число в строку матрицы
        }

        // Добавляем строку в матрицу, если она не пустая
        if (!row.empty()) {
            matrix.push_back(row);
        }
    }

    return matrix; // Возвращаем заполненную матрицу
}


vector<vector<double>> matrixMultiply(const vector<vector<double>>& A,
    const vector<vector<double>>& B) {
    // Проверка возможности умножения матриц
    if (A.empty() || B.empty() || A[0].size() != B.size()) {
        throw invalid_argument("Несовместимые размеры матриц для умножения");
    }

    // Определяем размеры результирующей матрицы
    size_t m = A.size();    // Число строк первой матрицы
    size_t n = B[0].size(); // Число столбцов второй матрицы
    size_t p = B.size();    // Общий размер (столбцы A и строки B)

    // Создаем матрицу-результат, заполненную нулями
    vector<vector<double>> result(m, vector<double>(n, 0.0));

    // Алгоритм умножения матриц (тройной цикл)
    for (size_t i = 0; i < m; ++i) {         // Цикл по строкам A
        for (size_t j = 0; j < n; ++j) {     // Цикл по столбцам B
            for (size_t k = 0; k < p; ++k) { // Скалярное произведение
                result[i][j] += A[i][k] * B[k][j];
            }
        }
    }

    return result; // Возвращаем результат умножения
}


void writeMatrix(const vector<vector<double>>& matrix, const string& filename) {
    // Открываем файл для записи
    ofstream file(filename);

    // Проверяем успешность открытия файла
    if (!file.is_open()) {
        throw runtime_error("Не удалось создать файл: " + filename);
    }

    // Записываем матрицу построчно
    for (const auto& row : matrix) {
        // Записываем элементы строки через пробел
        for (size_t i = 0; i < row.size(); ++i) {
            file << row[i]; // Записываем текущий элемент

            // Добавляем пробел, если это не последний элемент
            if (i != row.size() - 1) {
                file << " ";
            }
        }
        file << endl; // Переход на новую строку
    }
}


int main() {
    try {
        // 1. Чтение входных матриц
        auto A = readMatrix("A.txt");
        auto B = readMatrix("B.txt");

        // 2. Умножение матриц
        auto C = matrixMultiply(A, B);

        // 3. Запись результата
        writeMatrix(C, "C.txt");

    }
    catch (const exception& e) {
        // В случае ошибки возвращаем код 1
        return 1;
    }

    // Успешное завершение программы
    return 0;
}