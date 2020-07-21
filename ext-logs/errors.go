package extLogs

import (
	"fmt"
	"github.com/Alvarios/kushuh-go-utils/ext-apis/slack"
	"log"
	"time"
)

type ServerConfig struct {
	Webhook string
	Application string
	Environment string
}

func (s *ServerConfig) Print(f string, t string) string {
	return fmt.Sprintf(
		"%s\n\n*File*\n%s\n\n*Environment*\n%s\n\n*Noticed*\n%s",
		t,
		f,
		s.Environment,
		time.Now().Format("2006-01-02 3:4:5"),
	)
}

func (s *ServerConfig) Error(f string, t string) {
	_, _ = slack.Send(
		s.Webhook,
		fmt.Sprintf("Unexpected error in %s", s.Application),
		nil,
		[]map[string]interface{}{
			{
				"fallback" : fmt.Sprintf("Unexpected error in %s (%s)", s.Application, s.Environment),
				"color" : "#FF9300",
				"text" : s.Print(f, t),
			},
		},
	)
}

func (s *ServerConfig) Fatal(f string, t string) {
	_, _ = slack.Send(
		s.Webhook,
		fmt.Sprintf("Critical error in %s", s.Application),
		nil,
		[]map[string]interface{}{
			{
				"fallback" : fmt.Sprintf("Critical error in %s (%s)", s.Application, s.Environment),
				"color" : "#ff3232",
				"text" : s.Print(f, t),
			},
		},
	)

	log.Fatal(t)
}