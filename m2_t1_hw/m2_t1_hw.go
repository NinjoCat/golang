/*
Задача № 3. Вывести на экран в порядке возрастания три введенных числа
Пример:
Вход: 1, 9, 2
Выход: 1, 2, 9
*/

package main

import (
	"fmt"
)

func main() {

	var x int
	var y int
	var z int

	fmt.Println("Введите целое чило (1):")
	fmt.Scanf("%d\n", &x)

	fmt.Println("Введите целое чило (2):")
	fmt.Scanf("%d\n", &y)

	fmt.Println("Введите целое чило (3):")
	fmt.Scanf("%d\n", &z)

	fmt.Printf("Вы ввели числа: %d, %d, %d \n", x, y, z)


	if x > y {
		cache := x
		x = y
		y = cache
	}

	if x > z {
		cache := x
		x = z
		z = cache
	}

	if y > z {
		cache := y
		y = z
		z = cache
	}

	fmt.Printf("Результат сортировки чисел:   %d, %d, %d \n", x, y, z)
}
