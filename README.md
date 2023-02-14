# helm-values-gen

`helm-values-gen` generates a YAML payload that contains all default values in a given JSON schema.
This is useful when maintaining a helm project with a `values.schema.json` that already contains the default values, which are expected in `values.yaml`.

## Installation

Currently, `helm-values-gen` can be installed as go binary.
In the future it will be (additionally) possible to install it as a helm plugin.

```nohighlight
go install github.com/giantswarm/helm-values-gen@latest
```

## Usage

Executing `schemalint verify` without any options will check whether a file is valid JSON Schema and whether it is normalized.

```nohighlight
$ helm-values-gen values.schema.json -o values.yaml

Wrote default values to values.yaml
```

Use `--help` to learn about more options.
