package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"example.com/pulppress/internal/config"
	"example.com/pulppress/internal/db"
	"example.com/pulppress/internal/policy"
	"example.com/pulppress/internal/store"
	"example.com/pulppress/internal/svc"
)

func main() {
	cfg := config.Config{DBPath: "pulppress.sqlite", DataDir: "data", AdminToken: "dev-token"}.Normalized()
	_ = os.MkdirAll(cfg.DataDir, 0o755)
	sqlDB, err := db.Open(cfg.DBPath)
	if err != nil {
		log.Fatal(err)
	}
	defer sqlDB.Close()
	cat := svc.NewCatalog(store.New(sqlDB), cfg)
	ctx := context.Background()
	if err := cat.BootstrapAdmin(ctx); err != nil {
		log.Fatal(err)
	}
	items := policy.Seed()
	for _, it := range items {
		if _, err := cat.Create(ctx, it.Title, it.Body, it.Tags, 1); err != nil {
			log.Fatal(err)
		}
	}
	fmt.Println("seeded", len(items))
}
