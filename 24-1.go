package main

import "fmt"

//Напишите функцию, сортирующую массив длины 10 вставками.
func sortingInsert(input []int) {
	for i := 1; i < len(input); i++ {
		b := i
		for b > 0 && input[b-1] > input[b] {
			swap(input, i, b)
			b--
		}
	}
}

func swap(input []int, i, b int) {
	tmp := input[b-1]
	input[b-1] = input[b]
	input[b] = tmp
}

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
	fmt.Printf("Вашь массив : %d\n", arr)
	return arr
}

func main() {
	arr := creatArr()
	sortingInsert(arr)
	fmt.Printf("Отсортированный массив: %d\n", arr)
}
