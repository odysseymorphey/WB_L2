package main

import (
	"fmt"
	"sort"
)

// Strategy
type SortingStrategy interface {
	Sort([]int)
}

// ConcreteStrategy
type BubbleSortStrategy struct{}

func (b *BubbleSortStrategy) Sort(numbers []int) {
	fmt.Println("Сортировка пузырьком")
	sort.Slice(numbers, func(i, j int) bool {
		return numbers[i] < numbers[j]
	})
}

// ConcreteStrategy
type QuickSortStrategy struct{}

func (q *QuickSortStrategy) Sort(numbers []int) {
	fmt.Println("Быстрая сортировка")
	sort.Slice(numbers, func(i, j int) bool {
		return numbers[i] < numbers[j]
	})
}

// Context
type Sorter struct {
	strategy SortingStrategy
}

func (s *Sorter) SetStrategy(strategy SortingStrategy) {
	s.strategy = strategy
}

func (s *Sorter) Sort(numbers []int) {
	s.strategy.Sort(numbers)
}

func main() {
	sorter := &Sorter{}

	bubbleSort := &BubbleSortStrategy{}
	sorter.SetStrategy(bubbleSort)
	numbers := []int{5, 2, 7, 3, 9}
	sorter.Sort(numbers)
	fmt.Println("Отсортированный массив:", numbers)

	quickSort := &QuickSortStrategy{}
	sorter.SetStrategy(quickSort)
	numbers = []int{5, 2, 7, 3, 9}
	sorter.Sort(numbers)
	fmt.Println("Отсортированный массив:", numbers)
}

// Паттерн cтратегия - это поведенческий паттерн проектирования, 
// который позволяет определять семейство алгоритмов, инкапсулировать каждый из них 
// и делать их взаимозаменяемыми. Он позволяет изменять поведение объекта во время выполнения программы,
// не изменяя его самого.