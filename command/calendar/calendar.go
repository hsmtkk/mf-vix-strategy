package calendar

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/hsmtkk/mf-vix-strategy/config"
	"github.com/hsmtkk/mf-vix-strategy/expiredates"
	"github.com/hsmtkk/mf-vix-strategy/weekstofirstmonth"
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/spf13/cobra"
)

var Command = &cobra.Command{
	Use: "calendar",
	Run: run,
}

func run(cmd *cobra.Command, args []string) {
	expireDates, err := expiredates.GetExpireDates()
	if err != nil {
		log.Fatal(err)
	}

	restWeeks := weekstofirstmonth.CalculateWeeksToFirstMonth(expireDates[0])
	fmt.Printf("第1限月満期までの残り週数: %d\n", restWeeks)
	fmt.Println()

	printPut(expireDates, restWeeks)
	fmt.Println()
	printCall(expireDates, restWeeks)
}

// Put買い
func printPut(expireDates []time.Time, restWeeks int) {
	firstPuts := restWeeks
	secondPuts := 5 - firstPuts
	firstMonth := expireDates[0]
	secondMonth := expireDates[1]

	fmt.Println("PUT Delta 0.8を購入する")
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"限月", "最終取引日", "枚数"})
	t.AppendSeparator()
	t.AppendRows([]table.Row{
		{1, firstMonth.Format(config.DATE_FORMAT), firstPuts},
		{2, secondMonth.Format(config.DATE_FORMAT), secondPuts},
	})
	t.Render()
}

// Call買い
func printCall(expireDates []time.Time, restWeeks int) {
	secondCalls := restWeeks
	thirdCalls := 5 - secondCalls
	secondMonth := expireDates[1]
	thirdMonth := expireDates[2]

	fmt.Println("CALL Delta 0.8を購入する")
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"限月", "最終取引日", "枚数"})
	t.AppendSeparator()
	t.AppendRows([]table.Row{
		{2, secondMonth.Format(config.DATE_FORMAT), secondCalls},
		{3, thirdMonth.Format(config.DATE_FORMAT), thirdCalls},
	})
	t.Render()
}
