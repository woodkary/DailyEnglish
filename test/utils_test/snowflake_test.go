package test

import (
	snowflake "DailyEnglish/utils"
	"fmt"
	"testing"
	"time"

	sf "github.com/bwmarrin/snowflake"
)

func TestGenerateID(t *testing.T) {
	ids := make([]int64, 100)
	for i := range ids {
		ids[i] = snowflake.GenerateID(time.Now(), 1145141919810)
	}

	uniqueIDs := make(map[int64]bool)
	for _, id := range ids {
		if _, ok := uniqueIDs[id]; ok {
			t.Errorf("Duplicate ID generated: %d", id)
			return
		}
		uniqueIDs[id] = true
	}

	for _, id := range ids {
		if len(fmt.Sprint(id)) != 18 {
			t.Errorf("Invalid ID length: %d", len(fmt.Sprint(id)))
			return
		}

		timestamp := (id >> 22) + (sf.Epoch / 1000)
		expectedTimestamp := time.Now().Unix() - 30
		if timestamp < expectedTimestamp {
			t.Errorf("ID generated before expected time: %d < %d", timestamp, expectedTimestamp)
			return
		}
	}
}
