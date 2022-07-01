package util

import "github.com/spf13/cobra"

type RunFunc func(cmd *cobra.Command, args []string)
