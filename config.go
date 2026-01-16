package gosms

type SMSConfig struct {
	Provider    string // "aliyun", "tencent"
	AccessKey   string
	SecretKey   string
	SignName    string            // 短信签名
	TemplateID  string            // 默认模板ID（可选）
	ExtraParams map[string]string // 扩展参数（可选）
}

// 构造函数
func NewSMSConfig(provider, accessKey, secretKey string) *SMSConfig {
	return &SMSConfig{
		Provider:    provider,
		AccessKey:   accessKey,
		SecretKey:   secretKey,
		ExtraParams: map[string]string{},
	}
}

// 设置签名
func (c *SMSConfig) WithSign(sign string) *SMSConfig {
	c.SignName = sign
	return c
}

// 设置默认模板
func (c *SMSConfig) WithTemplate(templateID string) *SMSConfig {
	c.TemplateID = templateID
	return c
}

// 设置额外参数
func (c *SMSConfig) WithExtraParam(key, value string) *SMSConfig {
	c.ExtraParams[key] = value
	return c
}
