package algorithms

import (
	"math"
	"slices"
)

type Ordered interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr |
		~float32 | ~float64 |
		~string
}

const NumDigits = 10

// Like selection sort, not optimized
func SimpleSort[T Ordered](vec []T) {
	// Can go to len(vec)-1 because no need for processing at last index
	// But you can leave it out to make code seem simpler too!
	for i := 0; i < len(vec)-1; i++ {
		for j := i + 1; j < len(vec); j++ {
			if vec[j] < vec[i] {
				vec[i], vec[j] = vec[j], vec[i]
			}
		}
	}
}

// More optimized selection sort. It keeps track of
// smallest value and only swaps once. This is better
// especially if sorting large arrays

// "Select the Min"
func SelectionSort[T Ordered](vec []T) {
	// Can go to len(vec)-1 here because last one will already
	// be correct when everything else is sorted
	for i := 0; i < len(vec)-1; i++ {
		minIndex := i
		for j := i + 1; j < len(vec); j++ {
			if vec[j] < vec[minIndex] {
				minIndex = j
			}
		}
		if minIndex != i {
			vec[i], vec[minIndex] = vec[minIndex], vec[i]
		}
	}
}

// "Bubble the max to the right"
func BubbleSort[T Ordered](vec []T) {
	for i := 0; i < len(vec)-1; i++ {
		swapped := false
		for j := 0; j < len(vec)-1-i; j++ {
			if vec[j] > vec[j+1] {
				vec[j], vec[j+1] = vec[j+1], vec[j]
				swapped = true
			}
		}

		if !swapped {
			break
		}
	}
}

// Insert each new element to the sorted range in the left
func InsertionSort[T Ordered](vec []T) {
	// First element is already sorted
	for i := 1; i < len(vec); i++ {
		for j := i; j > 0 && vec[j] < vec[j-1]; j-- {
			vec[j], vec[j-1] = vec[j-1], vec[j]
		}
	}
}

// Divide and conquer! Divide into two parts and then do the work!
func MergeSort[T Ordered](vec []T) {
	// Instantly return because you don't want to do any of that extra work
	// This is critical actually
	if len(vec) <= 1 {
		return
	}

	tmp := make([]T, len(vec))
	mergeSortHelper(vec, tmp, 0, len(vec)-1)
}

func mergeSortHelper[T Ordered](vec []T, tmp []T, start int, end int) {
	if start >= end {
		return
	}

	mid := start + (end-start)/2
	mergeSortHelper(vec, tmp, start, mid)
	mergeSortHelper(vec, tmp, mid+1, end)
	merge(vec, tmp, start, mid, end)
}

func merge[T Ordered](vec []T, tmp []T, start int, mid int, end int) {
	i, j, k := start, mid+1, start

	for i <= mid && j <= end {
		if vec[i] <= vec[j] {
			tmp[k] = vec[i]
			i++
		} else {
			tmp[k] = vec[j]
			j++
		}
		k++
	}

	for i <= mid {
		tmp[k] = vec[i]
		i++
		k++
	}

	for j <= end {
		tmp[k] = vec[j]
		j++
		k++
	}

	for i = start; i <= end; i++ {
		vec[i] = tmp[i]
	}
}

// Pick a pivot, then fix the vector such that pivot is in its correct
// position and everything to its left is <= than itself, and everything
// to its right is > than itself
func QuickSort[T Ordered](vec []T) {
	if len(vec) <= 1 {
		return
	}

	quickSortHelper(vec, 0, len(vec)-1)
}

func quickSortHelper[T Ordered](vec []T, start int, end int) {
	if start >= end {
		return
	}

	pivot := partition(vec, start, end)
	quickSortHelper(vec, start, pivot-1)
	quickSortHelper(vec, pivot+1, end)
}

func partition[T Ordered](vec []T, start int, end int) int {
	mid := start + (end-start)/2
	pivotIndex := medianOfThree(vec, start, mid, end)
	vec[pivotIndex], vec[end] = vec[end], vec[pivotIndex]

	pivot := vec[end]
	i := start - 1

	for j := start; j < end; j++ {
		if vec[j] <= pivot {
			i++
			vec[i], vec[j] = vec[j], vec[i]
		}
	}

	vec[i+1], vec[end] = vec[end], vec[i+1]
	return i + 1
}

func medianOfThree[T Ordered](vec []T, i, j, k int) int {
	if (vec[i] > vec[j]) != (vec[i] > vec[k]) {
		return i
	} else if (vec[j] > vec[i]) != (vec[j] > vec[k]) {
		return j
	} else {
		return k
	}
}

// Use a max-heap and then remove the first element one by one, put it at the end
// Then fix the rest using heapify
func HeapSort[T Ordered](vec []T) {
	n := len(vec)
	buildHeap(vec)
	for i := n - 1; i >= 0; i-- {
		vec[0], vec[i] = vec[i], vec[0]
		heapify(vec, 0, i)
	}
}

func buildHeap[T Ordered](vec []T) {
	n := len(vec)
	for i := n/2 - 1; i >= 0; i-- {
		heapify(vec, i, n)
	}
}

// n needed to heapify a subset!
func heapify[T Ordered](vec []T, i int, n int) {
	largest := i
	left := 2*i + 1
	right := 2*i + 2

	if left < n && vec[left] > vec[largest] {
		largest = left
	}

	if right < n && vec[right] > vec[largest] {
		largest = right
	}

	if largest != i {
		vec[i], vec[largest] = vec[largest], vec[i]
		heapify(vec, largest, n)
	}
}

// maxVal here is the maximum value in the array
// i.e., the number of discinct values to be counted
func GeneralCountingSort(vec []uint) {
	if len(vec) <= 1 {
		return
	}

	max := slices.Max(vec)

	counts := make([]uint, max+1)
	sorted := make([]uint, len(vec))

	for _, val := range vec {
		counts[val]++
	}

	for i := 1; i < len(counts); i++ {
		counts[i] += counts[i-1]
	}

	for i := len(vec) - 1; i >= 0; i-- {
		sorted[counts[vec[i]]-1] = vec[i]
		counts[vec[i]]--
	}

	copy(vec, sorted)
}

func IntegerCountingSort(vec []uint) {
	if len(vec) <= 1 {
		return
	}

	max := slices.Max(vec)
	counts := make([]uint, max+1)

	for _, val := range vec {
		counts[val]++
	}

	index := 0
	var i uint
	for i = 0; i < uint(len(counts)); i++ {
		for counts[i] > 0 {
			vec[index] = i
			counts[i]--
			index++
		}
	}
}

func IntRadixSort(vec []uint) {
	if len(vec) <= 1 {
		return
	}

	max := slices.Max(vec)
	var exp uint = 1

	for (max / exp) > 0 {
		radixIntCountSort(vec, exp)
		exp *= 10
	}
}

func radixIntCountSort(vec []uint, exp uint) {
	output := make([]uint, len(vec))
	counts := make([]uint, NumDigits)

	for i := 0; i < len(vec); i++ {
		bucket := (vec[i] / exp) % NumDigits
		counts[bucket]++
	}

	for i := uint(1); i < NumDigits; i++ {
		counts[i] += counts[i-1]
	}

	for i := len(vec) - 1; i >= 0; i-- {
		bucket := (vec[i] / exp) % NumDigits
		output[counts[bucket]-1] = vec[i]
		counts[bucket]--
	}

	copy(vec, output)
}

func LessEfficientRadixSort(vec []uint) {
	if len(vec) <= 1 {
		return
	}

	max := slices.Max(vec)
	var divisor uint = 1

	for (max / divisor) > 0 {
		radixArray := make([][]uint, NumDigits)
		for i := 0; i < len(vec); i++ {
			radixIndex := (vec[i] / divisor) % NumDigits
			radixArray[radixIndex] = append(radixArray[radixIndex], vec[i])
		}

		k := 0
		for i := range radixArray {
			for j := range radixArray[i] {
				vec[k] = radixArray[i][j]
				k++
			}
			radixArray[i] = nil
		}

		divisor *= 10
	}
}

// REQUIRES: ASCII strings
func StringRadixSort(vec []string) {
	if len(vec) <= 1 {
		return
	}

	maxLen := len(vec[0])
	for i := 1; i < len(vec); i++ {
		if len(vec[i]) > maxLen {
			maxLen = len(vec[i])
		}
	}

	for i := maxLen - 1; i >= 0; i-- {
		radixStringCountSort(vec, i)
	}
}

func radixStringCountSort(vec []string, curIdx int) {
	output := make([]string, len(vec))
	counts := make([]uint, 129)

	var bucket uint8
	for i := 0; i < len(vec); i++ {
		if curIdx < len(vec[i]) {
			bucket = uint8(vec[i][curIdx]) + 1
		} else {
			bucket = 0 // for shorter strings
		}
		counts[bucket]++
	}

	for i := 1; i < len(counts); i++ {
		counts[i] += counts[i-1]
	}

	for i := len(vec) - 1; i >= 0; i-- {
		if curIdx < len(vec[i]) {
			bucket = uint8(vec[i][curIdx]) + 1
		} else {
			bucket = 0 // for shorter strings
		}
		output[counts[bucket]-1] = vec[i]
		counts[bucket]--
	}

	copy(vec, output)
}

func BucketSort(vec []float64) {
	if len(vec) <= 1 {
		return
	}

	min := math.Inf(1)
	max := math.Inf(-1)

	for _, val := range vec {
		// need to if's because edge case when first element is the max!

		if val < min {
			min = val
		}

		if val > max {
			max = val
		}
	}

	// edge case when no need for buckets! simply quicksort.
	if max == min {
		QuickSort(vec)
		return
	}

	numBuckets := int((max-min)/math.Sqrt(float64(len(vec)))) + 1 // need +1 here!
	buckets := make([][]float64, int(numBuckets))

	for _, val := range vec {
		index := int((val - min) / (max - min) * float64(numBuckets-1))
		buckets[index] = append(buckets[index], val)
	}

	output := make([]float64, len(vec))
	k := 0

	for i := 0; i < len(buckets); i++ {
		QuickSort(buckets[i])
		for _, val := range buckets[i] {
			output[k] = val
			k++
		}
	}

	copy(vec, output)
}
