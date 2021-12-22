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

package genfuncs_test

import (
	"fmt"
	"github.com/nwillc/genfuncs"
	"testing"
	"time"
)

var (
	lexicalOrder   = genfuncs.OrderedComparator[string]()
	reverseLexical = genfuncs.ReverseComparator(lexicalOrder)
)

func TestFunctionExamples(t *testing.T) {
	ExampleOrderedComparator()
	ExampleReverseComparator()
	ExampleStringerStringer()
}

func ExampleOrderedComparator() {
	fmt.Println(lexicalOrder("a", "b")) // -1
	fmt.Println(lexicalOrder("a", "a")) // 0
	fmt.Println(lexicalOrder("b", "a")) // 1
}

func ExampleReverseComparator() {
	fmt.Println(lexicalOrder("a", "b"))   // -1
	fmt.Println(reverseLexical("a", "b")) // 1
}

func ExampleStringerStringer() {
	var epoch time.Time
	fmt.Println(epoch.String()) // 0001-01-01 00:00:00 +0000 UTC
	stringer := genfuncs.StringerStringer[time.Time]()
	fmt.Println(stringer(epoch)) // 0001-01-01 00:00:00 +0000 UTC
}
