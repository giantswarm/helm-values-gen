package root

import (
	"github.com/giantswarm/microerror"
	"github.com/spf13/cobra"
)

type flag struct {
	output string
	force  bool
}

func (f *flag) init(cmd *cobra.Command) {
	cmd.Flags().StringVarP(&f.output, "output", "o", "", "output file")
	cmd.Flags().BoolVarP(&f.force, "force", "f", false, "force overwrite of output file")
}

func (f *flag) validate() error {

	if f.output == "" && f.force {
		return microerror.Maskf(flagsInvalidError, "output file must be specified when force is set")
	}

	return nil
}
