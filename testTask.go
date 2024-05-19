package main

import (
	"fmt"
	"sort"
)

func canSortBalls(n int, containers [][]int) string {
	// Вычисляем общее количество шаров каждого цвета и общую вместимость каждого контейнера
	totalBallsByColor := make([]int, n)
	totalCapacityByContainer := make([]int, n)

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			totalBallsByColor[j] += containers[i][j]
			totalCapacityByContainer[i] += containers[i][j]
		}
	}

	// Проверим, можно ли отсортировать
	sort.Ints(totalBallsByColor)
	sort.Ints(totalCapacityByContainer)

	for i := 0; i < n; i++ {
		if totalBallsByColor[i] != totalCapacityByContainer[i] {
			return "no"
		}
	}

	return "yes"
}

func main() {
	var n int
	fmt.Scan(&n)

	containers := make([][]int, n)
	for i := 0; i < n; i++ {
		containers[i] = make([]int, n)
		for j := 0; j < n; j++ {
			fmt.Scan(&containers[i][j])
		}
	}

	fmt.Println(canSortBalls(n, containers))
}
