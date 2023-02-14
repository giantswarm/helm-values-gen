package root

import (
	"github.com/spf13/cobra"
)

func (r *runner) preRun(cmd *cobra.Command, args []string) error {
	error := r.flag.validate()
	if error != nil {
		return error
	}
	return nil
}
