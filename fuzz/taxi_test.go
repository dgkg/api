package fuzz

import "testing"

func TestGetChartPlateNumber(t *testing.T) {
	res := getChartPlateNumber()
	t.Log(res)
}
