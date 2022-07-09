package service

import (
	"fmt"

	"github.com/lonelyevil/khl"
	"github.com/shuyangzhang/shiori/constant"
	"github.com/shuyangzhang/shiori/tools/message"
)

func Ping(ctx *khl.KmarkdownMessageContext, parameters ...string) {
	message.Send(ctx, "エンチャントアロー!")
}

func Version(ctx *khl.KmarkdownMessageContext, parameters ...string) {
	message.Reply(ctx, fmt.Sprintf("Version: %v", constant.Version))
}
