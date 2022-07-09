package handlers

import (
	"github.com/lonelyevil/khl"
	"github.com/shuyangzhang/shiori/tools/parser"
)

func commandHandler(ctx *khl.KmarkdownMessageContext) {
	if ctx.Common.Type != khl.MessageTypeKMarkdown || ctx.Extra.Author.Bot {
		return
	}

	withPrefix, command, params := parser.GetCommandWithParameters(ctx.Common.Content)

	if withPrefix {
		parser.RouteCommand(ctx, command, params)
	}
}
