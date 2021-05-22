// Code generated by lark_sdk_gen. DO NOT EDIT.

package lark

import (
	"context"
)

// AnswerHelpdeskTicketUserQuery 该接口用于回复用户提问结果至工单，需要工单仍处于进行中且未接入人工状态。仅支持自建应用。
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/helpdesk-v1/ticket/answer_user_query
func (r *HelpdeskService) AnswerHelpdeskTicketUserQuery(ctx context.Context, request *AnswerHelpdeskTicketUserQueryReq, options ...MethodOptionFunc) (*AnswerHelpdeskTicketUserQueryResp, *Response, error) {
	if r.cli.mock.mockHelpdeskAnswerHelpdeskTicketUserQuery != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] Helpdesk#AnswerHelpdeskTicketUserQuery mock enable")
		return r.cli.mock.mockHelpdeskAnswerHelpdeskTicketUserQuery(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "Helpdesk",
		API:                   "AnswerHelpdeskTicketUserQuery",
		Method:                "POST",
		URL:                   "https://open.feishu.cn/open-apis/helpdesk/v1/tickets/:ticket_id/answer_user_query",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
		NeedHelpdeskAuth:      true,
	}
	resp := new(answerHelpdeskTicketUserQueryResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

func (r *Mock) MockHelpdeskAnswerHelpdeskTicketUserQuery(f func(ctx context.Context, request *AnswerHelpdeskTicketUserQueryReq, options ...MethodOptionFunc) (*AnswerHelpdeskTicketUserQueryResp, *Response, error)) {
	r.mockHelpdeskAnswerHelpdeskTicketUserQuery = f
}

func (r *Mock) UnMockHelpdeskAnswerHelpdeskTicketUserQuery() {
	r.mockHelpdeskAnswerHelpdeskTicketUserQuery = nil
}

type AnswerHelpdeskTicketUserQueryReq struct {
	TicketID string                                 `path:"ticket_id" json:"-"` // 工单ID, 示例值："6945345902185807891"
	EventID  string                                 `json:"event_id,omitempty"` // 事件ID,可从订阅事件中提取, 示例值："abcd"
	Faqs     []*AnswerHelpdeskTicketUserQueryReqFaq `json:"faqs,omitempty"`     // faq结果列表
}

type AnswerHelpdeskTicketUserQueryReqFaq struct {
	ID    *string  `json:"id,omitempty"`    // faq服务台内唯一标识, 示例值："12345"
	Score *float64 `json:"score,omitempty"` // faq匹配得分, 示例值：0.9
}

type answerHelpdeskTicketUserQueryResp struct {
	Code int64                              `json:"code,omitempty"` // 错误码，非 0 表示失败
	Msg  string                             `json:"msg,omitempty"`  // 错误描述
	Data *AnswerHelpdeskTicketUserQueryResp `json:"data,omitempty"`
}

type AnswerHelpdeskTicketUserQueryResp struct{}
