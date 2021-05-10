//遍历文件生成docsify中的sidebar,readme
package main

import (
	"container/list"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func main() {

	var docsDir string
	fmt.Printf("Please enter docsify directory: ")
	fmt.Scanln(&docsDir)

	fd, err := os.ReadDir(docsDir)
	if err != nil {
		log.Fatal(err)
	}
	list := list.New()
	for _, catlog := range fd { //遍历目录
		if catlog.IsDir() && !strings.HasPrefix(catlog.Name(), "_") && !strings.HasPrefix(catlog.Name(), ".") {
			list.Init()
			catalogPath := filepath.Join(docsDir, catlog.Name())
			td, err := os.ReadDir(catalogPath)
			if err != nil {
				log.Fatal(err)
			}
			list.PushBack("- 文档\n")
			for _, file := range td { //遍历文件
				if file.Type().IsRegular() && !strings.HasPrefix(file.Name(), "_") && strings.Compare(file.Name(), "README.md") != 0 {
					list.PushBack("  - [" + strings.TrimSuffix(file.Name(), ".md") + "](/" + catlog.Name() + "/" + file.Name() + ")\n")
				}
			}
			//生成文件
			siderbar, err := os.Create(filepath.Join(catalogPath, "_sidebar.md"))
			if err != nil {
				log.Fatal(err)
			}
			defer siderbar.Close()

			readme, err := os.Create(filepath.Join(catalogPath, "README.md"))
			if err != nil {
				log.Fatal(err)
			}
			defer readme.Close()

			for e := list.Front(); e != nil; e = e.Next() {
				line := e.Value.(string)
				siderbar.WriteString(line)
				readme.WriteString(line)
			}

		}
	}
}
