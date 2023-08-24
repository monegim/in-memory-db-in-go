package main

const (
	ConditionEq = "="
	ConditionNe = "!="
)

func (t *Table) Insert(data ...any) {
	t.rows = append(t.rows, data)
}

func (t *Table) Select(cols []Expression) {
	//res := [][]any{}
	//for i := 0; i < len(t.rows); i++ {
	//	if t.r
	//}
}

func (t *Table) rowMatch(cond []Condition, i int) bool {
	if t.rows[i] == nil {
		return false
	}
	if cond == nil {
		return true
	}
	for _, c := range cond {
		j := t.getColumnIndex(c.Column)
		if !c.Eval(t.rows[i][j]) {
			return false
		}
	}
	return true
}

func (t *Table) getColumnIndex(name string) int {
	return t.colIdx[name]
}

func (c Condition) Eval(val any) bool {
	if c.Eq == ConditionEq {
		return c.Value == val
	} else if c.Eq == ConditionNe {
		return c.Value != val
	}
	return false
}
