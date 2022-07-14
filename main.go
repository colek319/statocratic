package main

import (
	"flag"
	"fmt"
	"github.com/vaguelytitled/statocratic/pkg/democracy"
	"github.com/vaguelytitled/statocratic/version"
)

func main() {
	versionFlag := flag.Bool("version", false, "Version")
	flag.Parse()

	if *versionFlag {
		fmt.Println("Build Date:", version.BuildDate)
		fmt.Println("Git Commit:", version.GitCommit)
		fmt.Println("Version:", version.Version)
		fmt.Println("Go Version:", version.GoVersion)
		fmt.Println("OS / Arch:", version.OsArch)
		return
	}
	fmt.Println("Hello.")
	dem := democracy.NewSimulation(1000000, 50, 100000, )
}

// vote:
// protest:
// policy should align with median sentiment

func electPolicy(d democracy.Democracy) democracy.Democracy {
	for _, s := range d.States {
		for _, s.PolicyPreference
	}
}
