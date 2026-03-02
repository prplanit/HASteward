package main

import (
	"fmt"
	"os"

	"gitlab.prplanit.com/precisionplanit/hasteward/cmd"
	"gitlab.prplanit.com/precisionplanit/hasteward/common"

	// Register engines via init()
	_ "gitlab.prplanit.com/precisionplanit/hasteward/engine/cnpg"
	_ "gitlab.prplanit.com/precisionplanit/hasteward/engine/galera"
)

func main() {
	common.InitLogging(false)
	if err := cmd.RootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}
