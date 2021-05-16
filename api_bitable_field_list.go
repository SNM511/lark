// Code generated by lark_sdk_gen. DO NOT EDIT.

package lark

import (
	"context"
)

// GetFieldList 根据 app_token 和 table_id，获取数据表的所有字段
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/bitable-v1/app-table-field/list
func (r *BitableAPI) GetFieldList(ctx context.Context, request *GetFieldListReq, options ...MethodOptionFunc) (*GetFieldListResp, *Response, error) {
	if r.cli.mock.mockBitableGetFieldList != nil {
		r.cli.logDebug(ctx, "[lark] Bitable#GetFieldList mock enable")
		return r.cli.mock.mockBitableGetFieldList(ctx, request, options...)
	}

	r.cli.logInfo(ctx, "[lark] Bitable#GetFieldList call api")
	r.cli.logDebug(ctx, "[lark] Bitable#GetFieldList request: %s", jsonString(request))

	req := &RawRequestReq{
		Method:              "GET",
		URL:                 "https://open.feishu.cn/open-apis/bitable/v1/apps/:app_token/tables/:table_id/fields",
		Body:                request,
		MethodOption:        newMethodOption(options),
		NeedUserAccessToken: true,
	}
	resp := new(getFieldListResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	if err != nil {
		r.cli.logError(ctx, "[lark] Bitable#GetFieldList GET https://open.feishu.cn/open-apis/bitable/v1/apps/:app_token/tables/:table_id/fields failed: %s", err)
		return nil, response, err
	} else if resp.Code != 0 {
		r.cli.logError(ctx, "[lark] Bitable#GetFieldList GET https://open.feishu.cn/open-apis/bitable/v1/apps/:app_token/tables/:table_id/fields failed, code: %d, msg: %s", resp.Code, resp.Msg)
		return nil, response, NewError("Bitable", "GetFieldList", resp.Code, resp.Msg)
	}

	r.cli.logDebug(ctx, "[lark] Bitable#GetFieldList request_id: %s, response: %s", response.RequestID, jsonString(resp.Data))

	return resp.Data, response, nil
}

func (r *Mock) MockBitableGetFieldList(f func(ctx context.Context, request *GetFieldListReq, options ...MethodOptionFunc) (*GetFieldListResp, *Response, error)) {
	r.mockBitableGetFieldList = f
}

func (r *Mock) UnMockBitableGetFieldList() {
	r.mockBitableGetFieldList = nil
}

type GetFieldListReq struct {
	ViewID    *string `query:"view_id" json:"-"`    // 视图 ID, 示例值："vewOVMEXPF"
	PageToken *string `query:"page_token" json:"-"` // 分页标记，第一次请求不填，表示从头开始遍历；分页查询结果还有更多项时会同时返回新的 page_token，下次遍历可采用该 page_token 获取查询结果, 示例值："fldwJ4YrtB"
	PageSize  *int    `query:"page_size" json:"-"`  // 分页大小, 示例值：10, 最大值：`100`
	AppToken  string  `path:"app_token" json:"-"`   // bitable app token, 示例值："appbcbWCzen6D8dezhoCH2RpMAh"
	TableID   string  `path:"table_id" json:"-"`    // table id, 示例值："tblsRc9GRRXKqhvW"
}

type getFieldListResp struct {
	Code int               `json:"code,omitempty"` // 错误码，非 0 表示失败
	Msg  string            `json:"msg,omitempty"`  // 错误描述
	Data *GetFieldListResp `json:"data,omitempty"` //
}

type GetFieldListResp struct {
	HasMore   bool                    `json:"has_more,omitempty"`   // 是否还有更多项
	PageToken string                  `json:"page_token,omitempty"` // 分页标记，当 has_more 为 true 时，会同时返回新的 page_token，否则不返回 page_token
	Items     []*GetFieldListRespItem `json:"items,omitempty"`      // 字段信息
}

type GetFieldListRespItem struct {
	FieldID   string      `json:"field_id,omitempty"`   // 多维表格字段 id
	FieldName string      `json:"field_name,omitempty"` // 多维表格字段名
	Type      int         `json:"type,omitempty"`       // 多维表格字段类型
	Property  interface{} `json:"property,omitempty"`   // 字段属性
}
