package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	_ "github.com/joho/godotenv/autoload"
	"github.com/lonelyevil/khl"
	"github.com/lonelyevil/khl/log_adapter/plog"
	"github.com/phuslu/log"
	"github.com/shuyangzhang/shiori/configs"
	"github.com/shuyangzhang/shiori/handlers"
)

func main() {
	logger := log.Logger{
		Level:  log.TraceLevel,
		Writer: &log.ConsoleWriter{},
	}

	configs.InitEnvConfigs()

	s := khl.New(configs.EnvConfigs.Token, plog.NewLogger(&logger))

	handlers.RegisterHandlers(s)

	s.Open()
	fmt.Println("Bot is now running.  Press CTRL-C to exit.")

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, os.Interrupt, syscall.SIGTERM)
	<-sc
	s.Close()
}
