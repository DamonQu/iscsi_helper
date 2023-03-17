package util

import "fmt"

type CommandLineOption_ int32

const (
	CommandLineOption_List_Sessions CommandLineOption_ = 0
	CommandLineOption_ISCSI_Login   CommandLineOption_ = 1
	CommandLineOption_ISCSI_Logout  CommandLineOption_ = 2
	CommandLineOption_MOUNT         CommandLineOption_ = 3
	CommandLineOption_Exit          CommandLineOption_ = 4
)

var CommandLineOptions = []string {
	"List ISCSI Sessions",
	"ISCSI Login",
	"ISCSI Logout",
	"Mount",
	"Exit",
}

func FirstPageHint() CommandLineOption_ {
	fmt.Println()
	selected := Hint("Select the action: ", "action you choose: ", CommandLineOptions, "%d")
	return mapCommandLineOption(selected)
}

func Hint(title, hint string, choices []string, scanFormat string) string {
	if title != "" {
		fmt.Println(title)
	}
	selectedChoice := false
	if choices != nil  && len(choices) != 0 {
		selectedChoice = true
		for index, choice := range choices {
			fmt.Printf(  "  [%d] %s\n", index + 1, choice)
		}
	}
	if selectedChoice {
		selected := 0
		for selected <= 0 || selected > len(choices) {
			fmt.Printf(hint)
			_, _ = fmt.Scanf(scanFormat, &selected)
			if selected <= 0 || selected > len(choices) {
				fmt.Printf("invalid index selected, the index should be in [1, %d].\n", len(choices))
			}
		}

		return choices[selected - 1]
	} else {
		var input string
		fmt.Printf(hint)
		_, _ = fmt.Scanf(scanFormat, &input)
		return input
	}
}

func mapCommandLineOption(option string) CommandLineOption_ {
	for index, lineOption := range CommandLineOptions {
		if lineOption == option {
			return CommandLineOption_(index)
		}
	}
	return CommandLineOption_Exit
}
