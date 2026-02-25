package gosms

type ErrorCode string

const (
	// 通用
	ErrSuccess         ErrorCode = "SUCCESS"
	ErrUnknown         ErrorCode = "UNKNOWN"
	ErrInvalidRequest  ErrorCode = "INVALID_REQUEST"
	ErrAuthFailed      ErrorCode = "AUTH_FAILED"
	ErrSignInvalid     ErrorCode = "SIGN_INVALID"
	ErrTemplateInvalid ErrorCode = "TEMPLATE_INVALID"
	ErrContentInvalid  ErrorCode = "CONTENT_INVALID"
	ErrTooManyMobiles  ErrorCode = "TOO_MANY_MOBILES"

	// 网络 / 系统
	ErrNetwork ErrorCode = "NETWORK_ERROR"
	ErrTimeout ErrorCode = "TIMEOUT"
)

type SMSError struct {
	Code     ErrorCode // 统一错误码
	Provider string    // mas / aliyun / tencent
	RawCode  string    // 原始错误码
	Message  string    // 可读信息
}

func (e *SMSError) Error() string {
	return string(e.Code) + ": " + e.Message
}
