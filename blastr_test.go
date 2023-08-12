package blastr

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const testNsec = "nsec1kk96g7rj34yfnyyjmgcjwa7kvjvddngxy2lwfx357w4lj8fnc20qpl0cr6"

func TestNew(t *testing.T) {
	// Valid nsec
	{
		b, err := New(testNsec)
		assert.NoError(t, err)
		assert.NotNil(t, b)
	}

	// Invalid nsec
	{
		b, err := New("nsec1nope")
		assert.NotNil(t, err)
		assert.Nil(t, b)
	}
}

func TestNewOptions(t *testing.T) {
	// Default options
	{
		b, err := New(testNsec)
		assert.NoError(t, err)
		assert.NotNil(t, b)
		assert.Equal(t, defaultOptions(), b.opts)
	}

	// Explicit options
	{
		opts := []Option{
			WithStrictErrors(),
			WithCustomRelays([]string{"wss://example.com"}),
		}

		b, err := New(testNsec, opts...)
		assert.NoError(t, err)
		assert.NotNil(t, b)

		assert.True(t, b.opts.strictErrors)
		assert.Equal(t, []string{"wss://example.com"}, b.opts.relayURLs)
	}
}
