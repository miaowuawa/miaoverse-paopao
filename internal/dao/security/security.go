package security

import (
	"strings"

	"github.com/alimy/tryst/cfg"
	"github.com/rocboss/paopao-ce/internal/core"
)

func NewPhoneVerifyService() core.PhoneVerifyService {
	smsVendor, _ := cfg.Val("Sms")
	switch strings.ToLower(smsVendor) {
	case "smsjuhe":
		return newJuheSmsServant()
	case "smsbao":
		return newSmsBaoServant()
	default:
		return newSmsBaoServant()
	}
}
