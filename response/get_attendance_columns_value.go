package response

type GetAttendanceColumnsValue struct {
	Response

	Columns struct {
		Id           int    `json:"id"`
		Type         int    `json:"type"`
		Name         string `json:"name"`
		Alias        string `json:"alias"`
		Status       int    `json:"status"`
		SubType      int    `json:"sub_type"`
		ExpressionId int    `json:"expression_id"`
	} `json:"columns"`
}
