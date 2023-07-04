
// B. Посадка в самолет
// Ограничение времени 	2 секунды
// Ограничение памяти 	512Mb
// Ввод 	стандартный ввод или input.txt
// Вывод 	стандартный вывод или output.txt
// В самолете n рядов и по три кресла слева и справа в каждом ряду. Крайние кресла (A и F) находятся у окна, центральные (C и D) – у прохода. На регистрацию приходят группы из одного, двух или трех пассажиров. Они желают сидеть рядом, то есть на одном ряду и на одной стороне: левой или правой. Например, группа из двух пассажиров может сесть на кресла B и C, но не может сесть на кресла C и D, потому что они разделены проходом, а также не может сесть на кресла A и C, потому что тогда они окажутся не рядом. Кроме того, один из пассажиров каждой группы очень требовательный – он хочет сесть либо у окна, либо у прохода. Конечно же, каждая группа из пассажиров хочет занять места в ряду с как можно меньшим номером, ведь тогда они скорее выйдут из самолета после посадки. Для каждой группы пассажиров определите, есть ли места в самолете, подходящие для них.
// Формат ввода
// Первая строка содержит число n (1≤n≤100) – количество рядов в самолете. Далее в n строках вводится изначальная рассадка в самолете по рядам (от первого до n-го), где символами . (точка) обозначены свободные места, символами # (решетка) обозначены занятые места, а символами _ (нижнее подчеркивание) обозначен проход между креслами C и D каждого ряда.

// Следующая строка содержит число m (1≤m≤100) – количество групп пассажиров. Далее в m строках содержатся описания групп пассажиров. Формат описания такой: numsideposition, где num – количество пассажиров (число 1, 2 или 3), side – желаемая сторона самолета (строка left или right), position – желаемое место требовательного пассажира (строка aisle или window).
// Формат вывода
// Если группа может сесть на места, удовлетворяющие ее требованиям, то выведите строку Passengers can take seats: и список их мест в формате rowletter, упорядоченный по возрастанию буквы места. Затем выведите в n строках получившуюся рассадку в самолете, в формате, описанном выше, причем места, занятые текущей группой пассажиров, должны быть обозначены символом X.

// Если группа не может найти места, удовлетворяющие ее требованиям, то выведите строку Cannot fulfill passengers requirements.

// Ответ сравнивается с правильным посимвольно, поэтому ваше решение не должно выводить никаких лишних символов, в том числе лишних переводов строк или пробельных символов в концах строк. В конце каждой строки (включая последнюю) должен быть выведен символ перевода строки.

// Пример
// Ввод

// 4
// ..._.#.
// .##_...
// .#._.##
// ..._...
// 7
// 2 left aisle
// 3 right window
// 2 left window
// 3 left aisle
// 1 right window
// 2 right window
// 1 right window

// Вывод
// Passengers can take seats: 1B 1C
// .XX_.#.
// .##_...
// .#._.##
// ..._...
// Passengers can take seats: 2D 2E 2F
// .##_.#.
// .##_XXX
// .#._.##
// ..._...
// Passengers can take seats: 4A 4B
// .##_.#.
// .##_###
// .#._.##
// XX._...
// Cannot fulfill passengers requirements
// Passengers can take seats: 1F
// .##_.#X
// .##_###
// .#._.##
// ##._...
// Passengers can take seats: 4E 4F
// .##_.##
// .##_###
// .#._.##
// ##._.XX
// Cannot fulfill passengers requirements

package main

import "fmt"

func myPrint(s *[100][6]int, qtyRows int) {
	for i:= 0; i < qtyRows; i++ {
		for j := 0; j < 3; j++ {
			if (*s)[i][j] == 0 {
				fmt.Print(".")
			} else if (*s)[i][j] == 1 {
				fmt.Print("#")
			} else {
				fmt.Print("X")
				(*s)[i][j] = 1
			}
		}
		fmt.Print("_")
		for j := 3; j < 6; j++ {
			if (*s)[i][j] == 0 {
				fmt.Print(".")
			} else if (*s)[i][j] == 1 {
				fmt.Print("#")
			} else {
				fmt.Print("X")
				(*s)[i][j] = 1
			}
		}
		fmt.Print("\n")
	}
}

func main() {
	var qtyRows, qtyPass  int
	var seatsOccupied [100][6]int
	var passWishes [100][3]int // qty Pass, 0/1 (left/right), 0/1 (windows/aisle)
	var passTmp int
	var tmp string
	var sign, signPrint bool
	fmt.Scan(&qtyRows)
	for i := 0; i < qtyRows; i++ {
		fmt.Scan(&tmp)
		for j := 0; j < 3; j++ {
			if tmp[j] == '.' {
				seatsOccupied[i][j] = 0
			} else if tmp[j] == '#' {
				seatsOccupied[i][j] = 1
			}
		}
		for j := 4; j < 7; j++ {
			if tmp[j] == '.' {
				seatsOccupied[i][j - 1] = 0
			} else if tmp[j] == '#' {
				seatsOccupied[i][j - 1] = 1
			}
		}
	}
	// fmt.Println("")
	// myPrint(&seatsOccupied, qtyRows)
	fmt.Scan(&qtyPass)
	for i := 0; i < qtyPass; i++ {
		fmt.Scan(&passTmp)
		passWishes[i][0] = passTmp
		fmt.Scan(&tmp)
		if (tmp[0] == 'l') {
			passWishes[i][1] = 0
		} else {
			passWishes[i][1] = 1
		}
		fmt.Scan(&tmp)
		if (tmp[0] == 'w') {
			passWishes[i][2] = 0
		} else {
			passWishes[i][2] = 1
		}
	}
	// fmt.Println("")
	// fmt.Println(passWishes)
	// fmt.Println("")
	for i := 0; i < qtyPass; i++ {
		if passWishes[i][1] == 0 && passWishes[i][2] == 0{ //left-window
			signPrint = false
			for j := 0; j < qtyRows && !signPrint; j++ {
				sign = false
				if seatsOccupied[j][0] != 0 {
					continue
				}
				for k := 0; k < passWishes[i][0]; k++ {
					if seatsOccupied[j][k] != 0 {
						sign = true
					}
				}
				if !sign {
					for l := 0; l < passWishes[i][0]; l++ {
						seatsOccupied[j][l] = 'X'
					}
					switch passWishes[i][0] {
						case 1: fmt.Printf("Passengers can take seats: %dA\n", j + 1)
						case 2: fmt.Printf("Passengers can take seats: %dA %dB\n", j + 1, j + 1)
						case 3: fmt.Printf("Passengers can take seats: %dA %dB %dC\n", j + 1, j + 1, j + 1)
					}
					myPrint(&seatsOccupied, qtyRows)
					signPrint = true
				}
			}
			if !signPrint {
				fmt.Println("Cannot fulfill passengers requirements")
			}
		}
		if passWishes[i][1] == 1 && passWishes[i][2] == 1 {  //right-aisle
			signPrint = false
			for j := 0; j < qtyRows && !signPrint; j++ {
				sign = false
				if seatsOccupied[j][3] != 0 {
					continue
				}
				for k := 0; k < passWishes[i][0]; k++ {
					if seatsOccupied[j][k + 3] != 0 {
						sign = true
					}
				}
				if !sign {
					for l := 0; l < passWishes[i][0]; l++ {
						seatsOccupied[j][l + 3] = 'X'
					}
					switch passWishes[i][0] {
						case 1: fmt.Printf("Passengers can take seats: %dD\n", j + 1)
						case 2: fmt.Printf("Passengers can take seats: %dD %dE\n", j + 1, j + 1)
						case 3: fmt.Printf("Passengers can take seats: %dD %dE %dF\n", j + 1, j + 1, j + 1)
					}
					myPrint(&seatsOccupied, qtyRows)
					signPrint = true
				}
			}
			if !signPrint {
				fmt.Println("Cannot fulfill passengers requirements")
			}
		}
		if passWishes[i][1] == 0 && passWishes[i][2] == 1 { //left-aisle
			signPrint = false
			for j := 0; j < qtyRows && !signPrint; j++ {
				sign = false
				if seatsOccupied[j][2] != 0 {
					continue
				}
				for k := 0; k < passWishes[i][0]; k++ {
					// fmt.Println(j, k)
					if seatsOccupied[j][2 - k] != 0 {
						sign = true
					}
				}
				if !sign {
					for l := 0; l < passWishes[i][0]; l++ {
						seatsOccupied[j][2 - l] = 'X'
					}
					switch passWishes[i][0] {
						case 1: fmt.Printf("Passengers can take seats: %dC\n", j + 1)
						case 2: fmt.Printf("Passengers can take seats: %dB %dC\n", j + 1, j + 1)
						case 3: fmt.Printf("Passengers can take seats: %dA %dB %dC\n", j + 1, j + 1, j + 1)
					}
					myPrint(&seatsOccupied, qtyRows)
					signPrint = true
				}
		}
			if !signPrint {
				fmt.Println("Cannot fulfill passengers requirements")
			}
		}
		if passWishes[i][1] == 1 && passWishes[i][2] == 0 { //right-window
			signPrint = false
			for j := 0; j < qtyRows && signPrint == false; j++ {
				sign = false
				if seatsOccupied[j][5] != 0 {
					continue
				}
				for k := 0; k < passWishes[i][0]; k++ {
					if seatsOccupied[j][5 - k] != 0 {
						sign = true
					}
				}
				if !sign {
					for l := 0; l < passWishes[i][0]; l++ {
						seatsOccupied[j][5 - l] = 'X'
					}
					switch passWishes[i][0] {
						case 1: fmt.Printf("Passengers can take seats: %dF\n", j + 1)
						case 2: fmt.Printf("Passengers can take seats: %dE %dF\n", j + 1, j + 1)
						case 3: fmt.Printf("Passengers can take seats: %dD %dE %dF\n", j + 1, j + 1, j + 1)
					}
					myPrint(&seatsOccupied, qtyRows)
					signPrint = true
				}
			}
			if !signPrint {
				fmt.Println("Cannot fulfill passengers requirements")
			}
		}
	}
}

