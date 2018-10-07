package models

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestModel_GenerateId(t *testing.T) {
	assert := assert.New(t)

	id := GenerateId()
	assert.Len(id, 32)

	id2 := GenerateId()
	assert.NotEqual(id, id2)
}
