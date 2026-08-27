package main

import (
	"fmt"
	"math/rand/v2"
)

// Нужно создать приложение с 2-мя горутинами, где:

// Первая создаёт slice из 10 случайных элементов от 0 до 100 и передаёт их по одному во вторую горутину.
// Вторая получает числа от 1-й и возводит в квадрат передавая результат в main.
// В main дожидаемся всех 10 чисел, которые были возведены в квадрат и выводим их в консоль.

func main() {
	nums := make(chan int)
	results := make(chan int)

	go func() {
		slice := make([]int, 10)
		for i := range slice {
			slice[i] = rand.IntN(101)
		}

		for _, val := range slice {
			nums <- val
		}
		close(nums)
	}()

	go func() {
		for val := range nums {
			square := val * val
			results <- square
		}
		close(results)
	}()

	for res := range results {
		fmt.Println(res)
	}
}
