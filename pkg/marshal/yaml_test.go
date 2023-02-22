package marshal

import (
	"strings"
	"testing"
)

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
