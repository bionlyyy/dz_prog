#include <iostream>
#include <fstream>
#include <vector>
#include <set>
#include <map>
#include <algorithm>
#include <locale>
#include <sstream>
#include <cstdlib>

using namespace std;


void setRussianLocale() {
    setlocale(LC_ALL, "Russian");
    locale::global(locale(""));
    cout.imbue(locale("rus_rus.866"));
    cerr.imbue(locale("rus_rus.866"));
}

// ������� ��� �������������� ������ � �����
int stringToInt(const string& s) {
    stringstream ss(s);
    int result;
    ss >> result;
    return result;
}

vector<int> readNumbers(const string& filename) {
    vector<int> numbers;
    ifstream file(filename, ios::binary);

    if (!file) {
        cerr << "������ �������� �����: " << filename << endl;
        return numbers;
    }

    string content((istreambuf_iterator<char>(file)),
        istreambuf_iterator<char>());
    file.close();

    // ��������� ����������� �����
    string currentNumber;
    for (char c : content) {
        if (isdigit(c) || (currentNumber.empty() && c == '-')) {
            currentNumber += c;
        }
        else if (!currentNumber.empty()) {
            numbers.push_back(stringToInt(currentNumber));
            currentNumber.clear();
        }
    }

    if (!currentNumber.empty()) {
        numbers.push_back(stringToInt(currentNumber));
    }

    return numbers;
}

void processFiles() {
    setRussianLocale();

    cout << "��������� ������..." << endl;

    vector<int> aNumbers = readNumbers("A.txt");
    vector<int> bNumbers = readNumbers("B.txt");

    // ������� �����������
    set<int> setA(aNumbers.begin(), aNumbers.end());
    set<int> setB(bNumbers.begin(), bNumbers.end());
    vector<int> commonNumbers;

    set_intersection(setA.begin(), setA.end(),
        setB.begin(), setB.end(),
        back_inserter(commonNumbers));

    // ������� ���������
    map<int, int> countA, countB;
    for (int num : aNumbers) countA[num]++;
    for (int num : bNumbers) countB[num]++;

    // ������ ����������
    ofstream outFile("C.txt");
    for (int num : commonNumbers) {
        int maxCount = max(countA[num], countB[num]);
        for (int i = 0; i < maxCount; ++i) {
            outFile << num << endl;
        }
    }

    cout << "��������� ������� � C.txt" << endl;
}

int main() {
    processFiles();
    return 0;
}