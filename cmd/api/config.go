package main

import (
	api "op-withdrawals-indexer/internal/api/service"

	"github.com/urfave/cli/v3"
)

var appConfig api.APIInitConfig

var appFlags []cli.Flag = []cli.Flag{
	&cli.StringFlag{
		Name:        "httpPort",
		Usage:       "HTTP Server Port (e.g. 8080)",
		Required:    true,
		Sources:     cli.EnvVars("HTTP_PORT"),
		Destination: &appConfig.HTTPPort,
	},
	&cli.StringFlag{
		Name:        "dbConnString",
		Usage:       "Database connection string",
		Required:    true,
		Sources:     cli.EnvVars("DB_CONN_STR"),
		Destination: &appConfig.DBConnStr,
	},
	&cli.IntFlag{
		Name:        "dbMaxOpenConn",
		Usage:       "Database max open connections",
		Sources:     cli.EnvVars("DB_MAX_OPEN_CONN"),
		Value:       30,
		DefaultText: "30",
		Destination: &appConfig.DBMaxOpenConn,
	},
	&cli.IntFlag{
		Name:        "dbMaxIdleConn",
		Usage:       "Database max idle connections",
		Sources:     cli.EnvVars("DB_MAX_IDLE_CONN"),
		Value:       30,
		DefaultText: "30",
		Destination: &appConfig.DBMaxIdleConn,
	},
	&cli.StringFlag{
		Name:        "dbMaxIdleTime",
		Usage:       "Database max idle time",
		Sources:     cli.EnvVars("DB_MAX_IDLE_TIME"),
		Value:       "15m",
		DefaultText: "15m",
		Destination: &appConfig.DBMaxIdleTime,
	},
}
