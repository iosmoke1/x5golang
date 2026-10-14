package main

import "fmt"

func OrderCategory(n int) (cat string, ok bool) {
	if n < 0 {
		return
	}
	if n < 10000 {
		cat = "small"
	} else if n < 50000 {
		cat = "regular"
	} else {
		cat = "large"
	}
	ok = true
	return
}

func main() {
	n := 5000
	cat, ok := OrderCategory(n)
	if !ok {
		fmt.Println("Ошибка: отрицательная сумма")
	} else {
		fmt.Printf("Категория заказа: %s\n", cat)
	}
	n = 15000
	cat, ok = OrderCategory(n)
	if !ok {
		fmt.Println("Ошибка: отрицательная сумма")
	} else {
		fmt.Printf("Категория заказа: %s\n", cat)
	}
	n = 55000
	cat, ok = OrderCategory(n)
	if !ok {
		fmt.Println("Ошибка: отрицательная сумма")
	} else {
		fmt.Printf("Категория заказа: %s\n", cat)
	}
	n = -5000
	cat, ok = OrderCategory(n)
	if !ok {
		fmt.Println("Ошибка: отрицательная сумма")
	} else {
		fmt.Printf("Категория заказа: %s\n", cat)
	}
}
