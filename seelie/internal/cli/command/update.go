package command

import (
	"fmt"

	"github.com/spf13/cobra"
)

// NewUpdateCmd ...
func NewUpdateCmd() *cobra.Command {
	return updateCmd
}

var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "update seelie client",
	Long:  "更新seelie客户端",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("todo: need to implement")
		// step 1. check version

		// step 2. update new binary file

		// step 3. extract and mv to replace old binary
	},
}
