package appinit

import (
	"log"
	"os"
)

const (
	EXIT_BAD_CONFIG         = 2
	EXIT_EAL_INIT_ERROR     = 3
	EXIT_EAL_LAUNCH_ERROR   = 4
	EXIT_MEMPOOL_INIT_ERROR = 5
	EXIT_FACE_INIT_ERROR    = 6
)

func Exitf(exitCode int, format string, v ...interface{}) {
	log.Printf(format, v...)
	os.Exit(exitCode)
}