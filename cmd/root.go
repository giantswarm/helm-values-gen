package cmd

import (
	"os"

	"github.com/giantswarm/helm-values-gen/v2/cmd/root"
)

func Execute() {
	rootCmd := root.New()
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
