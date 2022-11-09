package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreateTable(t *testing.T) {
	columns := []Column{
		{"A", IntegerType},
		{"B", FloatType},
	}

	table := CreateTable("test", columns)

	assert.Equal(t, 2, len(table.columns))
}
