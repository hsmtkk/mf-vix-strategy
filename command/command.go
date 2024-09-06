package command

import (
	"github.com/hsmtkk/mf-vix-strategy/command/calendar"
	"github.com/hsmtkk/mf-vix-strategy/command/halfvxxvxz"
	"github.com/hsmtkk/mf-vix-strategy/command/vxxvxz"
	"github.com/hsmtkk/mf-vix-strategy/sixmonth"
	"github.com/spf13/cobra"
)

var Command = &cobra.Command{}

func init() {
	Command.AddCommand(calendar.Command)
	Command.AddCommand(sixmonth.Command)
	Command.AddCommand(vxxvxz.Command)
	Command.AddCommand(halfvxxvxz.Command)
}
