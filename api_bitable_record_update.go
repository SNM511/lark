// Code generated by lark_sdk_gen. DO NOT EDIT.

package lark

import (
	"context"
)

// UpdateRecord 该接口用于更新数据表中的一条记录
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/bitable-v1/app-table-record/update
func (r *BitableAPI) UpdateRecord(ctx context.Context, request *UpdateRecordReq, options ...MethodOptionFunc) (*UpdateRecordResp, *Response, error) {
	if r.cli.mock.mockBitableUpdateRecord != nil {
		r.cli.logDebug(ctx, "[lark] Bitable#UpdateRecord mock enable")
		return r.cli.mock.mockBitableUpdateRecord(ctx, request, options...)
	}

	r.cli.logInfo(ctx, "[lark] Bitable#UpdateRecord call api")
	r.cli.logDebug(ctx, "[lark] Bitable#UpdateRecord request: %s", jsonString(request))

	req := &RawRequestReq{
		Method:              "PUT",
		URL:                 "https://open.feishu.cn/open-apis/bitable/v1/apps/:app_token/tables/:table_id/records/:record_id",
		Body:                request,
		MethodOption:        newMethodOption(options),
		NeedUserAccessToken: true,
	}
	resp := new(updateRecordResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	if err != nil {
		r.cli.logError(ctx, "[lark] Bitable#UpdateRecord PUT https://open.feishu.cn/open-apis/bitable/v1/apps/:app_token/tables/:table_id/records/:record_id failed: %s", err)
		return nil, response, err
	} else if resp.Code != 0 {
		r.cli.logError(ctx, "[lark] Bitable#UpdateRecord PUT https://open.feishu.cn/open-apis/bitable/v1/apps/:app_token/tables/:table_id/records/:record_id failed, code: %d, msg: %s", resp.Code, resp.Msg)
		return nil, response, NewError("Bitable", "UpdateRecord", resp.Code, resp.Msg)
	}

	r.cli.logDebug(ctx, "[lark] Bitable#UpdateRecord request_id: %s, response: %s", response.RequestID, jsonString(resp.Data))

	return resp.Data, response, nil
}

func (r *Mock) MockBitableUpdateRecord(f func(ctx context.Context, request *UpdateRecordReq, options ...MethodOptionFunc) (*UpdateRecordResp, *Response, error)) {
	r.mockBitableUpdateRecord = f
}

func (r *Mock) UnMockBitableUpdateRecord() {
	r.mockBitableUpdateRecord = nil
}

type UpdateRecordReq struct {
	UserIDType *IDType                `query:"user_id_type" json:"-"` // 用户 ID 类型, 示例值："open_id", 可选值有: `open_id`：用户的 open id, `union_id`：用户的 union id, `user_id`：用户的 user id, 默认值: `open_id`, 当值为 `user_id`, 字段权限要求: 获取用户 userid
	AppToken   string                 `path:"app_token" json:"-"`     // bitable app token, 示例值："appbcbWCzen6D8dezhoCH2RpMAh"
	TableID    string                 `path:"table_id" json:"-"`      // table id, 示例值："tblsRc9GRRXKqhvW"
	RecordID   string                 `path:"record_id" json:"-"`     // 单条记录的 id, 示例值："recqwIwhc6"
	Fields     map[string]interface{} `json:"fields,omitempty"`       // 记录字段
}

type updateRecordResp struct {
	Code int               `json:"code,omitempty"` // 错误码，非 0 表示失败
	Msg  string            `json:"msg,omitempty"`  // 错误描述
	Data *UpdateRecordResp `json:"data,omitempty"` //
}

type UpdateRecordResp struct {
	Record *UpdateRecordRespRecord `json:"record,omitempty"` // {,    "fields": {,        "人力评估": 2,,        "任务执行人": [,            {,                "id": "ou_debc524b2d8cb187704df652b43d29de",            },        ],,        "任务描述": "多渠道收集用户反馈",,        "对应 OKR": [,            "recqwIwhc6",,            "recOuEJMvN",        ],,        "截止日期": 1609516800000,,        "是否完成": true,,        "状态": "已结束",,        "相关部门": [,            "销售",,            "客服",        ],    },}
}

type UpdateRecordRespRecord struct {
	RecordID string                 `json:"record_id,omitempty"` // 记录 id
	Fields   map[string]interface{} `json:"fields,omitempty"`    // 记录字段
}
