// Code generated by lark_sdk_gen. DO NOT EDIT.

package lark

import (
	"context"
)

// GetFolderChildren 该接口用于根据 folderToken 获取该文件夹的文档清单，如 doc、sheet、folder。
//
// doc: https://open.feishu.cn/document/ukTMukTMukTM/uEjNzUjLxYzM14SM2MTN
func (r *DriveAPI) GetFolderChildren(ctx context.Context, request *GetFolderChildrenReq, options ...MethodOptionFunc) (*GetFolderChildrenResp, *Response, error) {
	if r.cli.mock.mockDriveGetFolderChildren != nil {
		r.cli.logDebug(ctx, "[lark] Drive#GetFolderChildren mock enable")
		return r.cli.mock.mockDriveGetFolderChildren(ctx, request, options...)
	}

	r.cli.logInfo(ctx, "[lark] Drive#GetFolderChildren call api")
	r.cli.logDebug(ctx, "[lark] Drive#GetFolderChildren request: %s", jsonString(request))

	req := &RawRequestReq{
		Method:              "GET",
		URL:                 "https://open.feishu.cn/open-apis/drive/explorer/v2/folder/:folderToken/children",
		Body:                request,
		MethodOption:        newMethodOption(options),
		NeedUserAccessToken: true,
	}
	resp := new(getFolderChildrenResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	if err != nil {
		r.cli.logError(ctx, "[lark] Drive#GetFolderChildren GET https://open.feishu.cn/open-apis/drive/explorer/v2/folder/:folderToken/children failed: %s", err)
		return nil, response, err
	} else if resp.Code != 0 {
		r.cli.logError(ctx, "[lark] Drive#GetFolderChildren GET https://open.feishu.cn/open-apis/drive/explorer/v2/folder/:folderToken/children failed, code: %d, msg: %s", resp.Code, resp.Msg)
		return nil, response, NewError("Drive", "GetFolderChildren", resp.Code, resp.Msg)
	}

	r.cli.logDebug(ctx, "[lark] Drive#GetFolderChildren request_id: %s, response: %s", response.RequestID, jsonString(resp.Data))

	return resp.Data, response, nil
}

func (r *Mock) MockDriveGetFolderChildren(f func(ctx context.Context, request *GetFolderChildrenReq, options ...MethodOptionFunc) (*GetFolderChildrenResp, *Response, error)) {
	r.mockDriveGetFolderChildren = f
}

func (r *Mock) UnMockDriveGetFolderChildren() {
	r.mockDriveGetFolderChildren = nil
}

type GetFolderChildrenReq struct {
	Types       []string `query:"types" json:"-"`      // 需要查询的文件类型，默认返回所有 children；types 可多选，可选类型有 doc、sheet、file、folder 。如 url?types=folder&types=sheet
	FolderToken string   `path:"folderToken" json:"-"` // 文件夹的 token，获取方式见 [对接前说明](/ssl:ttdoc/ukTMukTMukTM/uczNzUjL3czM14yN3MTN) 的第 3 项
}

type getFolderChildrenResp struct {
	Code int                    `json:"code,omitempty"`
	Msg  string                 `json:"msg,omitempty"`
	Data *GetFolderChildrenResp `json:"data,omitempty"`
}

type GetFolderChildrenResp struct {
	ParentToken string                         `json:"parentToken,omitempty"` // 文件夹的 token
	Children    *GetFolderChildrenRespChildren `json:"children,omitempty"`    // 文件夹的下的文件
}

type GetFolderChildrenRespChildren struct {
	Token string `json:"token,omitempty"` // 文件的 token
	Name  string `json:"name,omitempty"`  // 文件的标题
	Type  string `json:"type,omitempty"`  // 文件的类型
}
