// Copyright 2012 The Go Authors.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"fmt"
	"os"
	"strings"
)

/*
 * Helpers for building runtime.
 */

// mkzversion writes zversion.go:
//
//	package sys
//	const DefaultGoroot = <goroot>
//	const TheVersion = <version>
//	const Goexperiment = <goexperiment>
//	const StackGuardMultiplier = <multiplier value>
//
func mkzversion(dir, file string) {
	out := fmt.Sprintf(
		"// auto generated by go tool dist\n"+
			"\n"+
			"package sys\n"+
			"\n"+
			"const DefaultGoroot = `%s`\n"+
			"const TheVersion = `%s`\n"+
			"const Goexperiment = `%s`\n"+
			"const StackGuardMultiplier = %d\n\n", goroot_final, findgoversion(), os.Getenv("GOEXPERIMENT"), stackGuardMultiplier())

	writefile(out, file, writeSkipSame)
}

// mkzbootstrap writes cmd/internal/obj/zbootstrap.go:
//
//	package obj
//
//	const defaultGOROOT = <goroot>
//	const defaultGO386 = <go386>
//	const defaultGOARM = <goarm>
//	const defaultGOOS = runtime.GOOS
//	const defaultGOARCH = runtime.GOARCH
//	const defaultGO_EXTLINK_ENABLED = <goextlinkenabled>
//	const version = <version>
//	const stackGuardMultiplier = <multiplier value>
//	const goexperiment = <goexperiment>
//
// The use of runtime.GOOS and runtime.GOARCH makes sure that
// a cross-compiled compiler expects to compile for its own target
// system. That is, if on a Mac you do:
//
//	GOOS=linux GOARCH=ppc64 go build cmd/compile
//
// the resulting compiler will default to generating linux/ppc64 object files.
// This is more useful than having it default to generating objects for the
// original target (in this example, a Mac).
func mkzbootstrap(file string) {
	out := fmt.Sprintf(
		"// auto generated by go tool dist\n"+
			"\n"+
			"package obj\n"+
			"\n"+
			"import \"runtime\"\n"+
			"\n"+
			"const defaultGOROOT = `%s`\n"+
			"const defaultGO386 = `%s`\n"+
			"const defaultGOARM = `%s`\n"+
			"const defaultGOOS = runtime.GOOS\n"+
			"const defaultGOARCH = runtime.GOARCH\n"+
			"const defaultGO_EXTLINK_ENABLED = `%s`\n"+
			"const version = `%s`\n"+
			"const stackGuardMultiplier = %d\n"+
			"const goexperiment = `%s`\n",
		goroot_final, go386, goarm, goextlinkenabled, findgoversion(), stackGuardMultiplier(), os.Getenv("GOEXPERIMENT"))

	writefile(out, file, writeSkipSame)
}

// stackGuardMultiplier returns a multiplier to apply to the default
// stack guard size.  Larger multipliers are used for non-optimized
// builds that have larger stack frames.
func stackGuardMultiplier() int {
	for _, s := range strings.Split(os.Getenv("GO_GCFLAGS"), " ") {
		if s == "-N" {
			return 2
		}
	}
	return 1
}
