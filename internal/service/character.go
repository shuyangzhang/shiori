package service

import (
	"fmt"
	"strings"

	"github.com/lonelyevil/khl"
	"github.com/shuyangzhang/shiori/constant"
	"github.com/shuyangzhang/shiori/tools/message"
)

func NameToAliases(ctx *khl.KmarkdownMessageContext, parameters ...string) {
	var alias string
	if len(parameters) != 1 {
		panic("missing parameters or too many parameters received")
	} else {
		alias = parameters[0]
	}
	characterId, ok := constant.CharactersReverseMap[alias]
	var aliases []string
	if ok {
		aliases, _ = constant.Characters[characterId]
	} else {
		panic(fmt.Sprintf("can not find %v, please check your input or contanct developer via @manako", alias))
	}
	result := strings.Join(aliases, ",")

	message.Reply(ctx, result)
}
