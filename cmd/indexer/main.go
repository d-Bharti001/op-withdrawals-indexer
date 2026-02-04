package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	indexer "op-withdrawals-indexer/internal/indexer/services"

	"github.com/urfave/cli/v3"
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	cmd := &cli.Command{
		Name:   "OP Withdrawals Indexer",
		Usage:  "index op stack l2 chain's withdrawals",
		Flags:  appFlags,
		Action: appAction,
	}

	if err := cmd.Run(ctx, os.Args); err != nil {
		log.Println(err)
	}
}

func appAction(ctx context.Context, cmd *cli.Command) error {
	app, err := indexer.NewIndexer(ctx, appConfig)
	if err != nil {
		return err
	}

	defer app.Stop()

	if err := app.Start(); err != nil {
		return err
	}

	<-ctx.Done()

	return nil
}
