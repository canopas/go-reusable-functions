package file

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetFileBytesError(t *testing.T) {
	asserts := assert.New(t)
	temp, err := GetFileBytes("")
	asserts.Error(err)
	asserts.Empty(temp)
}

func TestGetFileBytesSuccess(t *testing.T) {
	asserts := assert.New(t)
	temp, err := GetFileBytes("test.txt")
	asserts.NoError(err)
	asserts.NotEmpty(temp)
}
