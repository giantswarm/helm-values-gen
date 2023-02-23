package marshal

import (
	"encoding/json"
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestToYaml(t *testing.T) {
	testCases := []struct {
		name string
		data interface{}
		want string
	}{
		{
			name: "case 0: simple",
			data: map[string]interface{}{
				"foo": "bar",
			},
			want: "foo: bar\n",
		},
		{
			name: "case 1: long string with possible line break",
			data: map[string]interface{}{
				"test": "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa bbb",
			},
			want: "test: aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa bbb\n",
		},
		{
			name: "case 2: integer",
			data: map[string]interface{}{
				"test": 1,
			},
			want: "test: 1\n",
		},
		{
			name: "case 3: float",
			data: map[string]interface{}{
				"test": 1.1,
			},
			want: "test: 1.1\n",
		},
		{
			name: "case 2: json.Number",
			data: map[string]interface{}{
				"test": json.Number("1"),
			},
			want: "test: 1\n",
		},
		{
			name: "case 5: nested object",
			data: map[string]interface{}{
				"test": map[string]interface{}{
					"foo": "bar",
				},
			},
			want: `test:
  foo: bar
`,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got, err := ToYaml(tc.data)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}

			if string(got) != tc.want {
				t.Fatalf("unexpected result: %s", cmp.Diff(string(got), tc.want))
			}
		})
	}

}

func TestToYamlWithLongStrings(t *testing.T) {
	data := map[string]interface{}{
		"test": "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa bbb",
	}
	marshalled, err := ToYaml(data)
	if err != nil {
		t.Fatalf("Error marshalling to YAML:\n%s", err)
	}
	linesSep := "\n"
	nLines := strings.Count(string(marshalled), linesSep)
	if nLines != 1 {
		t.Fatalf("Expected 1 line, got %d. yaml.Marshal() should not split long strings", nLines)
	}
}
