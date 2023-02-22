package marshal

import (
	"github.com/giantswarm/microerror"
	"gopkg.in/yaml.v3"
)

func ToYaml(data interface{}) ([]byte, error) {
	marshalled, err := yaml.Marshal(data)
	if err != nil {
		return nil, microerror.Maskf(marshalError, "Error marshalling to YAML:\n%s", err)
	}
	return marshalled, nil
}
