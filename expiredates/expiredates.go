package expiredates

import (
	"encoding/csv"
	"fmt"
	"os"
	"time"

	"github.com/hsmtkk/mf-vix-strategy/config"
)

func GetExpireDates() ([]time.Time, error) {
	f, err := os.Open(config.EXPIRE_DATES_CSV)
	if err != nil {
		return nil, fmt.Errorf("failed to open file %s: %w", config.EXPIRE_DATES_CSV, err)
	}
	defer f.Close()
	reader := csv.NewReader(f)
	reader.Comment = '#'
	records, err := reader.ReadAll()
	if err != nil {
		return nil, fmt.Errorf("failed to read CSV file: %w", err)
	}
	results := []time.Time{}
	for _, record := range records {
		t, err := time.Parse(config.DATE_FORMAT, record[0])
		if err != nil {
			return nil, fmt.Errorf("failed to parse date %s: %w", record, err)
		}
		// 過去の日付を除く
		if t.Before(time.Now()) {
			continue
		}
		results = append(results, t)
	}
	// 最低でも8ヶ月分、入力済みとする
	if len(results) < 8 {
		return nil, fmt.Errorf("not enough expire dates")
	}
	return results, nil
}
