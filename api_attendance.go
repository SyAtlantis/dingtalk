/*
 * Copyright 2020 zhaoyunxing.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package dingtalk

import (
	"net/http"

	"github.com/SyAtlantis/dingtalk/v3/constant"
	"github.com/SyAtlantis/dingtalk/v3/request"
	"github.com/SyAtlantis/dingtalk/v3/response"
)

// GetAttendanceGroups 批量获取考勤组详情
func (ding *DingTalk) GetAttendanceGroups(offset, size int) (rsp response.GetAttendanceGroup, err error) {
	return rsp, ding.Request(http.MethodPost, constant.GetAttendanceGroupsKey, nil,
		request.NewGetAttendanceGroup(offset, size), &rsp)
}

// GetAttendanceUserGroup 获取用户考勤组
func (ding *DingTalk) GetAttendanceUserGroup(userId string) (rsp response.GetAttendanceUserGroup, err error) {
	return rsp, ding.Request(http.MethodPost, constant.GetAttendanceUserGroupKey, nil,
		request.NewGetAttendanceUserGroup(userId), &rsp)
}

// GetAttendanceGroupMinimalism 批量获取考勤组摘要
func (ding *DingTalk) GetAttendanceGroupMinimalism(userId string, cursor int) (rsp response.GetAttendanceGroupMinimalism,
	err error) {
	return rsp, ding.Request(http.MethodPost, constant.GetAttendanceGroupMinimalismKey, nil,
		request.NewGetAttendanceGroupMinimalism(userId, cursor), &rsp)
}

// GetAttendanceGroupDetail 获取考勤组详情
func (ding *DingTalk) GetAttendanceGroupDetail(userId string, groupId int) (rsp response.GetAttendanceGroupDetail,
	err error) {
	return rsp, ding.Request(http.MethodPost, constant.GetAttendanceGroupDetailKey, nil,
		request.NewGetAttendanceGroupDetail(userId, groupId), &rsp)
}

// SearchAttendanceGroup 搜索考勤组摘要
func (ding *DingTalk) SearchAttendanceGroup(userId, groupName string) (rsp response.SearchAttendanceGroup, err error) {
	return rsp, ding.Request(http.MethodPost, constant.SearchAttendanceGroupKey, nil,
		request.NewSearchAttendanceGroup(userId, groupName), &rsp)
}

// CreateAttendanceGroup 创建考勤组
func (ding *DingTalk) CreateAttendanceGroup(res *request.CreateAttendanceGroup) (rsp response.CreateAttendanceGroup, err error) {
	return rsp, ding.Request(http.MethodPost, constant.CreateAttendanceGroupKey, nil, res, &rsp)
}

// GetAttendanceColumnsKey 获取考勤报表列定义
func (ding *DingTalk) GetAttendanceColumnsKey(res *request.GetAttendanceColumnsKey) (rsp response.GetAttendanceColumnsKey, err error) {
	return rsp, ding.Request(http.MethodPost, constant.GetAttendanceColumnsKey, nil, res, &rsp)
}

// GetAttendanceColumnsValue 获取考勤报表列值
func (ding *DingTalk) GetAttendanceColumnsValue(res *request.GetAttendanceColumnsValue) (rsp response.GetAttendanceColumnsValue, err error) {
	return rsp, ding.Request(http.MethodPost, constant.GetAttendanceColumnsValue, nil, res, &rsp)
}

// GetAttendanceUpdateData 获取用户考勤数据
func (ding *DingTalk) GetAttendanceUpdateData(res *request.GetAttendanceUpdateData) (rsp response.GetAttendanceUpdateData, err error) {
	return rsp, ding.Request(http.MethodPost, constant.GetAttendanceUpdateData, nil, res, &rsp)
}
