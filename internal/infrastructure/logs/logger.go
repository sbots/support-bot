package logs

import (
	log "github.com/sirupsen/logrus"
	"os"
)

func InitLogger(logLever string, _ bool) {
	level, err := log.ParseLevel(logLever)
	if err != nil {
		panic("can not parse log level")
	}
	log.SetOutput(os.Stdout)
	log.SetLevel(level)
}
