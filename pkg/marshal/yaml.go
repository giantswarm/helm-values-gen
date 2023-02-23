package marshal

import (
	"bytes"

	"github.com/giantswarm/microerror"
	"gopkg.in/yaml.v3"
	sigsYaml "sigs.k8s.io/yaml"
)

func ToYaml(data interface{}) ([]byte, error) {
	convertedData, err := convertJSONNumbers(data)
	if err != nil {
		return nil, microerror.Maskf(marshalError, "Error during YAML marshaling:\n%s", err)
	}

	var buf bytes.Buffer
	enc := yaml.NewEncoder(&buf)
	enc.SetIndent(2)
	err = enc.Encode(convertedData)
	if err != nil {
		return nil, microerror.Maskf(marshalError, "Error marshalling to YAML:\n%s", err)
	}
	marshalled := buf.Bytes()
	return marshalled, nil
}

// marshals to YAML and unmarshals back to interface{} to convert json.Number appropriately
func convertJSONNumbers(data interface{}) (interface{}, error) {
	b, err := sigsYaml.Marshal(data)
	if err != nil {
		return nil, microerror.Mask(err)
	}
	var converted interface{}
	err = sigsYaml.Unmarshal(b, &converted)
	return converted, microerror.Mask(err)

}
