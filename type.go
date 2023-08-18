package main

type Database struct {
	tables map[string]Table
}

type Table struct {
	Name    string
	columns []Column
	rows    [][]any
	colIdx  map[string]int
}

type Column struct {
	Name string
	Type int
}

type Expression struct {
	Column string
}
