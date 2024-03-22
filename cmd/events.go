package cmd

import (
	"log"

	"github.com/ThailanTec/cli-go-example/internal/calendar"
	"github.com/spf13/cobra"
)

var EventsCMD = &cobra.Command{
	Use:   "dia",
	Short: "Vou buscar na sua agenda os compromissos",
	Long:  "Vou buscar na sua agenda do google, todos os seus compromissos seu esquecido safado",
	Run: func(cmd *cobra.Command, args []string) {
		c := calendar.NewClient()
		err := c.GetAgendaID()
		if err != nil {
			log.Fatal(err)
		}
	},
}
