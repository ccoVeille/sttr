// Code generated by github.com/abhimanyu003/sttr/cmd/generate.go. DO NOT EDIT

package cmd

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/abhimanyu003/sttr/processors"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(titleCmd)
}

var titleCmd = &cobra.Command{
	Use:     "title",
	Short:   "Transform your text to Title Case",
	Aliases: []string{},
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

		p := processors.Title{}

		out, err = p.Transform(in, flags...)
		if err != nil {
			return err
		}

		_, err = fmt.Fprint(os.Stdout, out)
		return err
	},
}
