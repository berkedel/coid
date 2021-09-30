package main

import (
	"coid/util"
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var (
	compressMode bool
	rootCmd      = &cobra.Command{
		Use:   "util <uuid>",
		Short: "util is a command to compress and decompress uuid",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			if compressMode == true {
				fmt.Println(util.Compress(args[0]))
			} else {
				fmt.Println(util.Decompress(args[0]))
			}
		}}
)

func init() {
	rootCmd.PersistentFlags().BoolVarP(&compressMode, "compress", "c", false, "Toggle compress")
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
