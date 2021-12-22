package commons

import (
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
)

type DummyServiceOptions struct {
	Foo    string `env:"FOO"`
	Bar    string `env:"BAR"`
	Length int    `env:"LENGTH"`
}

func TestInitOptions(t *testing.T) {
	opts := &DummyServiceOptions{}
	err := InitOptions(opts, "dummy-service")
	assert.Nil(t, err)
	assert.NotNil(t, opts.Foo)
	assert.NotNil(t, opts.Bar)
	assert.Equal(t, "foo", opts.Foo)
	assert.Equal(t, "bar", opts.Bar)
	assert.Equal(t, 12, opts.Length)
	log.Println(opts)
}