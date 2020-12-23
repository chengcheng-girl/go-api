package qo

// Persistant Object 持久化对象
// Query  Object  查询对象
// View  Object  视图层对象
type Resp struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}
type UDNBody struct {
	ID    string `json:"id"`
	Items struct {
		NodeName     string `json:"node_name"`
		ApproveUser  string `json:"approve_user"`
		ApproveAdmin int8   `json:"approve_admin"`
	} `json:items`
}
type StInBody struct {
	AppId    int    `json:"app_id"`
	AppName  string `json:"app_name"`
	NsId     int    `json:"ns_id"`
	NsName   string `json:"ns_name"`
	Cluster  string `json:"cluster"`
	Data     string `json:"data"`
	DefineId int    `json:"define_id"`
	Remark   string `json:"remark"`
}