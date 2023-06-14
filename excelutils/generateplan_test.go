package excelutils

import (
	"log"
	"path"
	"testing"

	"github.com/mitchellh/go-homedir"
)

func TestGeneratePlan(t *testing.T) {
	homedir.DisableCache = false
	var month = [3]int{1, 2, 3}
	dir, err := homedir.Dir()
	if err != nil {
		log.Fatal(err)
	}
	var path = path.Join(dir, "plan.xlsx")
	GeneratePlan(2022, month[:], path)
}
