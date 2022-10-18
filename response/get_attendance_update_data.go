package response

type GetAttendanceUpdateData struct {
	Response

	AtCheckInfoForOpenVo `json:"result"`
}

type AtCheckInfoForOpenVo struct {
	WorkDate             string                        `json:"work_date"`
	AttendanceResultList []AtAttendanceResultForOpenVo `json:"attendance_result_list"`
	UserId               string                        `json:"userid"`
	ApproveList          []AtApproveForOpenVo          `json:"approve_list"`
	CheckRecordList      []AtAttendanceRecordForOpenVo `json:"check_record_list"`
	CorpId               string                        `json:"corpId"`
	ClassSettingInfo     AtClassSettingInfoForOpenVo   `json:"class_setting_info"`
}

type AtAttendanceResultForOpenVo struct {
	RecordId       int    `json:"record_id"`
	SourceType     string `json:"source_type"`
	PlanCheckTime  string `json:"plan_check_time"`
	ClassId        int    `json:"class_id"`
	LocationMethod string `json:"location_method"`
	LocationResult string `json:"location_result"`
	OutsideRemark  string `json:"outside_remark"`
	PlanId         int    `json:"plan_id"`
	UserAddress    string `json:"user_address"`
	GroupId        int    `json:"group_id"`
	UserCheckTime  string `json:"user_check_time"`
	ProcInstId     string `json:"procInst_id"`
	CheckType      string `json:"check_type"`
	TimeResult     string `json:"time_result"`
}

type AtApproveForOpenVo struct {
	DurationUnit string `json:"duration_unit"`
	Duration     string `json:"duration"`
	SubType      string `json:"sub_type"`
	TagName      string `json:"tag_name"`
	ProcInstId   string `json:"procInst_id"`
	BeginTime    string `json:"begin_time"`
	BizType      int    `json:"biz_type"`
	EndTime      string `json:"end_time"`
	GmtFinished  string `json:"gmt_finished"`
}

type AtAttendanceRecordForOpenVo struct {
	RecordId          int    `json:"record_id"`
	SourceType        string `json:"source_type"`
	UserAccuracy      string `json:"user_accuracy"`
	ValidMatched      bool   `json:"valid_matched"`
	UserCheckTime     string `json:"user_check_time"`
	UserLongitude     string `json:"user_longitude"`
	UserSsid          string `json:"user_ssid"`
	BaseAccuracy      string `json:"base_accuracy"`
	UserMacAddr       string `json:"user_mac_addr"`
	UserLatitude      string `json:"user_latitude"`
	BaseAddress       string `json:"base_address"`
	InvalidRecordMsg  string `json:"invalid_record_msg"`
	InvalidRecordType string `json:"invalid_record_type"`
}

type AtClassSettingInfoForOpenVo struct {
	RestTimeVoList []AtRestTimeVo `json:"rest_time_vo_list"`
}
type AtRestTimeVo struct {
	RestBeginTime int `json:"rest_begin_time"`
	RestEndTime   int `json:"rest_end_time"`
}
