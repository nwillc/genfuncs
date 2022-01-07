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

package genfuncs

var (
	// Orderings

	F32NumericOrder        = OrderedLessThan[float32]()
	F32ReverseNumericOrder = Reverse(F32NumericOrder)
	INumericOrder          = OrderedLessThan[int]()
	IReverseNumericOrder   = Reverse(INumericOrder)
	I64NumericOrder        = OrderedLessThan[int64]()
	I64ReverseNumericOrder = Reverse(I64NumericOrder)
	SLexicalOrder          = OrderedLessThan[string]()
	SReverseLexicalOrder   = Reverse(SLexicalOrder)

	// Predicates

	IsBlank    = IsEqualComparable("")
	IsNotBlank = IsBlank.Not()

	F32IsZero = IsEqualComparable(float32(0.0))
	F64IsZero = IsEqualComparable(0.0)
	IIsZero   = IsEqualComparable(0)
)

func IsEqualComparable[C comparable](c C) Predicate[C] {
	return func(a C) bool { return a == c }
}

func AreEqualComparable[C comparable](a, b C) bool {
	return a == b
}
