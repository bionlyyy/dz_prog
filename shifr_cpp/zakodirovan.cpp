#include <iostream>
#include <fstream>
#include <string>
#include <locale>

using namespace std;

// Атбаш
string atbashCipher(const string& text) {
    string result;

    // Русский алфавит в верхнем и нижнем регистрах
    const string upperAlphabet = "АБВГДЕЁЖЗИЙКЛМНОПРСТУФХЦЧШЩЪЫЬЭЮЯ";
    const string lowerAlphabet = "абвгдеёжзийклмнопрстуфхцчшщъыьэюя";

    for (char c : text) {
        if (upperAlphabet.find(c) != string::npos) {
            // Шифруем заглавные буквы
            size_t index = upperAlphabet.find(c);
            result += upperAlphabet[upperAlphabet.length() - 1 - index];
        }
        else if (lowerAlphabet.find(c) != string::npos) {
            // Шифруем строчные буквы
            size_t index = lowerAlphabet.find(c);
            result += lowerAlphabet[lowerAlphabet.length() - 1 - index];
        }
        else {
            
            result += c;
        }
    }

    return result;
}

// Цезарь
string caesarEncrypt(const string& text) {
    string result;
    const int shift = 3; // сдвиг

    // Русский алфавит в верхнем и нижнем регистрах
    const string upperAlphabet = "АБВГДЕЁЖЗИЙКЛМНОПРСТУФХЦЧШЩЪЫЬЭЮЯ";
    const string lowerAlphabet = "абвгдеёжзийклмнопрстуфхцчшщъыьэюя";

    for (char c : text) {
        if (upperAlphabet.find(c) != string::npos) {
            // Шифруем заглавные буквы
            size_t index = upperAlphabet.find(c);
            index = (index + shift) % upperAlphabet.length();
            result += upperAlphabet[index];
        }
        else if (lowerAlphabet.find(c) != string::npos) {
            // Шифруем строчные буквы
            size_t index = lowerAlphabet.find(c);
            index = (index + shift) % lowerAlphabet.length();
            result += lowerAlphabet[index];
        }
        else {
           
            result += c;
        }
    }

    return result;
}

int main() {
    
    setlocale(LC_ALL, "Russian");

    // Открываем входной файл
    ifstream inputFile("input.txt");
    if (!inputFile.is_open()) {
        cerr << "Ошибка открытия файла input.txt" << endl;
        return 1;
    }

    // Читаем весь файл в строку
    string text((istreambuf_iterator<char>(inputFile)),
        istreambuf_iterator<char>());
    inputFile.close();

    // Применяем оба шифра последовательно
    string atbashResult = atbashCipher(text);
    string finalResult = caesarEncrypt(atbashResult);

    // Записываем результат в выходной файл
    ofstream outputFile("output.txt");
    if (!outputFile.is_open()) {
        cerr << "Ошибка создания файла output.txt" << endl;
        return 1;
    }
    outputFile << "Зашифровано Цезарем:\n" + finalResult + "\n\n" + "Зашифровано Атбашем:\n" + atbashResult+ "\n\n";
    outputFile.close();


    return 0;
}