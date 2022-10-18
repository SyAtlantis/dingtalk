package request

type GetAttendanceUpdateData struct {
	UserId   string `json:"userid"`
	WorkDate string `json:"work_date"`
}

func NewGetAttendanceUpdateData() *GetAttendanceUpdateData {
	return &GetAttendanceUpdateData{}
}
