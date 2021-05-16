// Code generated by lark_sdk_gen. DO NOT EDIT.

package lark

import (
	"context"
)

// CreateComment 往云文档添加一条评论。
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/drive-v1/file-comment/create
func (r *DriveAPI) CreateComment(ctx context.Context, request *CreateCommentReq, options ...MethodOptionFunc) (*CreateCommentResp, *Response, error) {
	if r.cli.mock.mockDriveCreateComment != nil {
		r.cli.logDebug(ctx, "[lark] Drive#CreateComment mock enable")
		return r.cli.mock.mockDriveCreateComment(ctx, request, options...)
	}

	r.cli.logInfo(ctx, "[lark] Drive#CreateComment call api")
	r.cli.logDebug(ctx, "[lark] Drive#CreateComment request: %s", jsonString(request))

	req := &RawRequestReq{
		Method:                "POST",
		URL:                   "https://open.feishu.cn/open-apis/drive/v1/files/:file_token/comments",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
		NeedUserAccessToken:   true,
	}
	resp := new(createCommentResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	if err != nil {
		r.cli.logError(ctx, "[lark] Drive#CreateComment POST https://open.feishu.cn/open-apis/drive/v1/files/:file_token/comments failed: %s", err)
		return nil, response, err
	} else if resp.Code != 0 {
		r.cli.logError(ctx, "[lark] Drive#CreateComment POST https://open.feishu.cn/open-apis/drive/v1/files/:file_token/comments failed, code: %d, msg: %s", resp.Code, resp.Msg)
		return nil, response, NewError("Drive", "CreateComment", resp.Code, resp.Msg)
	}

	r.cli.logDebug(ctx, "[lark] Drive#CreateComment request_id: %s, response: %s", response.RequestID, jsonString(resp.Data))

	return resp.Data, response, nil
}

func (r *Mock) MockDriveCreateComment(f func(ctx context.Context, request *CreateCommentReq, options ...MethodOptionFunc) (*CreateCommentResp, *Response, error)) {
	r.mockDriveCreateComment = f
}

func (r *Mock) UnMockDriveCreateComment() {
	r.mockDriveCreateComment = nil
}

type CreateCommentReq struct {
	FileType     FileType                   `query:"file_type" json:"-"`      // 文档类型, 示例值："doc", 可选值有: `doc`：文档, `sheet`：表格, `file`：文件
	UserIDType   *IDType                    `query:"user_id_type" json:"-"`   // 用户 ID 类型, 示例值："open_id", 可选值有: `open_id`：用户的 open id, `union_id`：用户的 union id, `user_id`：用户的 user id, 默认值: `open_id`, 当值为 `user_id`, 字段权限要求: 获取用户 userid
	FileToken    string                     `path:"file_token" json:"-"`      // 文档token, 示例值："doccnGp4UK1UskrOEJwBXd3****"
	CommentID    *string                    `json:"comment_id,omitempty"`     // 评论ID, 示例值："6916106822734578184"
	UserID       *string                    `json:"user_id,omitempty"`        // 用户ID, 示例值："ou_cc19b2bfb93f8a44db4b4d6eab*****"
	CreateTime   *int                       `json:"create_time,omitempty"`    // 创建时间, 示例值：1610281603
	UpdateTime   *int                       `json:"update_time,omitempty"`    // 更新时间, 示例值：1610281603
	IsSolved     *bool                      `json:"is_solved,omitempty"`      // 是否已解决, 示例值：false
	SolvedTime   *int                       `json:"solved_time,omitempty"`    // 解决评论时间, 示例值：1610281603
	SolverUserID *string                    `json:"solver_user_id,omitempty"` // 解决评论者的用户ID, 示例值："null"
	ReplyList    *CreateCommentReqReplyList `json:"reply_list,omitempty"`     // 评论里的回复列表
}

type CreateCommentReqReplyList struct {
	Replies []*CreateCommentReqReplyListReplie `json:"replies,omitempty"` // 回复列表
}

type CreateCommentReqReplyListReplie struct {
	ReplyID    *string                                 `json:"reply_id,omitempty"`    // 回复ID, 示例值："6916106822734594568"
	UserID     *string                                 `json:"user_id,omitempty"`     // 用户ID, 示例值："ou_cc19b2bfb93f8a44db4b4d6eab2*****"
	CreateTime *int                                    `json:"create_time,omitempty"` // 创建时间, 示例值：1610281603
	UpdateTime *int                                    `json:"update_time,omitempty"` // 更新时间, 示例值：1610281603
	Content    *CreateCommentReqReplyListReplieContent `json:"content,omitempty"`     // 回复内容
}

type CreateCommentReqReplyListReplieContent struct {
	Elements []*CreateCommentReqReplyListReplieContentElement `json:"elements,omitempty"` // 回复的内容
}

type CreateCommentReqReplyListReplieContentElement struct {
	Type     string                                                 `json:"type,omitempty"`      // 回复的内容元素, 示例值："text_run", 可选值有: `text_run`：普通文本, `docs_link`：at 云文档链接, `person`：at 联系人
	TextRun  *CreateCommentReqReplyListReplieContentElementTextRun  `json:"text_run,omitempty"`  // 文本内容
	DocsLink *CreateCommentReqReplyListReplieContentElementDocsLink `json:"docs_link,omitempty"` // 文本内容
	Person   *CreateCommentReqReplyListReplieContentElementPerson   `json:"person,omitempty"`    // 文本内容
}

type CreateCommentReqReplyListReplieContentElementTextRun struct {
	Text string `json:"text,omitempty"` // 回复 普通文本, 示例值："comment text"
}

type CreateCommentReqReplyListReplieContentElementDocsLink struct {
	URL string `json:"url,omitempty"` // 回复 at云文档, 示例值："https://bytedance.feishu.cn/docs/doccnHh7U87HOFpii5u5G*****"
}

type CreateCommentReqReplyListReplieContentElementPerson struct {
	UserID string `json:"user_id,omitempty"` // 回复 at联系人, 示例值："ou_cc19b2bfb93f8a44db4b4d6eab*****"
}

type createCommentResp struct {
	Code int                `json:"code,omitempty"` // 错误码，非 0 表示失败
	Msg  string             `json:"msg,omitempty"`  // 错误描述
	Data *CreateCommentResp `json:"data,omitempty"` //
}

type CreateCommentResp struct {
	CommentID    string                      `json:"comment_id,omitempty"`     // 评论ID
	UserID       string                      `json:"user_id,omitempty"`        // 用户ID
	CreateTime   int                         `json:"create_time,omitempty"`    // 创建时间
	UpdateTime   int                         `json:"update_time,omitempty"`    // 更新时间
	IsSolved     bool                        `json:"is_solved,omitempty"`      // 是否已解决
	SolvedTime   int                         `json:"solved_time,omitempty"`    // 解决评论时间
	SolverUserID string                      `json:"solver_user_id,omitempty"` // 解决评论者的用户ID
	ReplyList    *CreateCommentRespReplyList `json:"reply_list,omitempty"`     // 评论里的回复列表
}

type CreateCommentRespReplyList struct {
	Replies []*CreateCommentRespReplyListReplie `json:"replies,omitempty"` // 回复列表
}

type CreateCommentRespReplyListReplie struct {
	ReplyID    string                                   `json:"reply_id,omitempty"`    // 回复ID
	UserID     string                                   `json:"user_id,omitempty"`     // 用户ID
	CreateTime int                                      `json:"create_time,omitempty"` // 创建时间
	UpdateTime int                                      `json:"update_time,omitempty"` // 更新时间
	Content    *CreateCommentRespReplyListReplieContent `json:"content,omitempty"`     // 回复内容
}

type CreateCommentRespReplyListReplieContent struct {
	Elements []*CreateCommentRespReplyListReplieContentElement `json:"elements,omitempty"` // 回复的内容
}

type CreateCommentRespReplyListReplieContentElement struct {
	Type     string                                                  `json:"type,omitempty"`      // 回复的内容元素, 可选值有: `text_run`：普通文本, `docs_link`：at 云文档链接, `person`：at 联系人
	TextRun  *CreateCommentRespReplyListReplieContentElementTextRun  `json:"text_run,omitempty"`  // 文本内容
	DocsLink *CreateCommentRespReplyListReplieContentElementDocsLink `json:"docs_link,omitempty"` // 文本内容
	Person   *CreateCommentRespReplyListReplieContentElementPerson   `json:"person,omitempty"`    // 文本内容
}

type CreateCommentRespReplyListReplieContentElementTextRun struct {
	Text string `json:"text,omitempty"` // 回复 普通文本
}

type CreateCommentRespReplyListReplieContentElementDocsLink struct {
	URL string `json:"url,omitempty"` // 回复 at云文档
}

type CreateCommentRespReplyListReplieContentElementPerson struct {
	UserID string `json:"user_id,omitempty"` // 回复 at联系人
}
