package gosms

import "time"

// 短信发送结果
type SMSResult struct {
	PhoneNumber string
	Success     bool
	Code        string
	Message     string
	RequestID   string
}

// 状态报告
type SMSReport struct {
	PhoneNumber string
	Status      string // DELIVERED, FAILED, UNKNOWN
	MsgID       string
	Timestamp   time.Time
}

// 上行短信
type SMSInbound struct {
	PhoneNumber string
	Content     string
	MsgID       string
	Timestamp   time.Time
}
