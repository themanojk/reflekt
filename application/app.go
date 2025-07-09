package application

import (
	"context"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/themanojk/reflekt/pkg/config"
	"github.com/themanojk/reflekt/store"
)

type App struct {
	router http.Handler
	store  store.Store
}

func New(cfg *config.Config) (*App, error) {
	// 1) Connect Mongo
	client, err := store.NewMongoClient(context.Background(), cfg.MongoURI)
	if err != nil {
		return nil, fmt.Errorf("mongo connect: %w", err)
	}

	// 2) Build your store
	mongoStore := store.NewMongoStore(client, cfg.MongoDB)

	// 3) Wire it up
	app := &App{
		store:  mongoStore,
		router: chi.NewRouter(),
	}

	app.loadRoutes()
	return app, nil
}

func (a *App) Start(ctx context.Context, cfg *config.Config) error {
	server := &http.Server{
		Addr:    cfg.HTTPAddr,
		Handler: a.router,
	}

	client, err := store.NewMongoClient(ctx, cfg.MongoURI)
	if err != nil {
		fmt.Println("mongo connect: %w", err)
	}
	defer func() {
		_ = client.Disconnect(ctx)
	}()

	println("Staring server")

	err = server.ListenAndServe()
	if err != nil {
		println("failed to start the server", err)
	}

	return err
}
