package response

type GetAttendanceColumnsValue struct {
	Response

	ColumnsValue `json:"result"`
}

type ColumnsValue struct {
	ColumnVals []ColumnValForTopVo `json:"column_vals"`
}

type ColumnValForTopVo struct {
	ColumnVals []ColumnDayAndVal `json:"column_vals"`
	ColumnVo   ColumnForTopVo    `json:"column_vo"`
	FixedValue string            `json:"fixed_value"`
}

type ColumnDayAndVal struct {
	Id int `json:"id"`
}

type ColumnForTopVo struct {
}
