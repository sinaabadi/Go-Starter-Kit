package utils

import (
	"fmt"
)

type LogWriter struct {
}

func (writer LogWriter) Write(bytes []byte) (int, error) {
	return fmt.Println(string(bytes))
}

