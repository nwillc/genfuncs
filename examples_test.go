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
	ExampleStringerToString()
	ExampleTransformArgs()
}

func ExampleMin() {
	fmt.Println(genfuncs.Min(1, 2))
	words := []string{"dog", "cat", "gorilla"}
	fmt.Println(genfuncs.Min(words...))
	// Output:
	// 1
	// cat
}

func ExampleMax() {
	fmt.Println(genfuncs.Max(1, 2))
	words := []string{"dog", "cat", "gorilla"}
	fmt.Println(genfuncs.Max(words...))
	// Output:
	// 2
	// gorilla
}

func ExampleStringerToString() {
	var epoch time.Time
	fmt.Println(epoch.String())
	stringer := genfuncs.StringerToString[time.Time]()
	fmt.Println(stringer(epoch))
	// Output:
	// 0001-01-01 00:00:00 +0000 UTC
	// 0001-01-01 00:00:00 +0000 UTC
}

func ExampleTransformArgs() {
	var unixTime = func(t time.Time) int64 { return t.Unix() }
	var chronoOrder = genfuncs.TransformArgs(unixTime, genfuncs.LessOrdered[int64])
	now := time.Now()
	fmt.Println(chronoOrder(now, now.Add(time.Second)))
	// Output: true
}
