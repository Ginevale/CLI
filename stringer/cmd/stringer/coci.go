package stringer

import (
	"fmt"

	"github.com/spf13/cobra"
)

var cocilovaCmd = &cobra.Command{
	Use:     "cocilova",
	Aliases: []string{"coci"},
	Short:   "Opinion on cocilova",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("hi my name is coci")
	},
}

func init() {
	rootCmd.AddCommand(cocilovaCmd)
}
