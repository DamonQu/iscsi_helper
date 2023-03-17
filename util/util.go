package util

import (
	"flag"
	"fmt"
	"os"
)



type FormatType_ int32

const (
	FormatType_ext4 FormatType_ = 0
	FormatType_ext3 FormatType_ = 1
	FormatType_xfs  FormatType_ = 2
)

type CommandVariable struct {
	Debug         bool // enable debug
}

func InitConfig(fs *flag.FlagSet) *CommandVariable {
	variables := &CommandVariable{
		Debug: false,
	}

	fs.BoolVar(&variables.Debug, "debug", false, "enable debugging")

	if err := fs.Parse(os.Args[1:]); err != nil {
		fmt.Print("unable to parse the configuration")
		os.Exit(1)
	}

	return variables
}