package main

import (
	"context"
	"log"
	indexer "op-withdrawals-indexer/internal/indexer/services"
	"os"

	"github.com/urfave/cli/v3"
)

func main() {
	cmd := &cli.Command{
		Name:   "OP Withdrawals Indexer",
		Usage:  "index op stack l2 chain's withdrawals",
		Flags:  appFlags,
		Action: appAction,
	}

	if err := cmd.Run(context.Background(), os.Args); err != nil {
		log.Fatalln(err)
	}
}

func appAction(ctx context.Context, cmd *cli.Command) error {
	app, err := indexer.NewIndexer(ctx, appConfig)
	if err != nil {
		return err
	}
	return app.Start()
}
