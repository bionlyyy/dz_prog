#include <iostream>   
#include <fstream>   
#include <vector>     
#include <sstream>    
#include <stdexcept>  

using namespace std; 


vector<vector<double>> readMatrix(const string& filename) {
    // ��������� ���� ��� ������
    ifstream file(filename);

    // ��������� ���������� �������� �����
    if (!file.is_open()) {
        throw runtime_error("�� ������� ������� ����: " + filename);
    }

    vector<vector<double>> matrix; // ������� ������ �������
    string line; // ����� ��� �������� ������ �� �����

    // ������ ���� ���������
    while (getline(file, line)) {
        vector<double> row; // ������ ��� ������� ������ �������
        stringstream ss(line); // ����� ��� ������� ������
        double value; // ���������� ��� �������� �����

        // ��������� ������ �� ��������� �����
        while (ss >> value) {
            row.push_back(value); // ��������� ����� � ������ �������
        }

        // ��������� ������ � �������, ���� ��� �� ������
        if (!row.empty()) {
            matrix.push_back(row);
        }
    }

    return matrix; // ���������� ����������� �������
}


vector<vector<double>> matrixMultiply(const vector<vector<double>>& A,
    const vector<vector<double>>& B) {
    // �������� ����������� ��������� ������
    if (A.empty() || B.empty() || A[0].size() != B.size()) {
        throw invalid_argument("������������� ������� ������ ��� ���������");
    }

    // ���������� ������� �������������� �������
    size_t m = A.size();    // ����� ����� ������ �������
    size_t n = B[0].size(); // ����� �������� ������ �������
    size_t p = B.size();    // ����� ������ (������� A � ������ B)

    // ������� �������-���������, ����������� ������
    vector<vector<double>> result(m, vector<double>(n, 0.0));

    // �������� ��������� ������ (������� ����)
    for (size_t i = 0; i < m; ++i) {         // ���� �� ������� A
        for (size_t j = 0; j < n; ++j) {     // ���� �� �������� B
            for (size_t k = 0; k < p; ++k) { // ��������� ������������
                result[i][j] += A[i][k] * B[k][j];
            }
        }
    }

    return result; // ���������� ��������� ���������
}


void writeMatrix(const vector<vector<double>>& matrix, const string& filename) {
    // ��������� ���� ��� ������
    ofstream file(filename);

    // ��������� ���������� �������� �����
    if (!file.is_open()) {
        throw runtime_error("�� ������� ������� ����: " + filename);
    }

    // ���������� ������� ���������
    for (const auto& row : matrix) {
        // ���������� �������� ������ ����� ������
        for (size_t i = 0; i < row.size(); ++i) {
            file << row[i]; // ���������� ������� �������

            // ��������� ������, ���� ��� �� ��������� �������
            if (i != row.size() - 1) {
                file << " ";
            }
        }
        file << endl; // ������� �� ����� ������
    }
}


int main() {
    try {
        // 1. ������ ������� ������
        auto A = readMatrix("A.txt");
        auto B = readMatrix("B.txt");

        // 2. ��������� ������
        auto C = matrixMultiply(A, B);

        // 3. ������ ����������
        writeMatrix(C, "C.txt");

    }
    catch (const exception& e) {
        // � ������ ������ ���������� ��� 1
        return 1;
    }

    // �������� ���������� ���������
    return 0;
}