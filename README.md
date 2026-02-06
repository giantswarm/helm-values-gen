# helm-values-gen

`helm-values-gen` generates a YAML payload that contains all default values in
a given JSON schema. This is useful when maintaining a helm project with a
`values.schema.json` that already contains the default values that are
expected in `values.yaml`.

## Installation

Currently, `helm-values-gen` can be installed as Go binary.
In the future it will be (additionally) possible to install it as a helm plugin.

```nohighlight
go install github.com/giantswarm/helm-values-gen@latest
```

## Usage

```nohighlight
$ helm-values-gen values.schema.json -o values.yaml

Wrote default values to values.yaml
```

Use `--help` to learn about more options.

## Details

As JSON schemas can be nested deeply, we need to define which default values
will be contained in the resulting payload.
The tool follows these steps:

1. Start in the root of the schema
2. If there is a default value, add it to the payload.
3. Recurse into all properties.

**Note**:
The tool does not recurse into the `items` keyword of an array.
To specify default values for an array you have to set a default on the level
of the array.

Consider the following example:

values.schema.json

```json
{
  "type": "object",
  "properties": {
    "addresses": {
      "type": "array",
      "items": {
        "type": "object",
        "properties": {
          "city": {
            "type": "string",
            "default": "New York"
          }
        }
      },
    },
    "lastName": {
      "type": "string",
      "default": "Doe"
    }
  }
}
```

values.yaml

```yaml
lastName: Doe
```

The default value `addresses[].city = "New York"` is not reflected in
`values.yaml`.

To specify a default value for an array consider the following example:

```json
{
  "type": "object",
  "properties": {
    "addresses": {
      "type": "array",
      "items": {
        "type": "object"
        "properties": {
          "city": {
            "type": "string",
            "default": "New York"
          }
        }
      },
      "default": [
        {
          "city": "New York"
        }
      ]
    },
    "lastName": {
      "type": "string",
      "default": "Doe"
    }
  }
}
```

values.yaml

```yaml
addresses:
  - city: New York
lastName: Doe
```

## GitHub Action

An action to run `schemalint verify` on the `values.schema.json` in app repositories in provided in `actions/verify-helm-schema`.

**Example workflow**:

```yaml
name: JSON schema validation
on:
  push: {}

jobs:
  generate:
    name: Check that values.yaml is generated from values.schema.json with helm-values-gen
    runs-on: ubuntu-latest
    steps:
      - name: Checkout code
        uses: actions/checkout@v3

      - name: Run helm-values-gen
        id: run-helm-values-gen
        uses: giantswarm/helm-values-gen/actions/ensure-generated@v1
```
