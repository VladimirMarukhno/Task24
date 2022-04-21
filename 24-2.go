package main

import "fmt"

//Напишите анонимную функцию, которая на вход получает массив типа integer,
//сортирует его пузырьком и переворачивает (либо сразу сортирует в обратном порядке, как посчитаете нужным).

func creatArr() []int {
	var size, elem int
	fmt.Println("Введите количество элементов в массиве.")
	_, err := fmt.Scan(&size)
	if err != nil {
		fmt.Println(err)
		defer creatArr()
	} else if size <= 0 {
		fmt.Printf("Количество элементво в массиве(%d) не может быть меньше или равно 0.\n", size)
	}
	arr := make([]int, size, size)
	for i := 0; i < size; i++ {
		fmt.Println("Введите следующий элемент.")
		_, err = fmt.Scan(&elem)
		if err != nil {
			fmt.Println(err)
			i--
		}
		arr[i] = elem
	}
	fmt.Printf("Ваш массив : %d\n", arr)
	return arr
}
func swap(ar []int, i, j int) {
	tmp := ar[i]
	ar[i] = ar[j]
	ar[j] = tmp
}

func main() {
	arr := creatArr()
	bubbleSort := func(input []int) {
		for i := 0; i < len(input); i++ {
			for j := i + 1; j < len(input); j++ {
				if input[i] < input[j] {
					swap(input, i, j)
				}
			}
		}
	}
	bubbleSort(arr)
	fmt.Printf("Отсортированный массив: %d\n", arr)
}
