/*
Задача №1
Вход:
    расстояние(50 - 10000 км),
    расход в литрах (5-25 литров) на 100 км и
    стоимость бензина(константа) = 48 руб

Выход: стоимость поездки в рублях
*/

package main

import "fmt"

const (
	gasolineСost = 48
)

func main() {
	var distance float64 //расстояние
	var gasCons float64  //расход бензина

	fmt.Println("Введите расстояние(50 - 10000 км):")
	fmt.Scanf("%f\n", &distance)

	fmt.Println("Введите расход в литрах (5-25 литров) на 100 км:")
	fmt.Scanf("%f\n", &gasCons)

	if distance >= 50 && distance <= 10000 && gasCons >= 5 && gasCons <= 25 {
		var sumPrice float64 // суммарная стоимость поездки
		sumPrice = (distance / 100) * gasCons * gasolineСost
		fmt.Printf("При расходе бензина %.2f на расстоянии %.2f, суммарная стоимость поездки равна %.2f.", gasCons, distance, sumPrice)
	} else {
		fmt.Println("операция не поддерживается")
	}
}
