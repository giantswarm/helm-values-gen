package root

import (
	"github.com/spf13/cobra"

	"github.com/giantswarm/helm-values-gen/pkg/project"
)

func New() *cobra.Command {
	flag := &flag{}
	runner := &runner{flag: flag}

	var rootCmd = &cobra.Command{
		Use:   "helm-values-gen PATH",
		Short: "Generate a payload from a JSON schema file that contains all the default values",
		Long: `Generate a payload from a JSON schema file that contains all the default values.

To generate the payload the tool will traverse into all properties. If some
property has a default, then all its parents will be added to the payload.
An exception to this rule are arrays. If an array has a default, then that
value will be used, if not then this tool won't traverse into the array's items.
`,
		Example: `  helm-values-gen schema.json
  helm-values-gen schema.json -o payload.json`,
		Args:       cobra.ExactArgs(1),
		ArgAliases: []string{"PATH"},
		RunE:       runner.run,
		PreRunE:    runner.preRun,
		Version:    project.Version(),
	}
	flag.init(rootCmd)

	return rootCmd

}
