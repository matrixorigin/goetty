package buf

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAlloc(t *testing.T) {
	alloc := newNonReusableAllocator()
	buf, err := alloc.Alloc(10)
	assert.NoError(t, err)
	assert.Equal(t, 10, len(buf))
}
