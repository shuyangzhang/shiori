package router

import (
	"github.com/lonelyevil/khl"
	"github.com/shuyangzhang/shiori/internal/service"
)

var CommandRouter = make(map[string]func(*khl.KmarkdownMessageContext, ...string))

func init() {
	registerCommand(service.Ping, "ping")
	registerCommand(service.Version, "version")
	registerCommand(service.NameToAliases, "alias", "别名", "外号")
	registerCommand(service.Arena, "jjc", "homework", "竞技场", "作业", "击剑")
}

func registerCommand(service func(*khl.KmarkdownMessageContext, ...string), commandsAndAliases ...string) {
	for _, command := range commandsAndAliases {
		CommandRouter[command] = service
	}
}
