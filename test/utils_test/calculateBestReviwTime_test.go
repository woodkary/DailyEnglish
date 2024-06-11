package test

import (
	utils "DailyEnglish/utils"
	"fmt"
	"testing"
)

func TestCalculate(t *testing.T) {
	a := 18
	b := "0,1,14"
	c := "0,0,1"
	day := utils.CalculateBestInterval(a, b, c)
	fmt.Println(day)
}
