package request

type SaveApprovalTemplate struct {
	SaveProcessRequest SaveProcessRequest `json:"saveProcessRequest"`
}
type SaveProcessRequest struct {
	AgentId           int               `json:"agentid"`
	ProcessCode       string            `json:"process_code,omitempty"`
	Name              string            `json:"name"`
	Description       string            `json:"description"`
	FormComponentList []FormComponentVo `json:"form_component_list"`
	FakeMode          bool              `json:"fake_mode,omitempty"`
}
type FormComponentVo struct {
	ComponentName string              `json:"component_name"`
	Props         FormComponentPropVo `json:"props,omitempty"`
}
type FormComponentPropVo struct {
	Id       string `json:"id"`
	Label    string `json:"label"`
	Required bool   `json:"required,omitempty"`
}

type GetApprovalTemplate struct {
	Name string `json:"name"`
}

type DeleteApprovalTemplate struct {
	Request DeleteProcessRequest `json:"request"`
}
type DeleteProcessRequest struct {
	AgentId          int    `json:"agentid"`
	ProcessCode      string `json:"process_code"`
	CleanRunningTask bool   `json:"clean_running_task,omitempty"`
}

type InitiateApprovalProcess struct {
	Request SaveFakeProcessInstanceRequest `json:"request"`
}
type SaveFakeProcessInstanceRequest struct {
	AgentId             int                    `json:"agentid"`
	ProcessCode         string                 `json:"process_code"`
	OriginatorUserId    string                 `json:"originator_user_id"`
	FormComponentValues []FormComponentValueVo `json:"form_component_values"`
	Url                 string                 `json:"url"`
	Title               string                 `json:"title,omitempty"`
}
type FormComponentValueVo struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

type UpdateApprovalProcess struct {
	Request UpdateProcessInstanceRequest `json:"request,omitempty"`
}
type UpdateProcessInstanceRequest struct {
	AgentId           int    `json:"agentid,omitempty"`
	ProcessInstanceId string `json:"process_instance_id"`
	Status            string `json:"status"`
	Result            string `json:"result"`
	CancelRunningTask bool   `json:"cancel_running_task,omitempty"`
}

type BatchUpdateApprovalProcess struct {
	Request BatchUpdateProcessInstanceRequest `json:"request"`
}
type BatchUpdateProcessInstanceRequest struct {
	Instances []UpdateProcessInstanceRequest2 `json:"instances"`
	AgentId   int                             `json:"agentid,omitempty"`
}
type UpdateProcessInstanceRequest2 struct {
	ProcessInstanceId string `json:"process_instance_id"`
	Status            string `json:"status"`
	Result            string `json:"result"`
}

type CreateWorkrecord struct {
	Request SaveTaskRequest `json:"request"`
}
type SaveTaskRequest struct {
	AgentId           int         `json:"agentid,omitempty"`
	ProcessInstanceId string      `json:"process_instance_id"`
	ActivityId        string      `json:"activity_id,omitempty"`
	Tasks             []TaskTopVo `json:"tasks"`
}
type TaskTopVo struct {
	UserId string `json:"userid"`
	Url    string `json:"url"`
}

type QueryWorkrecordList struct {
	UserId string `json:"userid"`
	Offset int    `json:"offset"`
	Count  int    `json:"count"`
	Status int    `json:"status"`
}

type UpdateWorkrecord struct {
	Request UpdateTaskRequest `json:"request"`
}
type UpdateTaskRequest struct {
	AgentId           string       `json:"agentid"`
	ProcessInstanceId string       `json:"process_instance_id"`
	Tasks             []TaskTopVo2 `json:"tasks"`
}
type TaskTopVo2 struct {
	TaskId string `json:"task_id"`
	Status string `json:"status"`
	Result string `json:"result"`
}

type CancelWorkrecord struct {
	Request UpdateTaskRequest2 `json:"request"`
}
type UpdateTaskRequest2 struct {
	AgentId           string   `json:"agentid"`
	ProcessInstanceId string   `json:"process_instance_id"`
	ActivityId        string   `json:"activity_id"`
	ActivityIdList    []string `json:"activity_id_list,omitempty"`
}
