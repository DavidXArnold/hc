package common

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestSerialForName(t *testing.T) {
	storage, err := NewFileStorage(os.TempDir())
	assert.Nil(t, err)
	name := "My Accessory"
	serial := GetSerialNumberForAccessoryName(name, storage)
	serial2 := GetSerialNumberForAccessoryName(name, storage)
	assert.Equal(t, serial, serial2)
}
