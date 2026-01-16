# 短信接口设计说明文档

## 1. 设计目标

* 对外提供统一短信接口，支持阿里云短信、腾讯短信等多个平台。
* 支持普通短信、一对一/一对多、多对多、模板短信。
* 支持状态报告与上行短信接收。
* 短信配置和发送请求均可通过 **链式 With 扩展参数**，统一风格，便于扩展。

---

## 2. 核心概念

* **SMSConfig**：短信平台基础配置，支持链式扩展参数。
* **SMSRequest**：短信发送请求，支持普通短信与模板短信，多对多发送。
* **SMSProvider**：短信平台适配器接口，阿里云、腾讯短信分别实现。
* **HTTP 接收服务**：接收短信状态报告和上行短信。

---

## 3. SMSConfig 设计

```go
type SMSConfig struct {
    Provider    string            // 平台：aliyun / tencent
    AccessKey   string
    SecretKey   string
    SignName    string            // 默认短信签名
    TemplateID  string            // 默认模板ID（可选）
    ExtraParams map[string]string // 扩展参数
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

// 链式方法
func (c *SMSConfig) WithSign(sign string) *SMSConfig {
    c.SignName = sign
    return c
}

func (c *SMSConfig) WithTemplate(templateID string) *SMSConfig {
    c.TemplateID = templateID
    return c
}

func (c *SMSConfig) WithExtraParam(key, value string) *SMSConfig {
    c.ExtraParams[key] = value
    return c
}
```

### 使用示例

```go
aliyunConfig := NewSMSConfig("aliyun", "AKID_EXAMPLE", "SECRET_EXAMPLE").
                    WithSign("公司签名").
                    WithTemplate("SMS_123456").
                    WithExtraParam("region", "cn-hangzhou").
                    WithExtraParam("timeout", "5s")
```

```go
tencentConfig := NewSMSConfig("tencent", "AKID_TENCENT", "SECRET_TENCENT").
                     WithSign("公司签名").
                     WithExtraParam("sdkAppID", "1400000000").
                     WithExtraParam("endpoint", "sms.tencentcloudapi.com")
```

---

## 4. SMSRequest 设计

```go
type SMSRequest struct {
    PhoneNumbers []string
    Content      string
    TemplateID   string
    TemplateVars map[string]string
    ExtraParams  map[string]string
}

// 构造函数
func NewSMSRequest(phones []string, content string) *SMSRequest {
    return &SMSRequest{
        PhoneNumbers: phones,
        Content:      content,
        TemplateVars: map[string]string{},
        ExtraParams:  map[string]string{},
    }
}

// 链式扩展
func (r *SMSRequest) WithTemplateID(tid string) *SMSRequest {
    r.TemplateID = tid
    return r
}

func (r *SMSRequest) WithTemplateVar(key, value string) *SMSRequest {
    r.TemplateVars[key] = value
    return r
}

func (r *SMSRequest) WithExtraParam(key, value string) *SMSRequest {
    r.ExtraParams[key] = value
    return r
}
```

### 使用示例

```go
req := NewSMSRequest([]string{"13800000000"}, "您的验证码是123456").
            WithExtraParam("outId", "20260116_001").
            WithExtraParam("expire", "5m")
```

---

## 5. SMSProvider 接口

```go
type SMSProvider interface {
    SendSingle(req SMSRequest) ([]SMSResult, error)       // 一对一/一对多
    SendMultiple(reqs []SMSRequest) ([]SMSResult, error) // 多对多
    SendTemplate(req SMSRequest) ([]SMSResult, error)    // 模板短信
    HandleReport(body []byte) ([]SMSReport, error)      // 状态报告
    HandleInbound(body []byte) ([]SMSInbound, error)    // 上行短信
}
```

---

## 6. 发送策略示例

### 一对一 / 一对多普通短信

```go
SendSMS(aliyun, NewSMSRequest([]string{"13800000000"}, "验证码123456"))
```

### 多对多短信

```go
reqs := []SMSRequest{
    {PhoneNumbers: []string{"13800000000"}, Content: "短信A"},
    {PhoneNumbers: []string{"13900000001"}, Content: "短信B"},
}
aliyun.SendMultiple(reqs)
```

### 模板短信

```go
req := NewSMSRequest([]string{"13800000000"}, "").
            WithTemplateID("SMS_123456").
            WithTemplateVar("code", "654321")
aliyun.SendTemplate(req)
```

---

## 7. 状态报告 & 上行短信接收

```go
r := gin.Default()

// 状态报告
r.POST("/sms/report", func(c *gin.Context) {
    body, _ := c.GetRawData()
    reports, _ := aliyun.HandleReport(body)
    fmt.Println("状态报告:", reports)
    c.String(200, "OK")
})

// 上行短信
r.POST("/sms/inbound", func(c *gin.Context) {
    body, _ := c.GetRawData()
    msgs, _ := aliyun.HandleInbound(body)
    fmt.Println("上行短信:", msgs)
    c.String(200, "OK")
})

r.Run(":8080")
```

---

## 8. 特点

* **统一接口**：调用方无需关心底层平台。
* **链式配置**：`SMSConfig` 与 `SMSRequest` 支持 `WithXXX` 扩展参数。
* **多种发送策略**：一对一、一对多、多对多、模板短信。
* **可扩展**：新增短信平台仅需实现 `SMSProvider` 接口。
* **接收支持**：状态报告和上行短信均可通过 HTTP 接收。

---

