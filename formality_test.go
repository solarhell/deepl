package deepl_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/solarhell/deepl"
)

func TestFormality_Value_String(t *testing.T) {
	tests := map[deepl.Formal]string{
		deepl.DefaultFormal: "default",
		deepl.LessFormal:    "less",
		deepl.MoreFormal:    "more",
	}

	for f, v := range tests {
		assert.Equal(t, f.Value(), v)
		assert.Equal(t, f.String(), v)
	}
}
