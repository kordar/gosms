package gosms

type SMSProvider interface {
	SendSingle(req SMSRequest) ([]SMSResult, error)      // 一对一或一对多
	SendMultiple(reqs []SMSRequest) ([]SMSResult, error) // 多对多
	SendTemplate(req SMSRequest) ([]SMSResult, error)    // 模板短信
	HandleReport(body []byte) ([]SMSReport, error)       // 状态报告回调
	HandleInbound(body []byte) ([]SMSInbound, error)     // 上行短信回调
}
