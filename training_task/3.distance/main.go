
// C. Расстояние
// Язык 	Ограничение времени 	Ограничение памяти 	Ввод 	Вывод
// Все языки 	2 секунды 	512Mb 	стандартный ввод или input.txt 	стандартный вывод или output.txt
// Python 3.7.3 	4 секунды 	512Mb
// Python 3.7 (PyPy 7.3.3) 	4 секунды 	512Mb
// Python 2.7 	4 секунды 	512Mb
// PHP 7.3.5 	4 секунды 	512Mb

// Рассмотрим целочисленный массив a длины n. 
// Назовём расстоянием от индекса i до множества индексов S величину dist(i,S)=∑j∈S∣∣ai−aj∣∣.

// Зафиксируем целое число k. Рассмотрим функцию f(i)=mindist(i,S), 
// где минимум берётся по множествам S размера k, не содержащим индекс i.

// Определите значение f(i) для всех i от 1 до n.

// Формат ввода
// В первой строке заданы два целых числа n и k (2≤n≤300000, 1≤k<n), описанные в условии.

// Во второй строке содержится n целых чисел ai (1≤ai≤109) — элементы массива a.
// Формат вывода
// Выведите n целых чисел: значения f(i) для i=1,i=2,…,i=n.
// Пример 1
// Ввод	4 2
// 		1 2 3 4
// Вывод	3 2 2 3

// Пример 2
// Ввод	5 3
// 		3 2 5 1 2
// Вывод	4 2 8 4 2

// Пример 3
// Ввод	6 2
// 		3 2 1 101 102 103
// Вывод	 2 3 3 2 3


// Примечания
// Рассмотрим первый пример.

// При i=1 оптмиальное S — это {2,3}; dist(1,{2,3})=|1−2|+|1−3|=3.

// При i=2 оптмиальное S — это {1,3}; dist(2,{1,3})=|2−1|+|2−3|=2.

// При i=3 оптмиальное S — это {2,4}; dist(3,{2,4})=|3−2|+|3−4|=2.

// При i=4 оптмиальное S — это {2,3}; dist(4,{2,3})=|4−2|+|4−3|=3. 