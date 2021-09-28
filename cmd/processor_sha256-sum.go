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
	rootCmd.AddCommand(sha256SumCmd)
}

var sha256SumCmd = &cobra.Command{
	Use:     "sha256-sum",
	Short:   "Get the SHA256 checksum of your text",
	Aliases: []string{"sha256"},
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
			if fi, err := os.Stat(args[0]); err == nil && !fi.IsDir() {
				d, err := ioutil.ReadFile(args[0])
				if err != nil {
					return err
				}
				in = string(d)
				flags = append(flags, processors.Flag{
					Name:  processors.FlagFile,
					Value: true,
				})
			} else {
				in = args[0]
			}
		}

		p := processors.SHA256{}

		out, err = p.Transform(in, flags...)
		if err != nil {
			return err
		}

		_, err = fmt.Fprint(os.Stdout, out)
		return err
	},
}
