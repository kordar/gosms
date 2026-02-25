package gosms

type SMSRequest struct {
	PhoneNumbers   []string
	Content        string            // 普通短信内容
	TemplateID     string            // 模板短信ID
	TemplateVars   map[string]string // 模板变量 (键值对)
	TemplateParams []string          // 模板参数 (有序列表，用于位置参数)
	ExtraParams    map[string]string // 可选扩展参数
}

func (r *SMSRequest) WithTemplateVars(vars map[string]string) *SMSRequest {
	r.TemplateVars = vars
	return r
}

// 构造函数
func NewSMSRequest(phones []string, content string) *SMSRequest {
	return &SMSRequest{
		PhoneNumbers:   phones,
		Content:        content,
		TemplateVars:   map[string]string{},
		TemplateParams: []string{},
		ExtraParams:    map[string]string{},
	}
}

// 链式扩展参数
func (r *SMSRequest) WithTemplateID(tid string) *SMSRequest {
	r.TemplateID = tid
	return r
}

func (r *SMSRequest) WithTemplateVar(key, value string) *SMSRequest {
	r.TemplateVars[key] = value
	return r
}

func (r *SMSRequest) WithTemplateParams(params []string) *SMSRequest {
	r.TemplateParams = params
	return r
}

func (r *SMSRequest) WithExtraParam(key, value string) *SMSRequest {
	r.ExtraParams[key] = value
	return r
}
