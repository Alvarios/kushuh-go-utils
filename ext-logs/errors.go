package extLogs

import (
	"fmt"
	"github.com/Alvarios/kushuh-go-utils/ext-apis/slack"
	"log"
)

func Error(u string, s string) {
	_, _ = slack.Send(u, s)
}

func Errorf(u string, s string, par ...interface{}) {
	_, _ = slack.Send(u, fmt.Sprintf(s, par...))
}

func Fatal(u string, s string) {
	_, _ = slack.Send(u, s)
	log.Fatal(s)
}

func Fatalf(u string, s string, par ...interface{}) {
	_, _ = slack.Send(u, fmt.Sprintf(s, par...))
	log.Fatalf(s, par...)
}