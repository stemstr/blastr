package blastr

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWithCustomRelays(t *testing.T) {
	var tests = []struct {
		name   string
		relays []string
	}{
		{"none", []string{}},
		{"single", []string{"wss://test.example.com"}},
		{"many", []string{"wss://test.example.com", "wss://foo.example"}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			opts := &Options{}
			WithCustomRelays(tt.relays)(opts)
			assert.Equal(t, tt.relays, opts.relayURLs)
		})
	}
}

func TestWithStrictErrors(t *testing.T) {
	opts := &Options{}
	WithStrictErrors()(opts)
	assert.True(t, opts.strictErrors)
}
