package convert

import "github.com/giantswarm/microerror"

var schemaConversionError = &microerror.Error{
	Kind: "schemaConversionError",
}

var executionFailedError = &microerror.Error{
	Kind: "executionFailedError",
}
