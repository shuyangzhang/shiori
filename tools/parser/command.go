package parser

import (
	"context"
	"fmt"
	"strings"

	"github.com/google/uuid"
	"github.com/lonelyevil/khl"
	"github.com/shuyangzhang/shiori/configs"
	"github.com/shuyangzhang/shiori/internal/router"
	log "github.com/sirupsen/logrus"
	"golang.org/x/exp/slices"
)

func GetCommandWithParameters(rawCommand string) (withPrefix bool, command string, params []string) {
	runeRawCommand := []rune(rawCommand)
	prefix := string(runeRawCommand[0])
	commandWithParameters := string(runeRawCommand[1:])

	if slices.Contains(configs.EnvConfigs.AllPrefixes, prefix) {
		commandSlice := strings.Fields(commandWithParameters)

		withPrefix = true
		command = commandSlice[0]
		params = commandSlice[1:]
	}

	return
}

func RouteCommand(ctx *khl.KmarkdownMessageContext, command string, params []string) {
	logId := uuid.NewString()

	commonCtx := context.Background()
	commonCtx = context.WithValue(commonCtx, "logId", logId)

	teardown := commandLogger(commonCtx, ctx, command, params)
	defer teardown(commonCtx)

	commandService, ok := router.CommandRouter[command]

	if ok {
		commandService(ctx, params...)
	} else {
		log.Warn(fmt.Sprintf("unknown command: %v", command))
	}

	if command == "testpanic" {
		panic("i am testing panic")
	}
}

func commandLogger(commonCtx context.Context, ctx *khl.KmarkdownMessageContext, command string, params []string) func(commonCtx context.Context) {
	log.WithFields(log.Fields{
		"logId":       commonCtx.Value("logId"),
		"channelName": ctx.Extra.ChannelName,
		"guildId":     ctx.Extra.GuildID,
	}).Info(
		fmt.Sprintf("user: %v, username: %v, used command: %v, with args: %v",
			ctx.Common.AuthorID,
			ctx.Extra.Author.Nickname,
			command,
			params))

	return func(commonCtx context.Context) {
		if err := recover(); err != nil {
			log.WithField("logId", commonCtx.Value("logId")).Error(
				fmt.Sprintf("panic occurred: %v", err),
			)
		}
	}
}

func init() {
	log.SetFormatter(&log.TextFormatter{
		FullTimestamp: true,
	})
}
