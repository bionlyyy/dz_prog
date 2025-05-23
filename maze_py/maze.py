from collections import deque  #deque для эффективной реализации очереди


def find_shortest_path(maze):
    """Находит кратчайший путь от S до E в лабиринте с помощью BFS"""

    #Поиск начальной (S) и конечной (E) точек в лабиринте
    start = end = None
    for i in range(len(maze)):
        for j in range(len(maze[i])):
            if maze[i][j] == 'S':
                start = (i, j)
            elif maze[i][j] == 'E':
                end = (i, j)

    #Проверка наличия обеих точек
    if not start or not end:
        return None  # Если нет старта или выхода, возвращаем None

    #Возможные направления движения (вверх, вниз, влево, вправо)
    directions = [(-1, 0), (1, 0), (0, -1), (0, 1)]

    #Инициализация очереди для BFS и словаря для отслеживания посещенных клеток
    queue = deque([start])  #Очередь начинается со стартовой позиции
    visited = {start: None}  #Словарь хранит связь "клетка -> предыдущая клетка"

    #Основной цикл BFS
    while queue:
        current = queue.popleft()  #Берем первую клетку из очереди

        #Если достигли конечной точки, прерываем цикл
        if current == end:
            break

        #Проверяем все соседние клетки
        for di, dj in directions:
            ni, nj = current[0] + di, current[1] + dj  #Координаты соседа

            #Проверяем, что сосед находится в пределах лабиринта
            if 0 <= ni < len(maze) and 0 <= nj < len(maze[0]):
                #Проверяем, что это не стена и клетка еще не посещена
                if maze[ni][nj] != '#' and (ni, nj) not in visited:
                    visited[(ni, nj)] = current  #Запоминаем, откуда пришли
                    queue.append((ni, nj))  #Добавляем в очередь для исследования

    #Восстановление пути (если он существует)
    if end not in visited:
        return None  #Путь не найден

    #Построение пути от конца к началу
    path = []
    current = end
    while current:
        path.append(current)
        current = visited[current]  #Идем назад по цепочке посещений
    path.reverse()  #Переворачиваем, чтобы получить путь от S до E

    return path


def main():
    """Основная функция программы"""
    try:
        #Чтение лабиринта из файла
        with open('input.txt', 'r') as f:
            #Создаем двумерный список (матрицу) лабиринта
            maze = [list(line.strip()) for line in f]

        #Поиск кратчайшего пути
        path = find_shortest_path(maze)

        #Запись результата в файл
        with open('output.txt', 'w') as f:
            if not path:
                f.write("Путь от S до E не существует!\n")
            else:
                #Создаем копию лабиринта для модификации
                maze_copy = [row[:] for row in maze]

                #Помечаем только клетки кратчайшего пути (исключая S и E)
                for i, j in path[1:-1]:
                    maze_copy[i][j] = '*'  #Звездочкой отмечаем путь

                #Записываем результат в файл
                for row in maze_copy:
                    f.write(''.join(row) + '\n')  #Объединяем список в строку

    except FileNotFoundError:
        print("Ошибка: input.txt не найден!")
    except Exception as e:
        print(f"Ошибка: {str(e)}")


if __name__ == "__main__":
    main()  # Запуск программы