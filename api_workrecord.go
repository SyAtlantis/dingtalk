package dingtalk

import (
	"github.com/SyAtlantis/dingtalk/v3/constant"
	"github.com/SyAtlantis/dingtalk/v3/request"
	"github.com/SyAtlantis/dingtalk/v3/response"
	"net/http"
)

// SaveApprovalTemplate 创建或更新审批模板
func (ding *DingTalk) SaveApprovalTemplate(res *request.SaveApprovalTemplate) (rsp response.SaveApprovalTemplate, err error) {
	return rsp, ding.Request(http.MethodPost, constant.SaveApprovalTemplate, nil, res, &rsp)
}

// GetApprovalTemplate 获取模板code
func (ding *DingTalk) GetApprovalTemplate(res *request.GetApprovalTemplate) (rsp response.GetApprovalTemplate, err error) {
	return rsp, ding.Request(http.MethodPost, constant.GetApprovalTemplate, nil, res, &rsp)
}

// DeleteApprovalTemplate 删除模板
func (ding *DingTalk) DeleteApprovalTemplate(res *request.DeleteApprovalTemplate) (rsp response.DeleteApprovalTemplate, err error) {
	return rsp, ding.Request(http.MethodPost, constant.DeleteApprovalTemplate, nil, res, &rsp)
}

// InitiateApprovalProcess 创建实例
func (ding *DingTalk) InitiateApprovalProcess(res *request.InitiateApprovalProcess) (rsp response.InitiateApprovalProcess, err error) {
	return rsp, ding.Request(http.MethodPost, constant.InitiateApprovalProcess, nil, res, &rsp)
}

// UpdateApprovalProcess 更新实例状态
func (ding *DingTalk) UpdateApprovalProcess(res *request.UpdateApprovalProcess) (rsp response.UpdateApprovalProcess, err error) {
	return rsp, ding.Request(http.MethodPost, constant.UpdateApprovalProcess, nil, res, &rsp)
}

// BatchUpdateApprovalProcess 批量更新实例状态
func (ding *DingTalk) BatchUpdateApprovalProcess(res *request.BatchUpdateApprovalProcess) (rsp response.BatchUpdateApprovalProcess, err error) {
	return rsp, ding.Request(http.MethodPost, constant.BatchUpdateApprovalProcess, nil, res, &rsp)
}

// CreateWorkrecord 创建待办事项
func (ding *DingTalk) CreateWorkrecord(res *request.CreateWorkrecord) (rsp response.CreateWorkrecord, err error) {
	return rsp, ding.Request(http.MethodPost, constant.CreateWorkrecord, nil, res, &rsp)
}

// QueryWorkrecordList 查询待办列表
func (ding *DingTalk) QueryWorkrecordList(res *request.QueryWorkrecordList) (rsp response.QueryWorkrecordList, err error) {
	return rsp, ding.Request(http.MethodPost, constant.QueryWorkrecordList, nil, res, &rsp)
}

// UpdateWorkrecord 更新待办状态
func (ding *DingTalk) UpdateWorkrecord(res *request.UpdateWorkrecord) (rsp response.UpdateWorkrecord, err error) {
	return rsp, ding.Request(http.MethodPost, constant.UpdateWorkrecord, nil, res, &rsp)
}

// CancelWorkrecord 批量取消待办
func (ding *DingTalk) CancelWorkrecord(res *request.CancelWorkrecord) (rsp response.CancelWorkrecord, err error) {
	return rsp, ding.Request(http.MethodPost, constant.CancelWorkrecord, nil, res, &rsp)
}
