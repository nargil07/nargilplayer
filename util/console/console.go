package console

import (
	"fmt"
	"github.com/fatih/color"
	"github.com/nargil07/nargilplayer/util"
)

func NewOptionsLine(first string, second string, description string) {
	opt := color.GreenString(util.FormatLength(first, 2))
	option := color.GreenString(util.FormatLength(second, 15))

	fmt.Printf("  %s, %s %s\n", opt, option, description)
}

func NewCommandLine(command string, description string) {
	cmd := color.GreenString(util.FormatLength(command, 15))
	fmt.Printf("  %s : %s", cmd, description)
}
