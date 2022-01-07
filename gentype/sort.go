/*
 *  Copyright (c) 2022,  nwillc@gmail.com
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

package gentype

import "github.com/nwillc/genfuncs"

// Sort sorts a slice by the LessThan order.
func (s Slice[T]) Sort(lessThan genfuncs.LessThan[T]) {
	n := len(s)
	s.quickSort(0, n, maxDepth(n), lessThan)
}

func (s Slice[T]) insertionSort(a, b int, lessThan genfuncs.LessThan[T]) {
	for i := a + 1; i < b; i++ {
		for j := i; j > a && lessThan(s[j], s[j-1]); j-- {
			s.Swap(j, j-1)
		}
	}
}

// siftDown implements the heap property on data[lo:hi].
// first is an offset into the array where the root of the heap lies.
func (s Slice[T]) siftDown(lo, hi, first int, lessThan genfuncs.LessThan[T]) {
	root := lo
	for {
		child := 2*root + 1
		if child >= hi {
			break
		}
		if child+1 < hi && lessThan(s[first+child], s[first+child+1]) {
			child++
		}
		if !lessThan(s[first+root], s[first+child]) {
			return
		}
		s.Swap(first+root, first+child)
		root = child
	}
}

func (s Slice[T]) heapSort(a, b int, lessThan genfuncs.LessThan[T]) {
	first := a
	lo := 0
	hi := b - a

	// Build heap with greatest element at top.
	for i := (hi - 1) / 2; i >= 0; i-- {
		s.siftDown(i, hi, first, lessThan)
	}

	// Pop elements, largest first, into end of data.
	for i := hi - 1; i >= 0; i-- {
		s.Swap(first, first+i)
		s.siftDown(lo, i, first, lessThan)
	}
}

// medianOfThree moves the median of the three values data[m0], data[m1], data[m2] into data[m1].
func (s Slice[T]) medianOfThree(m1, m0, m2 int, lessThan genfuncs.LessThan[T]) {
	// sort 3 elements
	if lessThan(s[m1], s[m0]) {
		s.Swap(m1, m0)
	}
	// data[m0] <= data[m1]
	if lessThan(s[m2], s[m1]) {
		s.Swap(m2, m1)
		// data[m0] <= data[m2] && data[m1] < data[m2]
		if lessThan(s[m1], s[m0]) {
			s.Swap(m1, m0)
		}
	}
	// now data[m0] <= data[m1] <= data[m2]
}

func (s Slice[T]) doPivot(lo, hi int, lessThan genfuncs.LessThan[T]) (midlo, midhi int) {
	m := int(uint(lo+hi) >> 1) // Written like this to avoid integer overflow.
	if hi-lo > 40 {
		// Tukey's ``Ninther,'' median of three medians of three.
		ss := (hi - lo) / 8
		s.medianOfThree(lo, lo+ss, lo+2*ss, lessThan)
		s.medianOfThree(m, m-ss, m+ss, lessThan)
		s.medianOfThree(hi-1, hi-1-ss, hi-1-2*ss, lessThan)
	}
	s.medianOfThree(lo, m, hi-1, lessThan)

	// Invariants are:
	//	data[lo] = pivot (set up by ChoosePivot)
	//	data[lo < i < a] < pivot
	//	data[a <= i < b] <= pivot
	//	data[b <= i < c] unexamined
	//	data[c <= i < hi-1] > pivot
	//	data[hi-1] >= pivot
	pivot := lo
	a, c := lo+1, hi-1

	for ; a < c && lessThan(s[a], s[pivot]); a++ {
	}
	b := a
	for {
		for ; b < c && !lessThan(s[pivot], s[b]); b++ { // data[b] <= pivot
		}
		for ; b < c && lessThan(s[pivot], s[c-1]); c-- { // data[c-1] > pivot
		}
		if b >= c {
			break
		}
		// data[b] > pivot; data[c-1] <= pivot
		s.Swap(b, c-1)
		b++
		c--
	}
	// If hi-c<3 then there are duplicates (by property of median of nine).
	// Let's be a bit more conservative, and set border to 5.
	protect := hi-c < 5
	if !protect && hi-c < (hi-lo)/4 {
		// Lets test some points for equality to pivot
		dups := 0
		if !lessThan(s[pivot], s[hi-1]) { // data[hi-1] = pivot
			s.Swap(c, hi-1)
			c++
			dups++
		}
		if !lessThan(s[b-1], s[pivot]) { // data[b-1] = pivot
			b--
			dups++
		}
		// m-lo = (hi-lo)/2 > 6
		// b-lo > (hi-lo)*3/4-1 > 8
		// ==> m < b ==> data[m] <= pivot
		if !lessThan(s[m], s[pivot]) { // data[m] = pivot
			s.Swap(m, b-1)
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
			for ; a < b && !lessThan(s[b-1], s[pivot]); b-- { // data[b] == pivot
			}
			for ; a < b && lessThan(s[a], s[pivot]); a++ { // data[a] < pivot
			}
			if a >= b {
				break
			}
			// data[a] == pivot; data[b-1] < pivot
			s.Swap(a, b-1)
			a++
			b--
		}
	}
	// Swap pivot into middle
	s.Swap(pivot, b-1)
	return b - 1, c
}

func (s Slice[T]) quickSort(a, b, maxDepth int, lessThan genfuncs.LessThan[T]) {
	for b-a > 12 { // Use ShellSort for slices <= 12 elements
		if maxDepth == 0 {
			s.heapSort(a, b, lessThan)
			return
		}
		maxDepth--
		mlo, mhi := s.doPivot(a, b, lessThan)
		// Avoiding recursion on the larger subproblem guarantees
		// a stack depth of at most lg(b-a).
		if mlo-a < b-mhi {
			s.quickSort(a, mlo, maxDepth, lessThan)
			a = mhi // i.e., quickSort(data, mhi, b)
		} else {
			s.quickSort(mhi, b, maxDepth, lessThan)
			b = mlo // i.e., quickSort(data, a, mlo)
		}
	}
	if b-a > 1 {
		// Do ShellSort pass with gap 6
		// It could be written in this simplified form cause b-a <= 12
		for i := a + 6; i < b; i++ {
			if lessThan(s[i], s[i-6]) {
				s.Swap(i, i-6)
			}
		}
		s.insertionSort(a, b, lessThan)
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
