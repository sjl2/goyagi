package sentry

import (
    "testing"

    "github.com/stretchr/testify/assert"
    "github.com/sjl2/goyagi/pkg/config"
)

func TestNew(t *testing.T) {
    cfg := config.Config{
        Environment: "test",
    }
    sentry, err := New(cfg)

    assert.NoError(t, err)
    assert.NotNil(t, sentry)
}
