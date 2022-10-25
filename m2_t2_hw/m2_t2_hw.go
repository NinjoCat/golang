/*
Задача № 4. Проверить, является ли четырехзначное число палиндромом
Пример:
Вход: 1221  Выход: 1221 - палиндром
Вход: 1234  Выход: 1234 - не палиндром
*/

package main

import (
	"fmt"
)

func main() {
	var dig int      //4х значное число
	var firstNum int // первая цифра
	var secNum int   // вторая цифра
	var thirdNum int // третья цифра
	var forthNum int // четвертая цифра

	var invertDig int // реверсивная запись числа

	fmt.Println("Введите четырехзначное целое число:")
	fmt.Scanf("%d\n", &dig)

	//1234
	if dig >= 1000 && dig <= 10000 {
		firstNum = dig / 1000                              // тысячи
		secNum = (dig - firstNum*1000) / 100               // сотни
		thirdNum = (dig - firstNum*1000 - secNum*100) / 10 // десятки
		forthNum = (dig - firstNum*1000 - secNum*100) % 10 // единицы
		fmt.Printf("Тысячи %d, Сотни %d, Десятки %d, Единицы %d \n", firstNum, secNum, thirdNum, forthNum)

		invertDig = forthNum*1000 + thirdNum*100 + secNum*10 + firstNum

		if invertDig == dig {
			fmt.Println("Успех! Число является пандромом")
		} else {
			fmt.Println("Число не является пандромом")
		}
	} else {
		fmt.Println("Операция не поддерживается.")
	}
}
