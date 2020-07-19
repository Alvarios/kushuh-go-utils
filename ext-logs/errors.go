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

func (s *ServerConfig) Print() string {
	return fmt.Sprintf(
		"*Environment*\n%s\n*Noticed*\n%s",
		s.Environment,
		time.Now().Format("2006-01-02 3:4:5"),
	)
}

func (s *ServerConfig) Error(t string) {
	_, _ = slack.Send(
		s.Webhook,
		fmt.Sprintf("Unexpected error in %s", s.Application),
		nil,
		[]map[string]interface{}{
			{
				"fallback" : fmt.Sprintf("Unexpected error in %s (%s)", s.Application, s.Environment),
				"color" : "#FF9300",
				"title" : t,
				"text" : s.Print(),
			},
		},
	)
}

func (s *ServerConfig) Errorf(t string, par ...interface{}) {
	_, _ = slack.Send(
		s.Webhook,
		fmt.Sprintf("Unexpected error in %s", s.Application),
		nil,
		[]map[string]interface{}{
			{
				"fallback" : fmt.Sprintf("Unexpected error in %s (%s)", s.Application, s.Environment),
				"color" : "#FF9300",
				"title" : fmt.Sprintf(t, par...),
				"text" : s.Print(),
			},
		},
	)
}

func (s *ServerConfig) Fatal(t string) {
	_, _ = slack.Send(
		s.Webhook,
		fmt.Sprintf("Critical error in %s", s.Application),
		nil,
		[]map[string]interface{}{
			{
				"fallback" : fmt.Sprintf("Critical error in %s (%s)", s.Application, s.Environment),
				"color" : "#ff3232",
				"title" : t,
				"text" : s.Print(),
			},
		},
	)

	log.Fatal(t)
}

func (s *ServerConfig) Fatalf(t string, par ...interface{}) {
	_, _ = slack.Send(
		s.Webhook,
		fmt.Sprintf("Critical error in %s", s.Application),
		nil,
		[]map[string]interface{}{
			{
				"fallback" : fmt.Sprintf("Critical error in %s (%s)", s.Application, s.Environment),
				"color" : "#ff3232",
				"title" : fmt.Sprintf(t, par...),
				"text" : s.Print(),
			},
		},
	)

	log.Fatalf(t, par...)
}