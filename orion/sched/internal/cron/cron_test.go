package cron

import (
	"testing"

	"weibo.com/opendcp/orion/models"

	"github.com/stretchr/testify/assert"
)

func TestToCronExpression(t *testing.T) {
	var (
		assert = assert.New(t)
		tt     = []struct {
			item models.CronItem
			exp  string
			fail bool
		}{
			{
				item: models.CronItem{WeekDay: 0, Time: "10:00"},
				exp:  "00 00 10 * * *",
			},
			{
				item: models.CronItem{WeekDay: 0, Time: "01:00"},
				exp:  "00 00 01 * * *",
			},
			{
				item: models.CronItem{WeekDay: 2, Time: "02:30"},
				exp:  "00 30 02 * * 1",
			},
			{
				item: models.CronItem{WeekDay: 0, Time: "23:00"},
				exp:  "00 00 23 * * *",
			},
		}
	)

	for i, v := range tt {
		assert.Equal(v.exp, toCronExpression(&v.item), "%d", i)
	}
}
