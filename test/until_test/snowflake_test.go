package utiltest

import (
	"DailyEnglish/utils/snowflake"
	"fmt"
	"testing"
	"time"

	sf "github.com/bwmarrin/snowflake"
)

func TestGenerateID(t *testing.T) {
	err := snowflake.Init("2023-01-01", 1)
	if err != nil {
		t.Errorf("Failed to initialize snowflake node: %v", err)
		return
	}

	ids := make([]int64, 100)
	for i := range ids {
		ids[i] = snowflake.GenerateID()
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
