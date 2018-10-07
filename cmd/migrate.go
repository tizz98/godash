package cmd

import (
	"fmt"
	"os"

	"github.com/go-pg/migrations"
	"github.com/spf13/cobra"
	"github.com/tizz98/godash/db"
	_ "github.com/tizz98/godash/db/migrations"
)

func init() {
	rootCmd.AddCommand(migrate)
}

var migrate = &cobra.Command{
	Use:   "migrate",
	Short: "Run migrations",
	Long: `This program runs command on the db. Supported commands are:
  - up [target] - runs all available migrations by default or up to target one if argument is provided.
  - down - reverts last migration.
  - reset - reverts all migrations.
  - version - prints current db version.
  - set_version [version] - sets db version without running migrations.`,
	Run: func(cmd *cobra.Command, args []string) {
		db := db.Connect()

		oldVersion, newVersion, err := migrations.Run(db, args...)

		if err != nil {
			exitf(err.Error())
		}
		if newVersion != oldVersion {
			fmt.Printf("migrated from version %d to %d\n", oldVersion, newVersion)
		} else {
			fmt.Printf("version is %d\n", oldVersion)
		}
	},
}

func errorf(s string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, s+"\n", args...)
}

func exitf(s string, args ...interface{}) {
	errorf(s, args...)
	os.Exit(1)
}
