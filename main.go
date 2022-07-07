package main

import (
	"fmt"
	"os"
	"os/signal"
	"strings"
	"syscall"

	_ "github.com/joho/godotenv/autoload"
	"github.com/lonelyevil/khl"
	"github.com/lonelyevil/khl/log_adapter/plog"
	"github.com/phuslu/log"
)

var (
	VERSION string = "0.1.0"
)

func main() {
	logger := log.Logger{
		Level:  log.TraceLevel,
		Writer: &log.ConsoleWriter{},
	}
	s := khl.New(os.Getenv("TOKEN"), plog.NewLogger(&logger))
	s.AddHandler(messageHandler)
	s.Open()
	// Wait here until CTRL-C or other term signal is received.
	fmt.Println("Bot is now running.  Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, os.Interrupt, syscall.SIGTERM)
	<-sc

	// Cleanly close down the KHL session.
	s.Close()
}

func messageHandler(ctx *khl.KmarkdownMessageContext) {
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
				Content:  fmt.Sprintf("Version: %s", VERSION),
				Quote:    ctx.Common.MsgID,
				Type:     khl.MessageTypeKMarkdown,
			},
		})
	}
}
