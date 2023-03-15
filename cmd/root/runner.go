package root

import (
	"github.com/giantswarm/helm-values-gen/v2/pkg/convert"
	"github.com/giantswarm/helm-values-gen/v2/pkg/jsonschema"
	"github.com/giantswarm/helm-values-gen/v2/pkg/marshal"
	"github.com/giantswarm/microerror"
	"github.com/spf13/cobra"
)

type runner struct {
	flag *flag
}

func (r *runner) run(cmd *cobra.Command, args []string) error {
	path := args[0]

	schema, err := jsonschema.Compile(path)
	if err != nil {
		return microerror.Mask(err)
	}

	var defaultValues interface{}
	defaultValues, err = convert.BuildDefaultValuesFor(schema)
	if err != nil {
		return microerror.Mask(err)
	}

	marshalled, err := marshal.ToYaml(defaultValues)
	if err != nil {
		return microerror.Mask(err)
	}

	err = handleOutput(r.flag, marshalled)
	if err != nil {
		return microerror.Mask(err)
	}

	return nil
}
