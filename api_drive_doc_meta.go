// Code generated by lark_sdk_gen. DO NOT EDIT.

package lark

import (
	"context"
)

// GetDocMeta 该接口用于根据 token 获取各类文件的元数据。
//
// 请求用户需要拥有该文件的访问（读）权限
//
// doc: https://open.feishu.cn/document/ukTMukTMukTM/uMjN3UjLzYzN14yM2cTN
func (r *DriveAPI) GetDocMeta(ctx context.Context, request *GetDocMetaReq, options ...MethodOptionFunc) (*GetDocMetaResp, *Response, error) {
	if r.cli.mock.mockDriveGetDocMeta != nil {
		r.cli.logDebug(ctx, "[lark] Drive#GetDocMeta mock enable")
		return r.cli.mock.mockDriveGetDocMeta(ctx, request, options...)
	}

	r.cli.logInfo(ctx, "[lark] Drive#GetDocMeta call api")
	r.cli.logDebug(ctx, "[lark] Drive#GetDocMeta request: %s", jsonString(request))

	req := &RawRequestReq{
		Method:              "POST",
		URL:                 "https://open.feishu.cn/open-apis/suite/docs-api/meta",
		Body:                request,
		MethodOption:        newMethodOption(options),
		NeedUserAccessToken: true,
	}
	resp := new(getDocMetaResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	if err != nil {
		r.cli.logError(ctx, "[lark] Drive#GetDocMeta POST https://open.feishu.cn/open-apis/suite/docs-api/meta failed: %s", err)
		return nil, response, err
	} else if resp.Code != 0 {
		r.cli.logError(ctx, "[lark] Drive#GetDocMeta POST https://open.feishu.cn/open-apis/suite/docs-api/meta failed, code: %d, msg: %s", resp.Code, resp.Msg)
		return nil, response, NewError("Drive", "GetDocMeta", resp.Code, resp.Msg)
	}

	r.cli.logDebug(ctx, "[lark] Drive#GetDocMeta request_id: %s, response: %s", response.RequestID, jsonString(resp.Data))

	return resp.Data, response, nil
}

func (r *Mock) MockDriveGetDocMeta(f func(ctx context.Context, request *GetDocMetaReq, options ...MethodOptionFunc) (*GetDocMetaResp, *Response, error)) {
	r.mockDriveGetDocMeta = f
}

func (r *Mock) UnMockDriveGetDocMeta() {
	r.mockDriveGetDocMeta = nil
}

type GetDocMetaReq struct {
	RequestDocs *GetDocMetaReqRequestDocs `json:"request_docs,omitempty"` // 请求文档，一次不超过200个
}

type GetDocMetaReqRequestDocs struct {
	DocsToken string `json:"docs_token,omitempty"` // 文件的 token，获取方式见 [对接前说明](/ssl:ttdoc/ukTMukTMukTM/uczNzUjL3czM14yN3MTN)的第 4 项
	DocsType  string `json:"docs_type,omitempty"`  // 文件类型  "doc"  or  "sheet" or "slide"
}

type getDocMetaResp struct {
	Code int             `json:"code,omitempty"`
	Msg  string          `json:"msg,omitempty"`
	Data *GetDocMetaResp `json:"data,omitempty"`
}

type GetDocMetaResp struct {
	DocsMetas *GetDocMetaRespDocsMetas `json:"docs_metas,omitempty"` // 文件元数据
}

type GetDocMetaRespDocsMetas struct {
	DocsToken        string `json:"docs_token,omitempty"`         // 文件token
	DocsType         string `json:"docs_type,omitempty"`          // 文件类型
	Title            string `json:"title,omitempty"`              // 标题
	OwnerID          string `json:"owner_id,omitempty"`           // 文件拥有者
	CreateTime       int    `json:"create_time,omitempty"`        // 创建时间（Unix时间戳）
	LatestModifyUser string `json:"latest_modify_user,omitempty"` // 最后编辑者
	LatestModifyTime int    `json:"latest_modify_time,omitempty"` // 最后编辑时间（Unix时间戳）
}
