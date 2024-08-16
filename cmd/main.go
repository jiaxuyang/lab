package main

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"strconv"
	"time"
)

func init() {
	rootCmd.AddCommand(tsCmd)
}

var tsCmd = &cobra.Command{
	Use:  "ts",
	Args: cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		var (
			seconds int64
			str     string
		)
		if i, err := strconv.ParseInt(args[0], 10, 64); err == nil {
			t := time.Unix(i, 0)
			seconds, str = t.Unix(), t.Format(time.DateTime)
		} else if t, err := time.ParseInLocation(time.DateTime, args[0], time.Local); err == nil {
			seconds, str = t.Unix(), t.Format(time.DateTime)
		} else if t, err := time.ParseInLocation(time.DateOnly, args[0], time.Local); err == nil {
			seconds, str = t.Unix(), t.Format(time.DateTime)
		} else {
			return err
		}
		fmt.Println(seconds)
		fmt.Println(str)
		return nil
	},
}

var rootCmd = &cobra.Command{
	Use:   "lab",
	Short: "lab is a useful cli tool",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		// Do Stuff Here
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func main() {
	Execute()
}
