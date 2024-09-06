package halfvxxvxz

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/hsmtkk/mf-vix-strategy/config"
	"github.com/hsmtkk/mf-vix-strategy/daystofirstmonth"
	"github.com/hsmtkk/mf-vix-strategy/expiredates"
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/spf13/cobra"
)

const TOTAL_QUANTITY = 2
const VXZ_START_INDEX = 4

var Command = &cobra.Command{
	Use: "halfvxxvxz",
	Run: run,
}

func run(cmd *cobra.Command, args []string) {
	expireDates, err := expiredates.GetExpireDates()
	if err != nil {
		log.Fatal(err)
	}
	printVXX(expireDates)
	fmt.Println()
	printVXZ(expireDates)
}

func printVXX(expireDates []time.Time) {
	firstMonth := expireDates[0]
	secondMonth := expireDates[1]

	restDays := daystofirstmonth.CalculateDaysToFirstMonth(expireDates[0])
	fmt.Printf("第1限月満期までの残り日数: %d\n", restDays)
	/*
		残り日数
		0-6: 0枚
		7-20: 1枚
		21-34: 2枚
	*/
	firstMonthPuts := (restDays+21)/14 - 1
	secondMonthPuts := TOTAL_QUANTITY - firstMonthPuts
	fmt.Println()

	fmt.Println("VXX売り複製")
	fmt.Println("PUT Delta 0.9を購入する")
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"限月", "最終取引日", "枚数"})
	t.AppendSeparator()
	t.AppendRows([]table.Row{
		{1, firstMonth.Format(config.DATE_FORMAT), firstMonthPuts},
		{2, secondMonth.Format(config.DATE_FORMAT), secondMonthPuts},
	})
	t.Render()
}

func printVXZ(expireDates []time.Time) {
	fmt.Println("VXZ買い複製")
	fmt.Println("CALL Delta 0.9を購入する")
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"限月", "最終取引日", "枚数"})
	t.AppendSeparator()
	for i := VXZ_START_INDEX; i < VXZ_START_INDEX+TOTAL_QUANTITY; i++ {
		t.AppendRow([]interface{}{i + 1, expireDates[i].Format(config.DATE_FORMAT), 1})
	}
	t.Render()
}
