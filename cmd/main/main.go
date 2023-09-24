package main

import (
	"github.com/ivanGuerreschi/whatversion/internal/utility"
	"github.com/ivanGuerreschi/whatversion/internal/version"
)

func main() {
	version.GetVersion(utility.GetHome() + "/.apps.csv")
}
