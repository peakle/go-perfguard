package main

import (
	"flag"
)

func addCommonFlags(r *runner, fs *flag.FlagSet) {
	fs.BoolVar(&r.autofix, "fix", false,
		`apply the suggested fixes automatically, where possible`)
	fs.StringVar(&r.goVersion, "go", "",
		`select the Go version to target; leave as empty string for the latest`)
	fs.BoolVar(&r.absFilenames, "abs", false,
		`print absolute filenames in the output`)
}