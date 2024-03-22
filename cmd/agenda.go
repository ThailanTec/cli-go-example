package cmd

import (
	"fmt"
	"log"

	"github.com/ThailanTec/cli-go-example/internal/calendar"
	"github.com/spf13/cobra"
)

var AgendaCMD = &cobra.Command{
	Use:   "agenda",
	Short: "Vou buscar na agenda",
	Long:  "Vou buscar na sua agenda do google, todos os seus compromissos seu esquecido safado",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		c := calendar.NewClient()
		err := c.InsertAgenda(args[0])
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("Deu green irmão")
	},
}
