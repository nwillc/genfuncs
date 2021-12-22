/*
 *  Copyright (c) 2021,  nwillc@gmail.com
 *
 *  Permission to use, copy, modify, and/or distribute this software for any
 *  purpose with or without fee is hereby granted, provided that the above
 *  copyright notice and this permission notice appear in all copies.
 *
 *  THE SOFTWARE IS PROVIDED "AS IS" AND THE AUTHOR DISCLAIMS ALL WARRANTIES
 *  WITH REGARD TO THIS SOFTWARE INCLUDING ALL IMPLIED WARRANTIES OF
 *  MERCHANTABILITY AND FITNESS. IN NO EVENT SHALL THE AUTHOR BE LIABLE FOR
 *  ANY SPECIAL, DIRECT, INDIRECT, OR CONSEQUENTIAL DAMAGES OR ANY DAMAGES
 *  WHATSOEVER RESULTING FROM LOSS OF USE, DATA OR PROFITS, WHETHER IN AN
 *  ACTION OF CONTRACT, NEGLIGENCE OR OTHER TORTIOUS ACTION, ARISING OUT OF
 *  OR IN CONNECTION WITH THE USE OR PERFORMANCE OF THIS SOFTWARE.
 */

package genfuncs

// InsertionSort sorts a slice by Comparator order using the insertion sort algorithm.
func InsertionSort[T any](slice []T, comparator Comparator[T]) {
	for i := 1; i < len(slice); i++ {
		key := slice[i]
		j := i - 1
		for ; j >= 0 && comparator(slice[j], key) == GreaterThan; j-- {
			slice[j+1] = slice[j]
		}
		slice[j+1] = key
	}
}

// HeapSort sorts a slice by Comparator order using the heap sort algorithm.
func HeapSort[T any](slice []T, comparator Comparator[T]) {
	// Build max heap
	n := len(slice)
	for i := n/2 - 1; i >= 0; i-- {
		heapify(slice, n, i, comparator)
	}

	// Heap sort
	for i := n - 1; i >= 0; i-- {
		Swap(slice, 0, i)
		// Heapify root element to get highest element at root again
		heapify(slice, i, 0, comparator)
	}
}

func heapify[T any](slice []T, n, i int, comparator Comparator[T]) {
	// Find largest among root, left child and right child
	largest := i
	left := left(i)
	right := right(i)

	if left < n && comparator(slice[left], slice[largest]) == GreaterThan {
		largest = left
	}
	if right < n && comparator(slice[right], slice[largest]) == GreaterThan {
		largest = right
	}

	// Swap and continue heapifying if root is not largest
	if largest != i {
		Swap(slice, i, largest)
		heapify(slice, n, largest, comparator)
	}
}

// QuickSort sorts a slice by Comparator order using the quick sort algorithm.
func QuickSort[T any](slice []T, comparator Comparator[T]) {
	quickSort(slice, 0, len(slice)-1, comparator)
}

func quickSort[T any](slice []T, start, end int, comparator Comparator[T]) {
	// base condition
	if start >= end {
		return
	}

	// rearrange elements across pivot
	pivot := partition(slice, start, end, comparator)

	// recur on subarray containing elements less than the pivot
	quickSort(slice, start, pivot-1, comparator)

	// recur on subarray containing elements more than the pivot
	quickSort(slice, pivot+1, end, comparator)
}

// Partition using the Lomuto partition scheme
func partition[T any](slice []T, start, end int, comparator Comparator[T]) int {

	// Pick the rightmost element as a pivot from the array
	pivot := slice[end]

	// elements less than the pivot will be pushed to the left of `pIndex`
	// elements more than the pivot will be pushed to the right of `pIndex`
	// equal elements can go either way
	pIndex := start

	// each time we find an element less than or equal to the pivot,
	// `pIndex` is incremented, and that element would be placed
	// before the pivot.
	for i := start; i < end; i++ {
		if comparator(slice[i], pivot) != GreaterThan {
			Swap(slice, i, pIndex)
			pIndex++
		}
	}

	// swap `pIndex` with pivot
	Swap(slice, end, pIndex)

	// return `pIndex` (index of the pivot element)
	return pIndex
}
