package cmd

import (
	"UNSWIoT/internal/models/sqlite"
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"github.com/spf13/cobra"
	"log"
)

var (
	rootCmd = &cobra.Command{
		Use:   "noRun",
		Short: "IoT security device management CLI",
	}

	queries *Queries
)

type Queries struct {
	Devices *sqlite.DeviceModel
}

func init() {
	db, err := sql.Open("sqlite3", "./backend/src/internal/db/IoT.db")

	if err != nil {
		log.Fatal(err)
	}
	queries = &Queries{
		Devices: &sqlite.DeviceModel{
			DB: db,
		},
	}
}

func Execute() error {
	return rootCmd.Execute()
}
