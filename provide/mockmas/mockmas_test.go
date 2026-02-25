package mockmas_test

import (
	"github.com/kordar/gosms"
	"github.com/kordar/gosms/provide/mockmas"
	"testing"
)

func TestSendSMS(t *testing.T) {
	cfg := gosms.NewSMSConfig("mas", "ak", "sk")
	p, _ := mockmas.New(cfg)

	mock := p.(*mockmas.Provider)

	req := gosms.NewSMSRequest(
		[]string{"13800138000"},
		"验证码1234",
	)

	res, err := p.SendSingle(*req)
	if err != nil {
		t.Fatal(err)
	}

	if !res[0].Success {
		t.Fatal("send failed")
	}

	if mock.LastSingle().Content != "验证码1234" {
		t.Fatal("content mismatch")
	}
}
