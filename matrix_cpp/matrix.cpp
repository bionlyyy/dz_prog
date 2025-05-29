#include <iostream>
#include <vector>
#include <fstream>
#include <cmath>
#include <iomanip>
#include <string>
#include <locale> 
using namespace std;

void setRussianLocale() {
    setlocale(LC_ALL, "Russian");
}

// Функция для чтения матрицы из файла
vector<vector<double>> readMatrix(const string& filename) {
    ifstream file(filename);
    vector<vector<double>> matrix;

    if (!file.is_open()) {
        cerr << "Ошибка: невозможно открыть файл " << filename << endl;
        return matrix;
    }

    string line;
    while (getline(file, line)) {
        vector<double> row;
        size_t pos = 0;
        while (pos < line.size()) {
            // Пропускаем пробелы
            while (pos < line.size() && line[pos] == ' ') pos++;
            if (pos >= line.size()) break;

            // Находим конец числа
            size_t endPos = line.find(' ', pos);
            if (endPos == string::npos) endPos = line.size();

            string numStr = line.substr(pos, endPos - pos);
            try {
                row.push_back(stod(numStr));
            }
            catch (...) {
                cerr << "Ошибка: неверный формат числа: " << numStr << endl;
                return {};
            }
            pos = endPos + 1;
        }
        if (!row.empty()) {
            matrix.push_back(row);
        }
    }

    file.close();
    return matrix;
}

// Функция для проверки, является ли матрица квадратной
bool isSquareMatrix(const vector<vector<double>>& matrix) {
    if (matrix.empty()) return false;
    size_t rows = matrix.size();
    for (const auto& row : matrix) {
        if (row.size() != rows) return false;
    }
    return true;
}

// Функция для вычисления следа матрицы 
double calculateTrace(const vector<vector<double>>& matrix) {
    if (!isSquareMatrix(matrix)) {
        cerr << "Ошибка: след можно вычислить только для квадратной матрицы" << endl;
        return 0;
    }

    double trace = 0.0;
    for (size_t i = 0; i < matrix.size(); ++i) {
        trace += matrix[i][i];
    }
    return trace;
}

// Функция для транспонирования матрицы
vector<vector<double>> transposeMatrix(const vector<vector<double>>& matrix) {
    if (matrix.empty()) return {};

    vector<vector<double>> result(matrix[0].size(), vector<double>(matrix.size()));

    for (size_t i = 0; i < matrix.size(); ++i) {
        for (size_t j = 0; j < matrix[0].size(); ++j) {
            result[j][i] = matrix[i][j];
        }
    }
    return result;
}

// Вспомогательная функция для вычисления определителя
double calculateDeterminant(const vector<vector<double>>& matrix);

// Функция для получения минора матрицы 
vector<vector<double>> getMinor(const vector<vector<double>>& matrix, size_t row, size_t col) {
    vector<vector<double>> minor(matrix.size() - 1, vector<double>(matrix.size() - 1));

    for (size_t i = 0, m = 0; i < matrix.size(); ++i) {
        if (i == row) continue;
        for (size_t j = 0, n = 0; j < matrix.size(); ++j) {
            if (j == col) continue;
            minor[m][n] = matrix[i][j];
            ++n;
        }
        ++m;
    }
    return minor;
}

// Рекурсивная функция для вычисления определителя матрицы
double calculateDeterminant(const vector<vector<double>>& matrix) {
    if (!isSquareMatrix(matrix)) {
        cerr << "Ошибка: определитель можно вычислить только для квадратной матрицы" << endl;
        return 0;
    }

    // Базовые случаи рекурсии
    if (matrix.size() == 1) {
        return matrix[0][0];
    }

    if (matrix.size() == 2) {
        return matrix[0][0] * matrix[1][1] - matrix[0][1] * matrix[1][0];
    }

    // Разложение по первой строке
    double determinant = 0.0;
    for (size_t j = 0; j < matrix.size(); ++j) {
        vector<vector<double>> minor = getMinor(matrix, 0, j);
        double sign = (j % 2 == 0) ? 1.0 : -1.0;
        determinant += sign * matrix[0][j] * calculateDeterminant(minor);
    }

    return determinant;
}

// Функция для записи результатов в файл
void writeResults(const string& filename, double trace,
    const vector<vector<double>>& transposed,
    double determinant) {
    ofstream file(filename);

    if (!file.is_open()) {
        cerr << "Ошибка: невозможно создать файл " << filename << endl;
        return;
    }

    file << "Результаты вычислений:\n";
    file << "======================\n\n";

    file << "1. След матрицы: " << trace << "\n\n";

    file << "2. Транспонированная матрица:\n";
    for (const auto& row : transposed) {
        for (double val : row) {
            file << setw(10) << val << " ";
        }
        file << "\n";
    }
    file << "\n";

    file << "3. Определитель матрицы: " << determinant << "\n";

    file.close();
}

int main() {
    setRussianLocale(); 

    vector<vector<double>> matrix = readMatrix("input.txt");

    if (matrix.empty()) {
        cerr << "Ошибка: не удалось прочитать матрицу из файла" << endl;
        return 1;
    }

    // Вычисляем характеристики матрицы
    double tr = calculateTrace(matrix);
    vector<vector<double>> transposed = transposeMatrix(matrix);
    double det = calculateDeterminant(matrix);

    // Записываем результаты в файл
    writeResults("output.txt", tr, transposed, det);

    return 0;
}