package main

import (
	"errors"
	"fmt"
)

func main() {
	var a, b float64
	fmt.Println("Введите два числа через пробел: ")
	_, err := fmt.Scan(&a, &b)
	if err != nil {
		fmt.Println("Ошибка ввода: ", err)
		return
	}
	res, err := divide(a, b)
	if err != nil {
		fmt.Println("Ошибка: ", err)
		return
	} else {
		fmt.Printf("%.2f\n", res)
	}
}

func divide(a float64, b float64) (float64, error) { // Явное лучше, чем неявное.
	var err error
	if b == 0 {
		err = errors.New("Делить на ноль нельзя!")
		return 0, err
	}
	result := a / b
	return result, err
}
