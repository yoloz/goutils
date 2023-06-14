package excelutil

import (
	"log"
	"path"
	"testing"

	"github.com/yoloz/goutils/fileutil"
)

func TestGeneratePlan(t *testing.T) {
	var month = [3]int{1, 2, 3}
	dir, err := fileutil.Userhome()
	if err != nil {
		log.Fatal(err)
	}
	var path = path.Join(dir, "plan.xlsx")
	GeneratePlan(2022, month[:], path)
}
