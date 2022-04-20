package calendarutil

import (
	"fmt"
	"testing"
)

func TestSolarToLuanr(t *testing.T) {
	fmt.Println(SolarToLuanr(2021, 10, 15))
}

func TestLunarToSolar(t *testing.T) {
	fmt.Println(LunarToSolar(2021, 9, 10, false))
}
