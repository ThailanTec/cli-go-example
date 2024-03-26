package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func NewRootCMD() *cobra.Command {
	rootCMD := &cobra.Command{
		Use:   "google",
		Short: "Trago suas informações em menos de 7 dias",
	}

	rootCMD.AddCommand(EventsCMD)
	rootCMD.AddCommand(AgendaCMD)
	rootCMD.AddCommand(WeekCMD)
	return rootCMD
}

func Execute() {
	if err := NewRootCMD().Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
