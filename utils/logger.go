package utils

import (
	"fmt"
	"github.com/sirupsen/logrus"
)

type LogWriter struct {
}
type LogFields logrus.Fields

func (writer LogWriter) Write(bytes []byte) (int, error) {
	return fmt.Println(string(bytes))
}

func Debug(customText string, fields LogFields) {
	logrus.WithFields(logrus.Fields(fields)).Info(customText)
}
