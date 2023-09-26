// Copyright 2023 Ivan Guerreschi. All rights reserved.
//
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"github.com/ivanGuerreschi/whatversion/internal/utility"
	"github.com/ivanGuerreschi/whatversion/internal/version"
)

func main() {
	version.GetVersion(utility.GetHome() + "/.apps.csv")
}
