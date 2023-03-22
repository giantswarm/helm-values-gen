package convert

import (
	"github.com/giantswarm/helm-values-gen/pkg/jsonschema"
	"github.com/giantswarm/microerror"
	"github.com/imdario/mergo"
)

// Creates a nested map that holds all default values specified by the schema.
func BuildDefaultValuesFor(schema *jsonschema.Schema) (interface{}, error) {
	var err error

	refDefaultValues := interface{}(nil)
	if schema.Ref != nil {
		refDefaultValues, err = BuildDefaultValuesFor(schema.GetRef())
		if err != nil {
			return nil, microerror.Mask(err)
		}
	}

	if schema.Default != nil {
		return schema.Default, nil
	}

	defaultValues, err := buildDefaultValuesForProperties(schema)
	if err != nil {
		return nil, microerror.Mask(err)
	}

	defaultValues, err = merge(refDefaultValues, defaultValues)
	if err != nil {
		return nil, microerror.Mask(err)
	}

	return defaultValues, nil
}

func buildDefaultValuesForProperties(schema *jsonschema.Schema) (interface{}, error) {
	defaultValues := map[string]interface{}{}

	if len(schema.Properties) == 0 {
		return nil, nil
	}

	for name, property := range schema.GetProperties() {
		defaultValue, err := BuildDefaultValuesFor(property)
		if err != nil {
			return nil, microerror.Mask(err)
		}
		if defaultValue != nil {
			defaultValues[name] = defaultValue
		}
	}

	return defaultValues, nil
}

// Merges two arbitrary values.
// a will take precedence over b.
func merge(b interface{}, a interface{}) (interface{}, error) {
	if a == nil {
		return b, nil
	}
	if b == nil {
		return a, nil
	}
	aMap, aIsMap := a.(map[string]interface{})
	bMap, bIsMap := b.(map[string]interface{})
	if !aIsMap || !bIsMap {
		return nil, microerror.Maskf(executionFailedError, "Merged values must be maps if they are not nil")
	}

	err := mergo.Merge(&aMap, bMap)
	if err != nil {
		return nil, microerror.Maskf(schemaConversionError, "Error merging default values with default values from referenced schema:\n%s", err)
	}

	return aMap, nil
}
