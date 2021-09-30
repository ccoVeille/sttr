// Code generated by github.com/abhimanyu003/sttr/cmd/generate.go. DO NOT EDIT

package cmd

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/abhimanyu003/sttr/processors"
	"github.com/spf13/cobra"
)

var bcrypt_flag_r uint

func init() {	
	bcryptCmd.Flags().UintVarP(&bcrypt_flag_r, "number-of-rounds", "r", 10, "Number of rounds")
	rootCmd.AddCommand(bcryptCmd)
}

var bcryptCmd = &cobra.Command{
	Use:     "bcrypt",
	Short:   "Get the Bcrypt hash of your text",
	Aliases: []string{"bcrypt-hash"},
	Args:    cobra.MaximumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		var err error
		in, out := "", ""

		flags := make([]processors.Flag, 0)
		if len(args) == 0 {
			all, err := ioutil.ReadAll(cmd.InOrStdin())
			if err != nil {
				return err
			}
			in = string(all)
		} else {
			in = args[0]
		}

		p := processors.Bcrypt{}
		flags = append(flags, processors.Flag{Short: "r", Value: bcrypt_flag_r})

		out, err = p.Transform(in, flags...)
		if err != nil {
			return err
		}

		_, err = fmt.Fprint(os.Stdout, out)
		return err
	},
}
