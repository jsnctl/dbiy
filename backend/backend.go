package backend

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
	Name  string
	CType ColumnType
}

type Table struct {
	Name    string
	Columns []Column
	Rows    [][]MemoryObject
}

func CreateTable(name string, columns []Column) *Table {
	t := Table{}
	t.Name = name
	t.Columns = columns
	return &t
}

type MemoryObject []byte
