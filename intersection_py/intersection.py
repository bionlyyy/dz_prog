from collections import Counter

def process_files():
    # Читаем числа из файлов, пропуская нечисловые значения
    def read_numbers(filename):
        numbers = []
        try:
            with open(filename, 'r', encoding='utf-16') as f:  
                for line in f:
                    for item in line.split():
                        try:
                            numbers.append(int(item))
                        except ValueError:
                            print(f"Пропущено нечисловое значение: '{item}' в {filename}")
        except UnicodeError:
            try:
                with open(filename, 'r', encoding='utf-8') as f:  
                    for line in f:
                        for item in line.split():
                            try:
                                numbers.append(int(item))
                            except ValueError:
                                print(f"Пропущено нечисловое значение: '{item}' в {filename}")
            except Exception as e:
                print(f"Ошибка чтения {filename}: {e}")
                return []
        return numbers

    # Читаем оба файла
    a_numbers = read_numbers('A.txt')
    b_numbers = read_numbers('B.txt')

    if not a_numbers or not b_numbers:
        print("Один из файлов пуст или содержит только нечисловые данные!")
        return

    # Находим пересечение чисел
    common_numbers = sorted(set(a_numbers) & set(b_numbers))

    # Считаем вхождения каждого числа
    count_a = Counter(a_numbers)
    count_b = Counter(b_numbers)

    # Записываем в C.txt каждое число max(count_a, count_b) раз
    with open('C.txt', 'w', encoding='utf-8') as f:
        for num in common_numbers:
            max_count = max(count_a[num], count_b[num])
            f.write((str(num) + '\n') * max_count)

if __name__ == "__main__":
    process_files()
