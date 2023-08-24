# in-memory-db-in-go
In-memory Database in Go

```go
tab := db.CreateTable("cities", []Column{
		{"id", ColumnInt},
		{"name", ColumnVarchar},
		{"population", ColumnInt},
	})
```
