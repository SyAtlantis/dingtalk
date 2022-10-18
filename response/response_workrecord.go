package response

type SaveApprovalTemplate struct {
	Response

	ProcessTopVo `json:"result"`
}
type ProcessTopVo struct {
	ProcessCode string `json:"process_code"`
}

type GetApprovalTemplate struct {
	Response

	ProcessCode string `json:"process_code"`
}

type DeleteApprovalTemplate struct {
	Response
}

type InitiateApprovalProcess struct {
	Response

	SaveFaceProcessInstanceResponse `json:"result"`
}
type SaveFaceProcessInstanceResponse struct {
	ProcessInstanceId string `json:"process_instance_id"`
}

type UpdateApprovalProcess struct {
	Response
}

type BatchUpdateApprovalProcess struct {
	Response
}

type CreateWorkrecord struct {
	Response

	Tasks []TaskTopVo `json:"tasks"`
}
type TaskTopVo struct {
	TaskId int    `json:"task_id"`
	UserId string `json:"userid"`
}

type QueryWorkrecordList struct {
	Response

	PageResult `json:"result"`
}

type PageResult struct {
	HasMore bool           `json:"has_more"`
	List    []WorkRecordVo `json:"list"`
}
type WorkRecordVo struct {
	Url        string       `json:"url"`
	TaskId     string       `json:"task_id"`
	InstanceId string       `json:"instance_id"`
	Title      string       `json:"title"`
	Forms      []FormItemVo `json:"forms"`
}
type FormItemVo struct {
	Title string `json:"title"`
}

type UpdateWorkrecord struct {
	Response
}

type CancelWorkrecord struct {
	Response
}
