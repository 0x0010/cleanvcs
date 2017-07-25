package utils

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCurrentPath(t *testing.T) {
	cp := CurrentPath()
	assert.NotNil(t, cp)
}
