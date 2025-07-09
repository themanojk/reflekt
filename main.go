package main

import (
	"context"
	"fmt"
	"time"

	"github.com/themanojk/reflekt/application"
	"github.com/themanojk/reflekt/pkg/config"
)

func main() {
	cfg, err := config.Load("configs/config.yaml")
	if err != nil {
		fmt.Println("failed to load config: %w", err)
	}

	app, err := application.New(cfg)
	if err != nil {
		fmt.Println("failed to start the server", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = app.Start(ctx, cfg)

	if err != nil {
		fmt.Println("failed to start the server", err)
	}
}
