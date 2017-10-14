package math

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMath(t *testing.T) {
	assert.Equal(t, int64(3), AddInt(1, 2))
	assert.Equal(t, int64(-1), SubInt(2, 3))
	assert.Equal(t, int64(12), MultInt(3, 4))
	assert.Equal(t, int64(2), DivInt(4, 2))
	assert.Equal(t, float64(1.5), AddFloat(1, 0.5))
	assert.Equal(t, float64(-6), SubFloat(2.5, 8.5))
	assert.Equal(t, float64(12.0), MultFloat(3, 4))
	assert.Equal(t, float64(2.5), DivFloat(5, 2))
}
