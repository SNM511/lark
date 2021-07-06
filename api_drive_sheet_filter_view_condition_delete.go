// Code generated by lark_sdk_gen. DO NOT EDIT.

package lark

import (
	"context"
)

// DeleteSheetFilterViewCondition 删除筛选视图的筛选范围某一列的筛选条件。
//
// doc: https://open.feishu.cn/document/ukTMukTMukTM/uUDN04SN0QjL1QDN/sheets-v3/spreadsheet-sheet-filter_view-condition/delete
func (r *DriveService) DeleteSheetFilterViewCondition(ctx context.Context, request *DeleteSheetFilterViewConditionReq, options ...MethodOptionFunc) (*DeleteSheetFilterViewConditionResp, *Response, error) {
	if r.cli.mock.mockDriveDeleteSheetFilterViewCondition != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] Drive#DeleteSheetFilterViewCondition mock enable")
		return r.cli.mock.mockDriveDeleteSheetFilterViewCondition(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "Drive",
		API:                   "DeleteSheetFilterViewCondition",
		Method:                "DELETE",
		URL:                   "https://open.feishu.cn/open-apis/sheets/v3/spreadsheets/:spreadsheet_token/sheets/:sheet_id/filter_views/:filter_view_id/conditions/:condition_id",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
		NeedUserAccessToken:   true,
	}
	resp := new(deleteSheetFilterViewConditionResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

func (r *Mock) MockDriveDeleteSheetFilterViewCondition(f func(ctx context.Context, request *DeleteSheetFilterViewConditionReq, options ...MethodOptionFunc) (*DeleteSheetFilterViewConditionResp, *Response, error)) {
	r.mockDriveDeleteSheetFilterViewCondition = f
}

func (r *Mock) UnMockDriveDeleteSheetFilterViewCondition() {
	r.mockDriveDeleteSheetFilterViewCondition = nil
}

type DeleteSheetFilterViewConditionReq struct {
	SpreadSheetToken string `path:"spreadsheet_token" json:"-"` // 表格 token, 示例值："shtcnmBA*****yGehy8"
	SheetID          string `path:"sheet_id" json:"-"`          // 子表 id, 示例值："0b**12"
	FilterViewID     string `path:"filter_view_id" json:"-"`    // 筛选视图 id, 示例值："pH9hbVcCXA"
	ConditionID      string `path:"condition_id" json:"-"`      // 筛选范围内的某列字母号, 示例值："E"
}

type deleteSheetFilterViewConditionResp struct {
	Code int64                               `json:"code,omitempty"` // 错误码，非 0 表示失败
	Msg  string                              `json:"msg,omitempty"`  // 错误描述
	Data *DeleteSheetFilterViewConditionResp `json:"data,omitempty"`
}

type DeleteSheetFilterViewConditionResp struct{}
