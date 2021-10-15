package calendarUtil

import (
	"fmt"
	"testing"
)

func TestSolarToLuanr(t *testing.T) {
	solarDate := "2021-10-15"
	fmt.Println(SolarToLuanr(solarDate))
}

func TestLunarToSolar(t *testing.T) {
	lunarDate := "2021-09-10"
	fmt.Println(LunarToSolar(lunarDate, false))
}
