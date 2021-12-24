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

// Sort sorts a slice by Comparator order.
func Sort[T any](slice []T, comparator Comparator[T]) {
	n := len(slice)
	quickSort(slice, 0, n, maxDepth(n), comparator)
}

func lessThanFor[T any](slice []T, comparator Comparator[T]) func(a, b int) bool {
	return func(a, b int) bool {
		return comparator(slice[a], slice[b]) == LessThan
	}
}

func insertionSort[T any](slice []T, a, b int, comparator Comparator[T]) {
	lessThan := lessThanFor(slice, comparator)
	for i := a + 1; i < b; i++ {
		for j := i; j > a && lessThan(j, j-1); j-- {
			Swap(slice, j, j-1)
		}
	}
}

// siftDown implements the heap property on data[lo:hi].
// first is an offset into the array where the root of the heap lies.
func siftDown[T any](slice []T, lo, hi, first int, comparator Comparator[T]) {
	lessThan := lessThanFor(slice, comparator)
	root := lo
	for {
		child := 2*root + 1
		if child >= hi {
			break
		}
		if child+1 < hi && lessThan(first+child, first+child+1) {
			child++
		}
		if !lessThan(first+root, first+child) {
			return
		}
		Swap(slice, first+root, first+child)
		root = child
	}
}

func heapSort[T any](slice []T, a, b int, comparator Comparator[T]) {
	first := a
	lo := 0
	hi := b - a

	// Build heap with greatest element at top.
	for i := (hi - 1) / 2; i >= 0; i-- {
		siftDown(slice, i, hi, first, comparator)
	}

	// Pop elements, largest first, into end of data.
	for i := hi - 1; i >= 0; i-- {
		Swap(slice, first, first+i)
		siftDown(slice, lo, i, first, comparator)
	}
}

// medianOfThree moves the median of the three values data[m0], data[m1], data[m2] into data[m1].
func medianOfThree[T any](slice []T, m1, m0, m2 int, comparator Comparator[T]) {
	lessThan := lessThanFor(slice, comparator)
	// sort 3 elements
	if lessThan(m1, m0) {
		Swap(slice, m1, m0)
	}
	// data[m0] <= data[m1]
	if lessThan(m2, m1) {
		Swap(slice, m2, m1)
		// data[m0] <= data[m2] && data[m1] < data[m2]
		if lessThan(m1, m0) {
			Swap(slice, m1, m0)
		}
	}
	// now data[m0] <= data[m1] <= data[m2]
}

func doPivot[T any](slice []T, lo, hi int, comparator Comparator[T]) (midlo, midhi int) {
	lessThan := lessThanFor(slice, comparator)
	m := int(uint(lo+hi) >> 1) // Written like this to avoid integer overflow.
	if hi-lo > 40 {
		// Tukey's ``Ninther,'' median of three medians of three.
		s := (hi - lo) / 8
		medianOfThree(slice, lo, lo+s, lo+2*s, comparator)
		medianOfThree(slice, m, m-s, m+s, comparator)
		medianOfThree(slice, hi-1, hi-1-s, hi-1-2*s, comparator)
	}
	medianOfThree(slice, lo, m, hi-1, comparator)

	// Invariants are:
	//	data[lo] = pivot (set up by ChoosePivot)
	//	data[lo < i < a] < pivot
	//	data[a <= i < b] <= pivot
	//	data[b <= i < c] unexamined
	//	data[c <= i < hi-1] > pivot
	//	data[hi-1] >= pivot
	pivot := lo
	a, c := lo+1, hi-1

	for ; a < c && lessThan(a, pivot); a++ {
	}
	b := a
	for {
		for ; b < c && !lessThan(pivot, b); b++ { // data[b] <= pivot
		}
		for ; b < c && lessThan(pivot, c-1); c-- { // data[c-1] > pivot
		}
		if b >= c {
			break
		}
		// data[b] > pivot; data[c-1] <= pivot
		Swap(slice, b, c-1)
		b++
		c--
	}
	// If hi-c<3 then there are duplicates (by property of median of nine).
	// Let's be a bit more conservative, and set border to 5.
	protect := hi-c < 5
	if !protect && hi-c < (hi-lo)/4 {
		// Lets test some points for equality to pivot
		dups := 0
		if !lessThan(pivot, hi-1) { // data[hi-1] = pivot
			Swap(slice, c, hi-1)
			c++
			dups++
		}
		if !lessThan(b-1, pivot) { // data[b-1] = pivot
			b--
			dups++
		}
		// m-lo = (hi-lo)/2 > 6
		// b-lo > (hi-lo)*3/4-1 > 8
		// ==> m < b ==> data[m] <= pivot
		if !lessThan(m, pivot) { // data[m] = pivot
			Swap(slice, m, b-1)
			b--
			dups++
		}
		// if at least 2 points are equal to pivot, assume skewed distribution
		protect = dups > 1
	}
	if protect {
		// Protect against a lot of duplicates
		// Add invariant:
		//	data[a <= i < b] unexamined
		//	data[b <= i < c] = pivot
		for {
			for ; a < b && !lessThan(b-1, pivot); b-- { // data[b] == pivot
			}
			for ; a < b && lessThan(a, pivot); a++ { // data[a] < pivot
			}
			if a >= b {
				break
			}
			// data[a] == pivot; data[b-1] < pivot
			Swap(slice, a, b-1)
			a++
			b--
		}
	}
	// Swap pivot into middle
	Swap(slice, pivot, b-1)
	return b - 1, c
}

func quickSort[T any](slice []T, a, b, maxDepth int, comparator Comparator[T]) {
	lessThan := lessThanFor(slice, comparator)
	for b-a > 12 { // Use ShellSort for slices <= 12 elements
		if maxDepth == 0 {
			heapSort(slice, a, b, comparator)
			return
		}
		maxDepth--
		mlo, mhi := doPivot(slice, a, b, comparator)
		// Avoiding recursion on the larger subproblem guarantees
		// a stack depth of at most lg(b-a).
		if mlo-a < b-mhi {
			quickSort(slice, a, mlo, maxDepth, comparator)
			a = mhi // i.e., quickSort(data, mhi, b)
		} else {
			quickSort(slice, mhi, b, maxDepth, comparator)
			b = mlo // i.e., quickSort(data, a, mlo)
		}
	}
	if b-a > 1 {
		// Do ShellSort pass with gap 6
		// It could be written in this simplified form cause b-a <= 12
		for i := a + 6; i < b; i++ {
			if lessThan(i, i-6) {
				Swap(slice, i, i-6)
			}
		}
		insertionSort(slice, a, b, comparator)
	}
}

// maxDepth returns a threshold at which quicksort should switch
// to heapsort. It returns 2*ceil(lg(n+1)).
func maxDepth(n int) int {
	var depth int
	for i := n; i > 0; i >>= 1 {
		depth++
	}
	return depth * 2
}
