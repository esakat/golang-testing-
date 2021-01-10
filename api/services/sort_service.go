package services

import "golang-testing/api/utils/sort"

func Sort(elements []int) {
	sort.BubbleSort(elements)
}