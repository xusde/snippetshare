package main

import (
	"flag"
	"log/slog"
	"net/http"
	"os"
)

type application struct {
	logger *slog.Logger
}

func main()  {
	// parse command line flags
	addr := flag.String("addr", ":4000", "HTTP network address")
	flag.Parse()

	// build custom logger
	logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		AddSource: true,
	}))
	app := &application{
		logger: logger,
	}
	
	// start server
	logger.Info("Starting server", "addr", *addr)
	err := http.ListenAndServe(*addr, app.routes())

	logger.Error(err.Error())
	os.Exit(1)
}