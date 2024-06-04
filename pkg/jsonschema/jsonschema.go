package jsonschema

import (
	"net/url"
	"os"
	"path/filepath"
	"runtime"

	"github.com/giantswarm/microerror"
	jsonschemalib "github.com/santhosh-tekuri/jsonschema/v6"
)

type Schema struct {
	*jsonschemalib.Schema
}

func (s *Schema) GetProperties() map[string]*Schema {
	properties := map[string]*Schema{}
	for name, property := range s.Properties {
		properties[name] = &Schema{property}
	}
	return properties
}

func (s *Schema) GetRef() *Schema {
	if s.Ref == nil {
		return nil
	}
	return &Schema{s.Ref}
}

// This processes the given input without specifying the draft to use.
// If the schema provides a valid `$schema` property, the one given will
// be used. If not, the latest draft will be used.
// In case of success, a string will be returned, otherwise an error.
func Compile(path string) (*Schema, error) {
	url, err := ToFileURL(path)

	if !fileExists(path) {
		return nil, microerror.Maskf(compilationError, "File %q not found", path)
	}

	if err != nil {
		return nil, microerror.Maskf(compilationError, "Error converting path to URL:\n%s", err)
	}

	compiler := jsonschemalib.NewCompiler()
	compiler.ExtractAnnotations = true

	schema, err := compiler.Compile(url)
	if err != nil {
		return nil, microerror.Maskf(compilationError, "Error compiling schema:\n%s", err)
	}
	return &Schema{schema}, nil
}

func ToFileURL(path string) (string, error) {
	path, err := filepath.Abs(path)
	if err != nil {
		return "", microerror.Mask(err)
	}
	path = filepath.ToSlash(path)
	if runtime.GOOS == "windows" {
		path = "/" + path
	}
	u, err := url.Parse("file://" + path)
	if err != nil {
		return "", microerror.Mask(err)
	}

	return u.String(), nil
}

func fileExists(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}
