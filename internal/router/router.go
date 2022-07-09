package router

import (
	"github.com/lonelyevil/khl"
	"github.com/shuyangzhang/shiori/internal/service"
)

var CommandRouter = make(map[string]func(*khl.KmarkdownMessageContext, ...string))

func init() {
	CommandRouter["ping"] = service.Ping
	CommandRouter["version"] = service.Version
}
