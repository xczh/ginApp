package main

import (
	"app"
	"fmt"
	"os"

	"github.com/spf13/pflag"
)

const AppName = "app"
const Copyright = "Copyright (C) 2020 xczh. All Rights Reserved."

var BuildVersion = "0.0.0"
var BuildGitHash = "(Not provided)"
var BuildTime = "(Not provided)"
var BuildGoVersion = "(Not provided)"

func main() {
	parseCommandLineFlag()
	app.Run(pflag.CommandLine)
}

func parseCommandLineFlag() {
	pflag.Usage = func() {
		_, _ = fmt.Fprintf(os.Stderr,
			`%s server %s
  Git Hash: %s
  Build Time: %s
  Build Go Version: %s
%s
		
Usage of %s:
`,
			AppName, BuildVersion, BuildGitHash, BuildTime, BuildGoVersion, Copyright, os.Args[0])
		pflag.CommandLine.PrintDefaults()
	}

	pflag.ErrHelp = fmt.Errorf("Print help")

	pflag.StringP("config", "c", "", "path to config file")
	pflag.BoolP("test", "t", false, "test config file only")

	pflag.Parse()
}
