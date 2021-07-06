// Code generated by lark_sdk_gen. DO NOT EDIT.

package lark

import (
	"context"
)

// QuerySheetFilterViewCondition
//
// ::: note
// 筛选条件含义可参考 [筛选视图的筛选条件指南](https://open.feishu.cn/document/ukTMukTMukTM/uUDN04SN0QjL1QDN/sheets-v3/spreadsheet-sheet-filter_view-condition/filter-view-condition-user-guide)
// :::
// 查询一个筛选视图的所有筛选条件，返回筛选视图的筛选范围内的筛选条件。
//
// doc: https://open.feishu.cn/document/ukTMukTMukTM/uUDN04SN0QjL1QDN/sheets-v3/spreadsheet-sheet-filter_view-condition/query
func (r *DriveService) QuerySheetFilterViewCondition(ctx context.Context, request *QuerySheetFilterViewConditionReq, options ...MethodOptionFunc) (*QuerySheetFilterViewConditionResp, *Response, error) {
	if r.cli.mock.mockDriveQuerySheetFilterViewCondition != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] Drive#QuerySheetFilterViewCondition mock enable")
		return r.cli.mock.mockDriveQuerySheetFilterViewCondition(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "Drive",
		API:                   "QuerySheetFilterViewCondition",
		Method:                "GET",
		URL:                   "https://open.feishu.cn/open-apis/sheets/v3/spreadsheets/:spreadsheet_token/sheets/:sheet_id/filter_views/:filter_view_id/conditions/query",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
		NeedUserAccessToken:   true,
	}
	resp := new(querySheetFilterViewConditionResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

func (r *Mock) MockDriveQuerySheetFilterViewCondition(f func(ctx context.Context, request *QuerySheetFilterViewConditionReq, options ...MethodOptionFunc) (*QuerySheetFilterViewConditionResp, *Response, error)) {
	r.mockDriveQuerySheetFilterViewCondition = f
}

func (r *Mock) UnMockDriveQuerySheetFilterViewCondition() {
	r.mockDriveQuerySheetFilterViewCondition = nil
}

type QuerySheetFilterViewConditionReq struct {
	SpreadSheetToken string `path:"spreadsheet_token" json:"-"` // 表格 token, 示例值："shtcnmBA*****yGehy8"
	SheetID          string `path:"sheet_id" json:"-"`          // 子表 id, 示例值："0b**12"
	FilterViewID     string `path:"filter_view_id" json:"-"`    // 筛选视图 id, 示例值："pH9hbVcCXA"
}

type querySheetFilterViewConditionResp struct {
	Code int64                              `json:"code,omitempty"` // 错误码，非 0 表示失败
	Msg  string                             `json:"msg,omitempty"`  // 错误描述
	Data *QuerySheetFilterViewConditionResp `json:"data,omitempty"`
}

type QuerySheetFilterViewConditionResp struct {
	Items []*QuerySheetFilterViewConditionRespItem `json:"items,omitempty"` // 筛选视图设置的所有筛选条件
}

type QuerySheetFilterViewConditionRespItem struct {
	ConditionID string   `json:"condition_id,omitempty"` // 设置筛选条件的列，使用字母号
	FilterType  string   `json:"filter_type,omitempty"`  // 筛选类型
	CompareType string   `json:"compare_type,omitempty"` // 比较类型
	Expected    []string `json:"expected,omitempty"`     // 筛选参数
}
