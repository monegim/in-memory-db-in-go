package main

func (t *Table) Insert(data ...any) {
	t.rows = append(t.rows, data)
}
