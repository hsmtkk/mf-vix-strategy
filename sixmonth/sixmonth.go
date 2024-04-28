package sixmonth

import (
	"fmt"
	"log"
	"os"

	"github.com/hsmtkk/mf-vix-strategy/config"
	"github.com/hsmtkk/mf-vix-strategy/expiredates"
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/spf13/cobra"
)

var Command = &cobra.Command{
	Use: "sixmonth",
	Run: run,
}

func run(cmd *cobra.Command, args []string) {
	expireDates, err := expiredates.GetExpireDates()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("PUT Delta 0.9を購入する")
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"限月", "満期"})
	t.AppendSeparator()
	for i := 0; i < 6; i++ {
		t.AppendRow(table.Row{i + 1, expireDates[i].Format(config.DATE_FORMAT)})
	}
	t.Render()
}
