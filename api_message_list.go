package lark

import (
	"context"
)

// GetMessageList 获取会话（包括单聊、群组）的历史消息。
//
// 注意事项:
// - 需要开启[机器人能力](https://open.feishu.cn/document/uQjL04CN/uYTMuYTMuYTM)
// - 获取群组消息时，机器人必须在群组中
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/im-v1/message/list
func (r *MessageAPI) GetMessageList(ctx context.Context, request *GetMessageListReq) (*GetMessageListResp, *Response, error) {
	req := &requestParam{
		Method:                "GET",
		URL:                   "https://open.feishu.cn/open-apis/im/v1/messages",
		Body:                  request,
		NeedTenantAccessToken: true,
		NeedAppAccessToken:    false,
		NeedHelpdeskAuth:      false,
		IsFile:                false,
	}
	resp := new(getMessageListResp)

	response, err := r.cli.request(ctx, req, resp)
	if err != nil {
		return nil, response, err
	} else if resp.Code != 0 {
		return nil, response, newError("Message", "GetMessageList", resp.Code, resp.Msg)
	}

	return resp.Data, response, nil
}

type GetMessageListReq struct {
	ContainerIDType ContainerIDType `query:"container_id_type" json:"-"` // 容器类型 ，目前可选值仅有"chat", 示例值："chat"
	ContainerID     string          `query:"container_id" json:"-"`      // 容器的id，即chat的id, 示例值："oc_234jsi43d3ssi993d43545f"
	StartTime       *string         `query:"start_time" json:"-"`        // 历史信息的起始时间, 示例值："1609296809"
	EndTime         *string         `query:"end_time" json:"-"`          // 历史信息的结束时间, 示例值："1608594809"
	PageToken       *string         `query:"page_token" json:"-"`        // 分页标记，第一次请求不填，表示从头开始遍历；分页查询结果还有更多项时会同时返回新的 page_token，下次遍历可采用该 page_token 获取查询结果, 示例值："GxmvlNRvP0NdQZpa7yIqf_Lv_QuBwTQ8tXkX7w-irAghVD_TvuYd1aoJ1LQph86O-XImC4X9j9FhUPhXQDvtrQ=="
	PageSize        *int            `query:"page_size" json:"-"`         // 分页大小, 示例值：10, 最大值：`50`
}

type getMessageListResp struct {
	Code int                 `json:"code,omitempty"` // 错误码，非 0 表示失败
	Msg  string              `json:"msg,omitempty"`  // 错误描述
	Data *GetMessageListResp `json:"data,omitempty"` //
}

type GetMessageListResp struct {
	HasMore   bool                      `json:"has_more,omitempty"`   // 是否还有更多项
	PageToken string                    `json:"page_token,omitempty"` // 分页标记，当 has_more 为 true 时，会同时返回新的 page_token，否则不返回 page_token
	Items     []*GetMessageListRespItem `json:"items,omitempty"`      // message[]
}

type GetMessageListRespItem struct {
	MessageID      string       `json:"message_id,omitempty"`       // 消息id open_message_id
	RootID         string       `json:"root_id,omitempty"`          // 根消息id open_message_id
	ParentID       string       `json:"parent_id,omitempty"`        // 父消息的id open_message_id
	MsgType        MsgType      `json:"msg_type,omitempty"`         // 消息类型 text post card image等等
	CreateTime     string       `json:"create_time,omitempty"`      // 消息生成的时间戳(毫秒)
	UpdateTime     string       `json:"update_time,omitempty"`      // 消息更新的时间戳
	Deleted        bool         `json:"deleted,omitempty"`          // 消息是否被撤回
	Updated        bool         `json:"updated,omitempty"`          // 消息是否被更新
	ChatID         string       `json:"chat_id,omitempty"`          // 所属的群
	Sender         *Sender      `json:"sender,omitempty"`           // 发送者，可以是用户或应用
	Body           *MessageBody `json:"body,omitempty"`             // 消息内容，json结构，格式说明参考： [消息content说明](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/im-v1/message/events/message_content)
	Mentions       []*Mention   `json:"mentions,omitempty"`         // 被艾特的人或应用的id
	UpperMessageID string       `json:"upper_message_id,omitempty"` // 上一层级的消息id open_message_id
}
