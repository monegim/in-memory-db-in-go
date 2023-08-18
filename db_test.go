package main

import (
	"github.com/google/go-cmp/cmp"
	"testing"
)

const (
	ColumnInt = 1 << iota
	ColumnVarchar
)

func TestDatabase(t *testing.T) {
	db := NewDatabase()
	col1 := Column{
		Name: "id",
		Type: ColumnInt,
	}
	col2 := Column{"name", ColumnVarchar}
	col3 := Column{"population", ColumnInt}
	tab := db.CreateTable("cities", []Column{
		col1,
		col2,
		col3,
	})
	if diff := cmp.Diff(len(db.tables), 1); diff != "" {
		t.Errorf("Select: (-want +got)\n%s", diff)
	}
	tab.Insert(1, "Tokyo", 37)
	tab.Insert(2, "Delhi", 28)
	tab.Insert(3, "Shanghai", 28)
}
