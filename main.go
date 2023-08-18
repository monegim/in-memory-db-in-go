package main

type Database struct {
	tables map[string]Table
}

type Table struct {
	Name    string
	columns []Column
	rows    [][]any
}

type Column struct {
	Name string
	Type int
}

func NewDatabase() *Database {
	return &Database{
		tables: make(map[string]Table),
	}
}

func (d *Database) CreateTable(name string, cols []Column) *Table {
	t := NewTable(name, cols)
	d.tables[t.Name] = t
	return &t

}

func NewTable(name string, cols []Column) Table {
	return Table{
		Name:    name,
		columns: cols,
	}
}
