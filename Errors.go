package in3d

import (
	"errors"
	"fmt"
	"os"

	"github.com/fatih/color"
)

var RedTX = color.New(color.FgRed).SprintFunc()

func FormatError(err error, msg ...string) error {
	if err == nil {
		return nil
	}
	errMsg := ""
	if len(msg) > 0 {
		for _, m := range msg {
			errMsg += RedTX(m) + ": "
		}
	}
	errMsg += err.Error()
	return fmt.Errorf("\n💔 %s", errors.New(errMsg))
}

func EoE(err error, msg ...string) {
	err = FormatError(err, msg...)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
