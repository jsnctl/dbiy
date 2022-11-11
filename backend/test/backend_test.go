package test

import (
	"github.com/jsnctl/dbiy/backend"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreateTable(t *testing.T) {
	columns := []backend.Column{
		{"A", backend.IntegerType},
		{"B", backend.FloatType},
	}

	table := backend.CreateTable("test", columns)

	assert.Equal(t, 2, len(table.Columns))
}
