package mockmas

import (
	"github.com/kordar/gosms"
	"sync"
	"time"
)

type Provider struct {
	mu sync.Mutex

	// 行为控制
	ForceError error
	ForceFail  bool
	NextMsgID  string

	// 记录请求
	Singles   []gosms.SMSRequest
	Multiples [][]gosms.SMSRequest
	Templates []gosms.SMSRequest
}

func New(cfg *gosms.SMSConfig) (gosms.SMSProvider, error) {
	return &Provider{
		NextMsgID: "mock-msg-001",
	}, nil
}

func (p *Provider) SendSingle(req gosms.SMSRequest) ([]gosms.SMSResult, error) {
	p.mu.Lock()
	defer p.mu.Unlock()

	p.Singles = append(p.Singles, req)

	if p.ForceError != nil {
		return nil, p.ForceError
	}

	if p.ForceFail {
		return []gosms.SMSResult{{
			Success: false,
			Code:    "MOCK_FAILED",
			Message: "mock send failed",
		}}, nil
	}

	return []gosms.SMSResult{{
		Success: true,
		Code:    "success",
		Message: p.NextMsgID,
	}}, nil
}

func (p *Provider) SendMultiple(reqs []gosms.SMSRequest) ([]gosms.SMSResult, error) {
	p.mu.Lock()
	defer p.mu.Unlock()

	p.Multiples = append(p.Multiples, reqs)

	if p.ForceError != nil {
		return nil, p.ForceError
	}

	var res []gosms.SMSResult
	for range reqs {
		res = append(res, gosms.SMSResult{
			Success: !p.ForceFail,
			Code:    "success",
			Message: p.NextMsgID,
		})
	}
	return res, nil
}

func (p *Provider) SendTemplate(req gosms.SMSRequest) ([]gosms.SMSResult, error) {
	p.mu.Lock()
	defer p.mu.Unlock()

	p.Templates = append(p.Templates, req)

	if p.ForceError != nil {
		return nil, p.ForceError
	}

	return []gosms.SMSResult{{
		Success: !p.ForceFail,
		Code:    "success",
		Message: p.NextMsgID,
	}}, nil
}

func (p *Provider) HandleReport(body []byte) ([]gosms.SMSReport, error) {
	return []gosms.SMSReport{{
		PhoneNumber: "13800000000",
		Status:      "DELIVERED",
		MsgID:       p.NextMsgID,
		Timestamp:   time.Now(),
	}}, nil
}

func (p *Provider) HandleInbound(body []byte) ([]gosms.SMSInbound, error) {
	return []gosms.SMSInbound{{
		PhoneNumber: "13800000000",
		Content:     "MOCK INBOUND",
		MsgID:       "mock-inbound-001",
		Timestamp:   time.Now(),
	}}, nil
}
