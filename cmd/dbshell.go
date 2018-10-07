package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/tizz98/godash/db"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"strings"
)

func init() {
	rootCmd.AddCommand(dbshell)
}

var dbshell = &cobra.Command{
	Use:   "dbshell",
	Short: "Connect to a database shell",
	Run: func(cmd *cobra.Command, args []string) {
		db := db.Connect()
		options := db.Options()

		split := strings.Split(options.Addr, ":")

		file, err := ioutil.TempFile("/tmp", ".pgpass.*.txt")
		if err != nil {
			log.Fatal(err)
		}
		defer os.Remove(file.Name())

		_, err = file.WriteString(fmt.Sprintf("%s:%s:%s:%s:%s", split[0], split[1], options.Database, options.User, options.Password))
		if err != nil {
			log.Fatal(err)
		}
		os.Setenv("PGPASSFILE", file.Name())
		defer os.Unsetenv("PGPASSFILE")

		command := exec.Command("psql", "-h", split[0], "-U", options.User, "-p", split[1], options.Database)
		command.Stdout = os.Stdout
		command.Stderr = os.Stderr
		command.Stdin = os.Stdin
		err = command.Run()
		if err != nil {
			log.Fatal(err)
		}
	},
}
