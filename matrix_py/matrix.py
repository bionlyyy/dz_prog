import numpy as np
def main():
    # 1. Чтение матрицы из файла input.txt
    with open('input.txt', 'r') as f:
        matrix = np.array([[float(num) for num in line.split()]
                           for line in f if line.strip()])

    # 2. Вычисление характеристик матрицы
    determinant = np.linalg.det(matrix)  # Определитель
    trace = np.trace(matrix)  # След матрицы
    transposed = matrix.T  # Транспонированная матрица

    # 3. Запись результатов в output.txt
    with open('output.txt', 'w', encoding='utf-8') as f:
        # Записываем определитель и след
        f.write(f"Определитель матрицы: {determinant:.2f}\n")
        f.write(f"След матрицы: {trace:.2f}\n\n")

        # Записываем транспонированную матрицу
        f.write("Транспонированная матрица:\n")
        np.savetxt(f, transposed, fmt="%8.2f")


if __name__ == "__main__":
    main()