package main

type Record struct {
	ID string
	A  float64
	B  int
}

type Backend struct {
	Tables map[string]*Table
}

func NewBackend() *Backend {
	return &Backend{
		Tables: map[string]*Table{},
	}
}

type ColumnType uint

const (
	StringType ColumnType = iota
	IntegerType
	FloatType
)

type Column struct {
	name  string
	cType ColumnType
}

type Table struct {
	name    string
	columns []Column
	rows    [][]MemoryObject
}

type MemoryObject []byte
