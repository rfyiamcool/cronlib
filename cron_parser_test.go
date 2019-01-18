package cronlib

import (
	"testing"
	"time"
)

func TestParse(t *testing.T) {
	cron, err := Parse("1 0 0 */1 * *")
	if err != nil {
		t.Error(err)
	}

	now1 := time.Date(2018, 12, 11, 05, 22, 33, 0, time.UTC)
	nowAt := cron.Next(now1)
	if nowAt != time.Date(2018, 12, 12, 0, 0, 1, 0, time.UTC) {
		t.Error("err")
	}
}
