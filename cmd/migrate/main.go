package main

import (
	"fmt"
	"os"

	"github.com/rahulmjain24/go-server/internal/migration"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: create <migration_name>")
		fmt.Println("Usage: run")
		return
	}

	switch command := os.Args[1]; command {
	case "create":
		name := os.Args[2]
		if err := migration.CreateMigration(name); err != nil {
			fmt.Println("Migration creation failed:", err)
			os.Exit(1)
		}
	case "run":
		if err := migration.RunMigrations(); err != nil {
			fmt.Println("Migration failed:", err)
			os.Exit(1)
		}
	default:
		break
	}
}
