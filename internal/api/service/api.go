package api

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"

	"op-withdrawals-indexer/internal/api/dbstore"
	"op-withdrawals-indexer/internal/api/docs"
	"op-withdrawals-indexer/internal/api/handlers/withdrawals"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	httpSwagger "github.com/swaggo/http-swagger/v2"
)

type APIService struct {
	ctx      context.Context
	cancel   context.CancelFunc
	stopOnce sync.Once

	dbProvider dbstore.DBStoreProvider

	server *http.Server
	port   string
}

func NewAPIService(ctx context.Context, cfg APIInitConfig) (*APIService, error) {
	ctx, cancel := context.WithCancel(ctx)

	// Initialize database
	db, err := dbstore.NewPostgresStore(
		cfg.DBConnStr,
		cfg.DBMaxOpenConn,
		cfg.DBMaxIdleConn,
		cfg.DBMaxIdleTime,
	)
	if err != nil {
		cancel()
		return nil, err
	}

	// Mount router

	r := chi.NewRouter()

	// Basic CORS
	// for more ideas, see: https://developer.github.com/v3/#cross-origin-resource-sharing
	r.Use(cors.Handler(cors.Options{
		// AllowedOrigins:   []string{"https://foo.com"}, // Use this to allow specific origin hosts
		AllowedOrigins: []string{"https://*", "http://*"},
		// AllowOriginFunc:  func(r *http.Request, origin string) bool { return true },
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))

	// A good base middleware stack
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	// Set a timeout value on the request context (ctx), that will signal
	// through ctx.Done() that the request has timed out and further
	// processing should be stopped.
	r.Use(middleware.Timeout(60 * time.Second))

	// Swagger docs
	r.Get("/swagger/*", httpSwagger.Handler(
		httpSwagger.URL(fmt.Sprintf("http://localhost:%s/swagger/doc.json", cfg.HTTPPort)),
	))

	withdrawalsHandler := withdrawals.NewWithdrawalsHandler(db)

	// Withdrawal history
	r.Get("/withdrawal-history", withdrawalsHandler.GetWithdrawalHistory)

	return &APIService{
		ctx:        ctx,
		cancel:     cancel,
		dbProvider: db,
		server: &http.Server{
			Addr:         fmt.Sprintf(":%s", cfg.HTTPPort),
			Handler:      r,
			ReadTimeout:  10 * time.Second,
			WriteTimeout: 30 * time.Second,
			IdleTimeout:  time.Minute,
		},
	}, nil
}

func (app *APIService) Start() error {

	// Swagger docs configuration
	docs.SwaggerInfo.Host = fmt.Sprintf("localhost:%s", app.port)
	docs.SwaggerInfo.BasePath = "/"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}

	log.Println("Starting HTTP server on port", app.port)

	return app.server.ListenAndServe()
}

func (app *APIService) Stop() {
	app.stopOnce.Do(func() {
		log.Println("Stopping API Service...")

		ctx, cancel := context.WithTimeout(context.Background(), ShutdownTimeout)
		defer cancel()

		if app.server != nil {
			err := app.server.Shutdown(ctx)
			if err != nil {
				log.Println("Error while shutting down http server :::", err)
			} else {
				log.Println("HTTP server shut down.")
			}
		}

		if app.dbProvider != nil {
			err := app.dbProvider.CloseConnection()
			if err != nil {
				log.Println("Error while closing database connection :::", err)
			} else {
				log.Println("Database connection closed.")
			}
		}

		log.Println("API Service stopped.")
	})
}
