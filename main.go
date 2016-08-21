package main

import (
	"fmt"
	"os"
	"runtime"

	"gopkg.in/alecthomas/kingpin.v2"
)

var (
	Name         string = "blacklist-checker"
	BuildVersion string
	BuildHash    string
	BuildDate    string
)

var (
	verbose     = kingpin.Flag("verbose", "Verbose mode.").Bool()
	version     = kingpin.Flag("version", "Show version and terminate").Short('v').Bool()
	baseversion = kingpin.Flag("baseversion", "Show base version").Short('b').Bool()
	nameserver  = kingpin.Flag("nameserver", "Name server to use").Short('n').Default("8.8.8.8").String()
)

func init() {
	kingpin.CommandLine.HelpFlag.Short('h')
	kingpin.Parse()

	if *version {
		fmt.Printf("%s version %s build %s (%s), build on %s\n", Name, BuildVersion, BuildHash, runtime.GOARCH, BuildDate)
		os.Exit(0)
	}

	if *baseversion {
		fmt.Printf("%s#%s (%s)\n", BuildVersion, BuildHash, runtime.GOARCH)
		os.Exit(0)
	}
}

func main() {
}
