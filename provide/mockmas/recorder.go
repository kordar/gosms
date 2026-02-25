package mockmas

import "github.com/kordar/gosms"

func (p *Provider) LastSingle() *gosms.SMSRequest {
	if len(p.Singles) == 0 {
		return nil
	}
	return &p.Singles[len(p.Singles)-1]
}
