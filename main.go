package main

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
	colIdx := make(map[string]int, len(cols))
	for i, v := range cols {
		colIdx[v.Name] = i
	}
	return Table{
		Name:    name,
		columns: cols,
		colIdx:  colIdx,
	}
}
