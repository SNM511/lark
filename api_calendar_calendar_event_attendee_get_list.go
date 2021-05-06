package lark

import (
	"context"
)

// GetCalendarEventAttendeeList 获取日程的参与人列表，若参与者列表中有群组，请使用 [获取参与人群成员列表](/ssl:ttdoc/uAjLw4CM/ukTMukTMukTM/reference/calendar-v4/calendar-event-attendee-chat_member/list) 。
//
// - 当前身份必须对日历有访问权限。
// - 当前身份必须有权限查看日程的参与人列表。
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/calendar-v4/calendar-event-attendee/list
func (r *CalendarAPI) GetCalendarEventAttendeeList(ctx context.Context, request *GetCalendarEventAttendeeListReq) (*GetCalendarEventAttendeeListResp, *Response, error) {
	req := &requestParam{
		Method:                "GET",
		URL:                   "https://open.feishu.cn/open-apis/calendar/v4/calendars/:calendar_id/events/:event_id/attendees",
		Body:                  request,
		NeedTenantAccessToken: true,
		NeedAppAccessToken:    false,
		NeedHelpdeskAuth:      false,
		IsFile:                false,
	}
	resp := new(getCalendarEventAttendeeListResp)

	response, err := r.cli.request(ctx, req, resp)
	if err != nil {
		return nil, response, err
	} else if resp.Code != 0 {
		return nil, response, newError("Calendar", "GetCalendarEventAttendeeList", resp.Code, resp.Msg)
	}

	return resp.Data, response, nil
}

type GetCalendarEventAttendeeListReq struct {
	UserIDType *IDType `query:"user_id_type" json:"-"` // 用户 ID 类型, 示例值："open_id", 可选值有: `open_id`：用户的 open id, `union_id`：用户的 union id, `user_id`：用户的 user id, 默认值: `open_id`, 当值为 `user_id`, 字段权限要求: 获取用户 userid
	PageToken  *string `query:"page_token" json:"-"`   // 分页标记，第一次请求不填，表示从头开始遍历；分页查询结果还有更多项时会同时返回新的 page_token，下次遍历可采用该 page_token 获取查询结果, 示例值："780TRhwXXXXX"
	PageSize   *int    `query:"page_size" json:"-"`    // 分页大小, 示例值：10, 最大值：`100`
	CalendarID string  `path:"calendar_id" json:"-"`   // 日历 ID, 示例值："feishu.cn_xxxxxxxxxx@group.calendar.feishu.cn"
	EventID    string  `path:"event_id" json:"-"`      // 日程 ID, 示例值："xxxxxxxxx_0"
}

type getCalendarEventAttendeeListResp struct {
	Code int                               `json:"code,omitempty"` // 错误码，非 0 表示失败
	Msg  string                            `json:"msg,omitempty"`  // 错误描述
	Data *GetCalendarEventAttendeeListResp `json:"data,omitempty"` //
}

type GetCalendarEventAttendeeListResp struct {
	Items     []*GetCalendarEventAttendeeListRespItem `json:"items,omitempty"`      // 日程的参与者列表
	HasMore   bool                                    `json:"has_more,omitempty"`   // 是否还有更多项
	PageToken string                                  `json:"page_token,omitempty"` // 分页标记，当 has_more 为 true 时，会同时返回新的 page_token，否则不返回 page_token
}

type GetCalendarEventAttendeeListRespItem struct {
	Type            CalendarEventAttendeeType                         `json:"type,omitempty"`              // 参与人类型, 可选值有: `user`：用户, `chat`：群组, `resource`：会议室, `third_party`：邮箱
	AttendeeID      string                                            `json:"attendee_id,omitempty"`       // 参与人ID
	RsvpStatus      string                                            `json:"rsvp_status,omitempty"`       // 参与人RSVP状态, 可选值有: `needs_action`：参与人尚未回复状态，或表示会议室预约中, `accept`：参与人回复接受，或表示会议室预约成功, `tentative`：参与人回复待定, `decline`：参与人回复拒绝，或表示会议室预约失败, `removed`：参与人或会议室已经从日程中被移除
	IsOptional      bool                                              `json:"is_optional,omitempty"`       // 参与人是否为「可选参加」，无法编辑群参与人的此字段, 默认值: `false`
	IsOrganizer     bool                                              `json:"is_organizer,omitempty"`      // 参与人是否为日程组织者
	IsExternal      bool                                              `json:"is_external,omitempty"`       // 参与人是否为外部参与人；外部参与人不支持编辑
	DisplayName     string                                            `json:"display_name,omitempty"`      // 参与人名称
	ChatMembers     []*GetCalendarEventAttendeeListRespItemChatMember `json:"chat_members,omitempty"`      // 群中的群成员，当type为Chat时有效；群成员不支持编辑
	UserID          string                                            `json:"user_id,omitempty"`           // 参与人的用户id，依赖于user_id_type返回对应的取值，当is_external为true时，此字段只会返回open_id或者union_id
	ChatID          string                                            `json:"chat_id,omitempty"`           // chat类型参与人的群组chat_id
	RoomID          string                                            `json:"room_id,omitempty"`           // resource类型参与人的会议室room_id
	ThirdPartyEmail string                                            `json:"third_party_email,omitempty"` // third_party类型参与人的邮箱
}

type GetCalendarEventAttendeeListRespItemChatMember struct {
	RsvpStatus  string `json:"rsvp_status,omitempty"`  // 参与人RSVP状态, 可选值有: `needs_action`：参与人尚未回复状态，或表示会议室预约中, `accept`：参与人回复接受，或表示会议室预约成功, `tentative`：参与人回复待定, `decline`：参与人回复拒绝，或表示会议室预约失败, `removed`：参与人或会议室已经从日程中被移除
	IsOptional  bool   `json:"is_optional,omitempty"`  // 参与人是否为「可选参加」, 默认值: `false`
	DisplayName string `json:"display_name,omitempty"` // 参与人名称
	IsOrganizer bool   `json:"is_organizer,omitempty"` // 参与人是否为日程组织者
	IsExternal  bool   `json:"is_external,omitempty"`  // 参与人是否为外部参与人
}
