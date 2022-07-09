package handlers

import (
	"fmt"
	"strings"

	"github.com/lonelyevil/khl"
	"github.com/shuyangzhang/shiori/constant"
	"github.com/shuyangzhang/shiori/tools"
)

func commandHandler(ctx *khl.KmarkdownMessageContext) {
	if ctx.Common.Type != khl.MessageTypeKMarkdown || ctx.Extra.Author.Bot {
		return
	}

	withPrefix, command, params := tools.GetCommandWithParameters(ctx.Common.Content)

	if withPrefix {
		tools.CommandRouter(ctx, command, params)
	}

	if withPrefix {
		ctx.Session.MessageCreate(&khl.MessageCreate{
			MessageCreateBase: khl.MessageCreateBase{
				TargetID: ctx.Common.TargetID,
				Content:  fmt.Sprintf("your command is %s, your args are %s", command, params),
				Quote:    ctx.Common.MsgID,
				Type:     khl.MessageTypeKMarkdown,
			},
		})
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
