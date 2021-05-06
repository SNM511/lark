package lark

import (
	"context"
)

// GetCalendarEventList 该接口用于以当前身份（应用 / 用户）获取日历下的日程列表。
//
// 当前身份必须对日历有访问权限。
// 调用时首先使用 page_token 分页拉取存量数据，之后使用 sync_token 增量同步变更数据。
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/calendar-v4/calendar-event/list
func (r *CalendarAPI) GetCalendarEventList(ctx context.Context, request *GetCalendarEventListReq) (*GetCalendarEventListResp, *Response, error) {
	req := &requestParam{
		Method:                "GET",
		URL:                   "https://open.feishu.cn/open-apis/calendar/v4/calendars/:calendar_id/events",
		Body:                  request,
		NeedTenantAccessToken: true,
		NeedAppAccessToken:    false,
		NeedHelpdeskAuth:      false,
		IsFile:                false,
	}
	resp := new(getCalendarEventListResp)

	response, err := r.cli.request(ctx, req, resp)
	if err != nil {
		return nil, response, err
	} else if resp.Code != 0 {
		return nil, response, newError("Calendar", "GetCalendarEventList", resp.Code, resp.Msg)
	}

	return resp.Data, response, nil
}

type GetCalendarEventListReq struct {
	PageSize   *int    `query:"page_size" json:"-"`  // 一次请求要求返回最大数量，默认500，取值范围为[50. 1000], 示例值：50, 默认值: `500`, 取值范围：`50` ～ `1000`
	PageToken  *string `query:"page_token" json:"-"` // 上次请求Response返回的分页标记，首次请求时为空, 示例值："xxxxx"
	SyncToken  *string `query:"sync_token" json:"-"` // 上次请求Response返回的增量同步标记，分页请求未结束时为空, 示例值："xxxxx"
	CalendarID string  `path:"calendar_id" json:"-"` // 日历ID, 示例值："feishu.cn_xxxxxxxxxx@group.calendar.feishu.cn"
}

type getCalendarEventListResp struct {
	Code int                       `json:"code,omitempty"` // 错误码，非 0 表示失败
	Msg  string                    `json:"msg,omitempty"`  // 错误描述
	Data *GetCalendarEventListResp `json:"data,omitempty"` //
}

type GetCalendarEventListResp struct {
	HasMore   bool                            `json:"has_more,omitempty"`   // 是否有下一页数据
	PageToken string                          `json:"page_token,omitempty"` // 下次请求需要带上的分页标记，90 天有效期
	SyncToken string                          `json:"sync_token,omitempty"` // 下次请求需要带上的增量同步标记，90 天有效期
	Items     []*GetCalendarEventListRespItem `json:"items,omitempty"`      // 日程列表
}

type GetCalendarEventListRespItem struct {
	EventID          string                                  `json:"event_id,omitempty"`           // 日程ID
	Summary          string                                  `json:"summary,omitempty"`            // 日程标题, 最大长度：`1000` 字符
	Description      string                                  `json:"description,omitempty"`        // 日程描述, 最大长度：`8192` 字符
	StartTime        *GetCalendarEventListRespItemStartTime  `json:"start_time,omitempty"`         // 日程开始时间
	EndTime          *GetCalendarEventListRespItemEndTime    `json:"end_time,omitempty"`           // 日程结束时间
	Vchat            *GetCalendarEventListRespItemVchat      `json:"vchat,omitempty"`              // 视频会议信息，仅当日程至少有一位attendee时生效
	Visibility       string                                  `json:"visibility,omitempty"`         // 日程公开范围，新建日程默认为Default；仅新建日程时对所有参与人生效，之后修改该属性仅对当前身份生效, 可选值有: `default`：默认权限，仅向他人显示是否“忙碌”, `public`：公开，显示日程详情, `private`：私密，仅自己可见
	AttendeeAbility  string                                  `json:"attendee_ability,omitempty"`   // 参与人权限, 可选值有: `none`：无法编辑日程、无法邀请其它参与人、无法查看参与人列表, `can_see_others`：无法编辑日程、无法邀请其它参与人、可以查看参与人列表, `can_invite_others`：无法编辑日程、可以邀请其它参与人、可以查看参与人列表, `can_modify_event`：可以编辑日程、可以邀请其它参与人、可以查看参与人列表
	FreeBusyStatus   string                                  `json:"free_busy_status,omitempty"`   // 日程占用的忙闲状态，新建日程默认为Busy；仅新建日程时对所有参与人生效，之后修改该属性仅对当前身份生效, 可选值有: `busy`：忙碌, `free`：空闲
	Location         *GetCalendarEventListRespItemLocation   `json:"location,omitempty"`           // 日程地点
	Color            int                                     `json:"color,omitempty"`              // 日程颜色，颜色RGB值的int32表示。仅对当前身份生效；客户端展示时会映射到色板上最接近的一种颜色；值为0或-1时默认跟随日历颜色。
	Reminders        []*GetCalendarEventListRespItemReminder `json:"reminders,omitempty"`          // 日程提醒列表
	Recurrence       string                                  `json:"recurrence,omitempty"`         // 重复日程的重复性规则, 最大长度：`2000` 字符
	Status           string                                  `json:"status,omitempty"`             // 日程状态, 可选值有: `tentative`：未回应, `confirmed`：已确认, `cancelled`：日程已取消
	IsException      bool                                    `json:"is_exception,omitempty"`       // 日程是否是一个重复日程的例外日程
	RecurringEventID string                                  `json:"recurring_event_id,omitempty"` // 例外日程的原重复日程的event_id
	Schemas          []*GetCalendarEventListRespItemSchema   `json:"schemas,omitempty"`            // 日程自定义信息
}

type GetCalendarEventListRespItemStartTime struct {
	Date      string `json:"date,omitempty"`      // 仅全天日程使用该字段，如2018-09-01。需满足 RFC3339 格式。不能与 time_stamp 同时指定
	Timestamp string `json:"timestamp,omitempty"` // 秒级时间戳，如1602504000(表示2020/10/12 20:0:00 +8时区)
	Timezone  string `json:"timezone,omitempty"`  // 时区名称，使用IANA Time Zone Database标准，如Asia/Shanghai；全天日程时区固定为UTC，非全天日程时区默认为Asia/Shanghai
}

type GetCalendarEventListRespItemEndTime struct {
	Date      string `json:"date,omitempty"`      // 仅全天日程使用该字段，如2018-09-01。需满足 RFC3339 格式。不能与 time_stamp 同时指定
	Timestamp string `json:"timestamp,omitempty"` // 秒级时间戳，如1602504000(表示2020/10/12 20:0:00 +8时区)
	Timezone  string `json:"timezone,omitempty"`  // 时区名称，使用IANA Time Zone Database标准，如Asia/Shanghai；全天日程时区固定为UTC，非全天日程时区默认为Asia/Shanghai
}

type GetCalendarEventListRespItemVchat struct {
	MeetingUrl string `json:"meeting_url,omitempty"` // 视频会议URL, 长度范围：`1` ～ `2000` 字符
}

type GetCalendarEventListRespItemLocation struct {
	Name      string  `json:"name,omitempty"`      // 地点名称, 长度范围：`1` ～ `512` 字符
	Address   string  `json:"address,omitempty"`   // 地点地址, 长度范围：`1` ～ `255` 字符
	Latitude  float64 `json:"latitude,omitempty"`  // 地点坐标纬度信息，对于国内的地点，采用GCJ-02标准，海外地点采用WGS84标准
	Longitude float64 `json:"longitude,omitempty"` // 地点坐标经度信息，对于国内的地点，采用GCJ-02标准，海外地点采用WGS84标准
}

type GetCalendarEventListRespItemReminder struct {
	Minutes int `json:"minutes,omitempty"` // 日程提醒时间的偏移量，正数时表示在日程开始前X分钟提醒，负数时表示在日程开始后X分钟提醒,新建或更新日程时传入该字段，仅对当前身份生效, 取值范围：`-20160` ～ `20160`
}

type GetCalendarEventListRespItemSchema struct {
	UiName   string `json:"ui_name,omitempty"`   // UI项名称 TODO文档
	UiStatus string `json:"ui_status,omitempty"` // UI项自定义状态, 可选值有: `hide`：隐藏显示, `readonly`：只读, `editable`：可编辑, `unknown`：未知UI项自定义状态，仅用于读取时兼容
	AppLink  string `json:"app_link,omitempty"`  // 按钮点击后跳转的链接, 最大长度：`2000` 字符
}
