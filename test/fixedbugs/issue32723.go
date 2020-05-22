// errorcheck

// Copyright 2019 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Complex literal comparison

package p

const x = 1i
const y = 1i < 2i // ERROR "invalid operation: .*not defined on untyped complex"
const z = x < 2i  // ERROR "invalid operation: .*not defined on untyped complex"

func f() {
	_ = 1i < 2i // ERROR "invalid operation: .*not defined on untyped complex"
	_ = 1i < 2  // ERROR "invalid operation: .*not defined on untyped complex"
	_ = 1 < 2i  // ERROR "invalid operation: .*not defined on untyped complex"

	c := 1i
	_ = c < 2i // ERROR "invalid operation: .*not defined on complex128"
}
