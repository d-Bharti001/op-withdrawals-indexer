package api

import "time"

const ShutdownTimeout = 40 * time.Second

type APIInitConfig struct {
	HTTPPort string

	DBConnStr     string
	DBMaxOpenConn int
	DBMaxIdleConn int
	DBMaxIdleTime string
}
