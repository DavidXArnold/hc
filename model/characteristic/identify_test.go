package characteristic

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIdentify(t *testing.T) {
	i := NewIdentify(true)
	assert.Equal(t, i.Type, CharTypeIdentify)
	assert.True(t, i.Identify())
	i.SetIdentify(false)
	assert.False(t, i.Identify())
}
