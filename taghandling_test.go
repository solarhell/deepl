package deepl_test

import (
	"testing"

	"github.com/solarhell/deepl"
)

func TestTagHandlingStrategy_Value_String(t *testing.T) {
	tests := map[deepl.TagHandlingStrategy]string{
		deepl.DefaultTagHandling: "",
		deepl.XMLTagHandling:     "xml",
		deepl.HTMLTagHandling:    "html",
	}

	for strategy, want := range tests {
		if strategy.Value() != want {
			t.Errorf("expected strategy value %q; got %q", want, strategy.Value())
		}

		want = string(strategy)
		if strategy.String() != want {
			t.Errorf("expected strategy string value %q; got %q", want, strategy.String())
		}
	}
}
