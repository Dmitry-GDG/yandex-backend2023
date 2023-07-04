
// D. Посчитать графы
// 	Все языки 	PHP 7.3.5
// Ограничение времени 	2 секунды 	4 секунды
// Ограничение памяти 	512Mb 	512Mb
// Ввод 	стандартный ввод или input.txt
// Вывод 	стандартный вывод или output.txt
// Назовём неориентированный граф простым, если в нём нет петель и кратных рёбер. Назовём простой неориентированный граф хорошим, если в нём ровно у одной вершины степень равна n−1, то есть в графе есть ровно одна вершина, соединённая со всеми остальными ребром.

// Дано число n, требуется посчитать количество хороших графов на n вершинах. Два графа называются различными, если существует пара вершин (u,v) такая, что в одном графе есть ребро (u,v), а в другом нет.

// Так как ответ может быть крайне большим, выведите отстаток от его деления на 109+7.

// Формат ввода
// В единственной строке задано одно целое число n (1≤n≤5000).
// Формат вывода
// Выведите одно число — ответ на задачу по модулю 109+7.
// 		Ввод	Вывод
// Пример1 	1		1
// Пример2		2		0
// ВПример3	3		3
// Пример4		4		16
// Пример5		2021	113707034
// Пример6		5000	855711688

// Примечания
// В первом примере из условия возможен один единственный граф, и у него ровно одна вершина имеет степень 0:

// PIC

// Во втором примере из условия возможно два графа, но в первом случае ни одна вершина не имеет степень 1, а во втором случае две вершины имеют степень 1:

// PIC

// В третьем примере из условия подходят только такие графы:

// PIC

// В четвёртом примере из условия подходят такие графы:

// PIC