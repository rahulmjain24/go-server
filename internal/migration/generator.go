package migration

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func CreateMigration(name string) error {
	timestamp := time.Now().Format("20060102150405")
	safeName := strings.ToLower(strings.ReplaceAll(name, " ", "_"))
	base := fmt.Sprintf("%s_%s", timestamp, safeName)
	migrationsDir := "db/migrations"

	if err := os.MkdirAll(migrationsDir, os.ModePerm); err != nil {
		return fmt.Errorf("failed to create migrations folder: %v", err)
	}

	upFile := filepath.Join(migrationsDir, base+".up.sql")
	downFile := filepath.Join(migrationsDir, base+".down.sql")

	for _, file := range []string{upFile, downFile} {
		f, err := os.Create(file)
		if err != nil {
			return fmt.Errorf("failed to create file %s: %v", file, err)
		}
		f.Close()
	}

	fmt.Println("Migration files created:")
	fmt.Println("  ↳", upFile)
	fmt.Println("  ↳", downFile)
	return nil
}
