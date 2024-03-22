package cmd

import (
	"log"

	"github.com/ThailanTec/cli-go-example/internal/calendar"
	"github.com/spf13/cobra"
)

var WeekCMD = &cobra.Command{
	Use:   "semana",
	Short: "Comprimisos da semana nengue",
	Long:  "Vou buscar seus compromissos da semana safado",
	Run: func(cmd *cobra.Command, args []string) {
		c := calendar.NewClient()
		err := c.GetAgendaID()
		if err != nil {
			log.Fatal(err)
		}
		c.ListWeekEvents(c.CalendarID)
	},
}
