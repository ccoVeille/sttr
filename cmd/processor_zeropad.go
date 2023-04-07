// Code generated by github.com/abhimanyu003/sttr/cmd/generate.go. DO NOT EDIT

package cmd

import (
	"fmt"
	"io"
	"os"

	"github.com/abhimanyu003/sttr/processors"
	"github.com/spf13/cobra"
)

var (		
	zeropad_flag_n uint		
	zeropad_flag_p string
)

func init() {	
	zeropadCmd.Flags().UintVarP(&zeropad_flag_n, "number-of-zeros", "n", 5, "Number of zeros to be padded")
	zeropadCmd.Flags().StringVarP(&zeropad_flag_p, "prefix", "p", "", "The number get prefixed with this")
	rootCmd.AddCommand(zeropadCmd)
}

var zeropadCmd = &cobra.Command{
	Use:     "zeropad [string]",
	Short:   "Pad a number with zeros",
	Aliases: []string{},
	Args:    cobra.MaximumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		var err error
		var in []byte
		var out string

		if len(args) == 0 {
			in, err = io.ReadAll(cmd.InOrStdin())
			if err != nil {
				return err
			}
		} else {
			if fi, err := os.Stat(args[0]); err == nil && !fi.IsDir() {
				d, err := os.ReadFile(args[0])
				if err != nil {
					return err
				}
				in = d
			} else {
				in = []byte(args[0])
			}
		}

		flags := make([]processors.Flag, 0)
		p := processors.Zeropad{}
		flags = append(flags, processors.Flag{Short: "n", Value: zeropad_flag_n})
		flags = append(flags, processors.Flag{Short: "p", Value: zeropad_flag_p})

		out, err = p.Transform(in, flags...)
		if err != nil {
			return err
		}

		_, err = fmt.Fprint(os.Stdout, out)
		return err
	},
}