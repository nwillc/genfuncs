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
	"os"
	"testing"
	"time"

	"github.com/nwillc/genfuncs"
)

func TestFunctionExamples(t *testing.T) {
	if _, ok := os.LookupEnv("RUN_EXAMPLES"); !ok {
		t.Skip("skipping: RUN_EXAMPLES not set")
	}
	// Functions
	ExampleMax()
	ExampleMin()
	// ExampleSLexicalOrder()
	// ExampleSReverseLexicalOrder()
	ExampleStringerToString()
	ExampleTransformArgs()
}

func ExampleMin() {
	fmt.Println(genfuncs.Min(1, 2)) // 1
	words := []string{"dog", "cat", "gorilla"}
	fmt.Println(genfuncs.Min(words...)) // cat
}

func ExampleMax() {
	fmt.Println(genfuncs.Max(1, 2)) // 2
	words := []string{"dog", "cat", "gorilla"}
	fmt.Println(genfuncs.Max(words...)) // gorilla
}

// func ExampleSLexicalOrder() {
// 	fmt.Println(genfuncs.LessOrdered("a", "b")) // true
// 	fmt.Println(genfuncs.LessOrdered("a", "a")) // false
// 	fmt.Println(genfuncs.LessOrdered("b", "a")) // false
// }
//
// func ExampleSReverseLexicalOrder() {
// 	fmt.Println(genfuncs.SLexicalOrdered("a", "b"))        // true
// 	fmt.Println(genfuncs.SReverseLexicalOrdered("a", "b")) // false
// }

func ExampleStringerToString() {
	var epoch time.Time
	fmt.Println(epoch.String()) // 0001-01-01 00:00:00 +0000 UTC
	stringer := genfuncs.StringerToString[time.Time]()
	fmt.Println(stringer(epoch)) // 0001-01-01 00:00:00 +0000 UTC
}

func ExampleTransformArgs() {
	var unixTime = func(t time.Time) int64 { return t.Unix() }
	var chronoOrder = genfuncs.TransformArgs(unixTime, genfuncs.LessOrdered[int64])
	now := time.Now()
	fmt.Println(chronoOrder(now, now.Add(time.Second))) // true
}
