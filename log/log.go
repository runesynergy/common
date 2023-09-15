package log

import (
	"fmt"
	"log"
	"os"

	"common.runesynergy.dev/config"
	"github.com/pkg/errors"
)

var debug = config.BoolOrDefault("debug", false)

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.SetOutput(os.Stdout)
}

func Panic(err any, depth ...int) {
	d := 2
	if len(depth) > 0 {
		d -= depth[0]
	}
	_ = log.Output(d, fmt.Sprint(err))
	panic(err)
}

func Debug(args ...any) {
	if debug {
		_ = log.Output(2, fmt.Sprintln(args...))
	}
}

func Debugf(format string, args ...any) {
	if debug {
		_ = log.Output(2, fmt.Sprintf(format, args...))
	}
}

func Print(args ...any) {
	_ = log.Output(2, fmt.Sprintln(args...))
}

func Printf(format string, args ...any) {
	_ = log.Output(2, fmt.Sprintf(format, args...))
}

func Error(err error) error {
	Printf("Error: %+v", err)
	return err
}

func Errorf(err error, format string, args ...any) error {
	return Error(errors.Wrapf(err, format, args...))
}
