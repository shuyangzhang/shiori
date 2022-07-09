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
	"github.com/shuyangzhang/shiori/handlers"
)

func main() {
	logger := log.Logger{
		Level:  log.TraceLevel,
		Writer: &log.ConsoleWriter{},
	}
	s := khl.New(os.Getenv("TOKEN"), plog.NewLogger(&logger))
	handlers.RegisterHandlers(s)
	s.Open()
	// Wait here until CTRL-C or other term signal is received.
	fmt.Println("Bot is now running.  Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, os.Interrupt, syscall.SIGTERM)
	<-sc

	// Cleanly close down the KHL session.
	s.Close()
}
