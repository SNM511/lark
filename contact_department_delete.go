package lark

import (
	"context"
)

// DeleteDepartment 该接口用于向通讯录中删除部门。
//
// 应用需要同时拥有待删除部门及其父部门的通讯录授权。应用商店应用无权限调用该接口。
// https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/contact-v3/department/delete
func (r *ContactAPI) DeleteDepartment(ctx context.Context, request *DeleteDepartmentReq) (*DeleteDepartmentResp, *Response, error) {
	req := &requestParam{
		Method:                "DELETE",
		URL:                   "https://open.feishu.cn/open-apis/contact/v3/departments/:department_id",
		Body:                  request,
		NeedTenantAccessToken: true,
		NeedAppAccessToken:    false,
	}
	resp := new(deleteDepartmentResp)

	response, err := r.cli.request(ctx, req, resp)
	if err != nil {
		return nil, response, err
	} else if resp.Code != 0 {
		return nil, response, newError("Contact", "DeleteDepartment", resp.Code, resp.Msg)
	}

	return resp.Data, response, nil
}

type DeleteDepartmentReq struct {
	UserIDType       *IDType `query:"user_id_type" json:"-"`       // 用户 ID 类型,**示例值**："open_id",**可选值有**：,- `open_id`：用户的 open id,- `union_id`：用户的 union id,- `user_id`：用户的 user id,**默认值**：`open_id`,**当值为 `user_id`，字段权限要求**：,<md-perm href="/ssl:ttdoc/ukTMukTMukTM/uQjN3QjL0YzN04CN2cDN">获取用户 userid</md-perm>
	DepartmentIDType *IDType `query:"department_id_type" json:"-"` // 此次调用中使用的部门ID的类型,**示例值**："open_department_id",**可选值有**：,- `department_id`：以自定义department_id来标识部门,- `open_department_id`：以open_department_id来标识部门,**默认值**：`open_department_id`
	DepartmentID     string  `path:"department_id" json:"-"`       // 部门ID，需要与查询参数中传入的department_id_type类型保持一致。,**示例值**："od-4e6ac4d14bcd5071a37a39de902c7141",**数据校验规则**：,- 最大长度：`128` 字符,- 正则校验：`^0|[^od][A-Za-z0-9]*`
}

type deleteDepartmentResp struct {
	Code int                   `json:"code,omitempty"` // 错误码，非 0 表示失败
	Msg  string                `json:"msg,omitempty"`  // 错误描述
	Data *DeleteDepartmentResp `json:"data,omitempty"`
}

type DeleteDepartmentResp struct{}
