package handlers

import (
	"fmt"
	"strings"

	"github.com/lonelyevil/khl"
	"github.com/shuyangzhang/shiori/constant"
)

func commandHandler(ctx *khl.KmarkdownMessageContext) {
	if ctx.Common.Type != khl.MessageTypeKMarkdown || ctx.Extra.Author.Bot {
		return
	}

	if strings.HasPrefix(ctx.Common.Content, ",ping") {
		ctx.Session.MessageCreate(&khl.MessageCreate{
			MessageCreateBase: khl.MessageCreateBase{
				TargetID: ctx.Common.TargetID,
				Content:  "エンチャントアロー!",
				Quote:    ctx.Common.MsgID,
				Type:     khl.MessageTypeKMarkdown,
			},
		})
	}

	if strings.HasPrefix(ctx.Common.Content, ",version") {
		ctx.Session.MessageCreate(&khl.MessageCreate{
			MessageCreateBase: khl.MessageCreateBase{
				TargetID: ctx.Common.TargetID,
				Content:  fmt.Sprintf("Version: %s", constant.Version),
				Quote:    ctx.Common.MsgID,
				Type:     khl.MessageTypeKMarkdown,
			},
		})
	}
}
