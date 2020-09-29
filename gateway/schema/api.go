package schema

type Api struct {
	// 标题
	Title string `json:"title"`
	// 所属分组
	Gid int32 `json:"gid"`
	// 项目编号
	Pid int32 `json:"pid"`
	// 可以请求的方式
	RequestMethod []string `json:"request_method"`
	// 请求路径
	RequestURL string `json:"path"`
	// 转发协议
	Proto string `json:"proto"`
	// 上游名称，或者Target
	Upstreams string `json:"upstreams"`
	// 匹配的host
	Host string `json:"host"`
	// 转发方式
	TargetMethod string `json:"target_method"`
	// 转发路径
	TargetURL string `json:"target_url"`
	// 请求超时时间
	TimeOut int32 `json:"time_out"`
	// 请求重试次数
	RetryCount int32 `json:"retry_count"`
	// 错误返回
	ErrorResponse string `json:"error_response"`
}
