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

package sequences_test

import (
	"fmt"
	"github.com/nwillc/genfuncs"
	"github.com/nwillc/genfuncs/container"
	"github.com/nwillc/genfuncs/container/maps"
	"github.com/nwillc/genfuncs/container/sequences"
	"github.com/nwillc/genfuncs/internal/tests"
	"strings"
	"testing"
)

func TestFunctionExamples(t *testing.T) {
	tests.MaybeRunExamples(t)
	ExampleAssociate()
	ExampleAssociateWith()
	ExampleFold()
}

func ExampleAssociate() {
	byLastName := func(n string) *genfuncs.Result[*maps.Entry[string, string]] {
		parts := strings.Split(n, " ")
		return genfuncs.NewResult(maps.NewEntry(parts[1], n))
	}
	names := sequences.NewSequence[string]("fred flintstone", "barney rubble")
	sequences.Associate(names, byLastName).
		OnSuccess(func(nameMap container.GMap[string, string]) {
			fmt.Println(nameMap["rubble"])
		})

	// Output: barney rubble
}

func ExampleAssociateWith() {
	oddEven := func(i int) *genfuncs.Result[string] {
		if i%2 == 0 {
			return genfuncs.NewResult("EVEN")
		}
		return genfuncs.NewResult("ODD")
	}
	numbers := sequences.NewSequence[int](1, 2, 3, 4)
	sequences.AssociateWith[int, string](numbers, oddEven).OnSuccess(func(odsEvensMap container.GMap[int, string]) {
		fmt.Println(odsEvensMap[2])
		fmt.Println(odsEvensMap[3])
	})
	// Output:
	// EVEN
	// ODD
}

func ExampleFold() {
	sequence := sequences.NewSequence(1, 2, 3, 4)
	sum := sequences.Fold(sequence, 0, func(prior, value int) int {
		return prior + value
	})
	fmt.Println(sum)
	// Output: 10
}
