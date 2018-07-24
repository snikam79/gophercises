package cobra

import (
	"fmt"

	"secret"

	"github.com/spf13/cobra"
)

var setCmd = &cobra.Command{
	Use:   "set",                                  // it would the name of the command
	Short: "Sets a secret in your secret storage", // command description
	//anonymous function
	Run: func(cmd *cobra.Command, args []string) {
		v := secret.File(encodingKey, secretsPath())
		key, value := args[0], args[1]
		err := v.SetKey(key, value)
		if err != nil {
			panic(err)
		}
		fmt.Println("Value set successfully!")
	},
}

// we are doing this for making setCmd available.
func init() {
	RootCmd.AddCommand(setCmd)
}
