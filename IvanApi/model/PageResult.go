package model

// 响应数据
type PageResult struct {
	Total     int         `json:"total"`     // 总条数
	PageSize  int         `json:"pageSize"`  // 页码
	PageIndex int         `json:"pageIndex"` // 页签
	Result    interface{} `json:"result"`    // 结果
}

type ResponseResult struct {
	Code   int         `json:"code"`   // 状态
	Msg    string      `json:"msg"`    // 消息
	Result interface{} `json:"result"` // 结果
}
