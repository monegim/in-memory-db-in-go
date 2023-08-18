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
