package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	api "op-withdrawals-indexer/internal/api/service"

	"github.com/urfave/cli/v3"
)

// @title			OP Withdrawals Indexer API
// @version		1.0
// @description	Frontend API service for OP withdrawals indexer
// @BasePath		/
func main() {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	cmd := &cli.Command{
		Name:   "OP Withdrawals Indexer API",
		Usage:  "api for viewing saved withdrawals",
		Flags:  appFlags,
		Action: appAction,
	}

	if err := cmd.Run(ctx, os.Args); err != nil {
		log.Println(err)
	}
}

func appAction(ctx context.Context, cmd *cli.Command) error {
	app, err := api.NewAPIService(ctx, appConfig)
	if err != nil {
		return err
	}

	defer app.Stop()

	err = app.Start()
	if err != nil {
		return err
	}

	<-ctx.Done()

	return nil
}
