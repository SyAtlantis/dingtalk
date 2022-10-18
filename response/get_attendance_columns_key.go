package response

type GetAttendanceColumnsKey struct {
	Response

	ColumnsKey `json:"result"`
}

type ColumnsKey struct {
	Columns []Columns `json:"columns"`
}

type Columns struct {
	Id           int    `json:"id"`
	Type         int    `json:"type"`
	Name         string `json:"name"`
	Alias        string `json:"alias"`
	Status       int    `json:"status"`
	SubType      int    `json:"sub_type"`
	ExpressionId int    `json:"expression_id"`
}
