package calendar

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os"
	"time"

	gCalendar "google.golang.org/api/calendar/v3"
	"google.golang.org/api/option"
)

const Agenda = "Rolés"

type Calendar struct {
	Service    *gCalendar.Service
	CalendarID string
}

func NewClient() *Calendar {
	ctx := context.Background()
	credential, err := os.ReadFile("./gcp.json")
	if err != nil {
		log.Fatal("Deu erro no Json mano")
	}

	service, err := gCalendar.NewService(ctx, option.WithCredentialsJSON(credential))
	if err != nil {
		log.Fatalf("deu erro no service meu mano: %s", err)
	}

	return &Calendar{
		Service: service,
	}
}

func (s *Calendar) GetAgendaID() (e error) {
	list, err := s.Service.CalendarList.List().Do()
	if err != nil {
		return err
	}
	for _, v := range list.Items {
		if v.Summary == Agenda {
			s.CalendarID = v.Id
		}
	}

	return nil
}

func (s *Calendar) InsertAgenda(id string) (e error) {
	entry := &gCalendar.CalendarListEntry{
		Id: id,
	}

	_, err := s.Service.CalendarList.Insert(entry).Do()
	if err != nil {
		return errors.New("erro ao adicionar agenda")
	}

	return nil
}
func (s *Calendar) ListWeekEvents(id string) error {
	now := time.Now()
	weekday := now.Weekday()
	starDate := now.AddDate(0, 7, -int(weekday))
	endDate := starDate.AddDate(0, 0, 7)

	events, err := s.Service.Events.List(s.CalendarID).TimeMin(starDate.Format(time.RFC3339)).TimeMax(endDate.Format(time.RFC3339)).Do()
	if err != nil {
		return err
	}

	for _, v := range events.Items {
		fmt.Printf("%s | %s | at %s\n", v.Summary, v.Status, v.Start.DateTime)
	}

	return nil
}
