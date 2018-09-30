package models

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestModel_GenerateId(t *testing.T) {
	assert := assert.New(t)

	m := &Model{}
	id := m.GenerateId()

	assert.Equal(32, len(id))

	id2 := m.GenerateId()

	assert.NotEqual(id, id2)
}
