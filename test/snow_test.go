package test

import (
	"fmt"
	"testing"
	"time"

	snow "DailyEnglish/utils"
)

func TestSnow(t *testing.T) {
	fmt.Println(snow.GenerateID(time.Now(), 114514191810))
}
