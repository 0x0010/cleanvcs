package utils

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestCurrentPath(t *testing.T) {
	cp := CurrentPath();
	assert.NotNil(t, cp)
}
