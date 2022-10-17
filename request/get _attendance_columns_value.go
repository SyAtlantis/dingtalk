package request

type GetAttendanceColumnsValue struct {
	UserId       string `json:"userid"`
	ColumnIdList string `json:"column_id_list"`
	FromDate     string `json:"from_date"`
	ToDate       string `json:"to_date"`
}

func NewGetAttendanceColumnsValue() *GetAttendanceColumnsValue {
	return &GetAttendanceColumnsValue{}
}
