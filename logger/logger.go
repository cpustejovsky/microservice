package logger

import (
	"log"
	"os"
)

var Error = log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.LUTC|log.Llongfile)

var Info = log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime|log.LUTC)
